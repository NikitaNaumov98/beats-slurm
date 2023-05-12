package job_mem

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

var slurmdir string
var steps bool

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("slurm", "job_mem", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	uid int
	jobid int
	step string
	memusage int
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	slurmdir = "/sys/fs/cgroup/memory/slurm"

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		uid: 0,
		jobid: 0,
		step: "",
		memusage: 0,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	var curr_uid_dir string
	var curr_job_dir string
	var curr_step_dir string

	entries, err := os.ReadDir(slurmdir)
	if err != nil {
		return fmt.Errorf("failed to access slurm cgroup directory: %w", err)
	}

	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "uid_") {
			m.uid, err = strconv.Atoi(strings.Split(e.Name(), "_")[1])
			if err != nil {
				m.Logger().Errorf("failed to convert uid into int: %s", err)
				continue
			}
			curr_uid_dir = slurmdir + "/" + e.Name()
			entries_uid, err := os.ReadDir(curr_uid_dir)
			if err != nil {
				m.Logger().Errorf("failed to access directories associated with uid: %s", err)
				continue
			}
			for _, f := range entries_uid {
				if strings.HasPrefix(f.Name(), "job_") {
					steps = false
					m.jobid, err = strconv.Atoi(strings.Split(f.Name(), "_")[1])
					if err != nil {
						m.Logger().Errorf("failed to convert jobid into int: %s", err)
						continue
					}
					curr_job_dir = curr_uid_dir + "/" + f.Name() + "/"
					entries_job, err := os.ReadDir(curr_job_dir)
					if err != nil {
						m.Logger().Errorf("failed to access directories associated with jobs: %s", err)
						continue
					}
					for _, g := range entries_job {
						if strings.HasPrefix(g.Name(), "step_") {
							steps = true
							m.step = strings.Split(g.Name(), "_")[1]
							curr_step_dir = curr_job_dir + g.Name() + "/"
							usageval, err := os.ReadFile(curr_step_dir + "memory.usage_in_bytes")
							if err != nil {
								m.Logger().Errorf("failed to get value of memory.usage_in_bytes: %s", err)
								continue
							}
							m.memusage, err = strconv.Atoi(strings.TrimSpace(string(usageval)))
							if err != nil {
								m.Logger().Errorf("failed to parse value of memory.usage_in_bytes as int: %s", err)
								continue
							}
						}
					}
					if !steps {
						usageval, err := os.ReadFile(curr_job_dir + "memory.max_usage_in_bytes")
						if err != nil {
							m.Logger().Errorf("failed to get value of memory.max_usage_in_bytes: %s", err)
							continue
						}
						m.memusage, err = strconv.Atoi(strings.TrimSpace(string(usageval)))
						if err != nil {
							m.Logger().Errorf("failed to parse value of memory.max_usage_in_bytes as int: %s", err)
							continue
						}
					}

					report.Event(mb.Event{
						MetricSetFields: mapstr.M{
							"uid": m.uid,
							"jobid": m.jobid,
							"step": m.step,
							"memusage": m.memusage,
						},
					})
				}
			}
		}
	}

	return nil
}
