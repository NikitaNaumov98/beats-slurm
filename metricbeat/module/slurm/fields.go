// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package slurm

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "slurm", asset.ModuleFieldsPri, AssetSlurm); err != nil {
		panic(err)
	}
}

// AssetSlurm returns asset data.
// This is the base64 encoded zlib format compressed contents of module/slurm.
func AssetSlurm() string {
	return "eJy0lD9z4jAQxXt/ijf03Adwcc1dc1eloc7I1gICyetopSHOp8/IGOK/CUPIdkjo996unrXGkZocYqN3GRBMsJRj1f5eZYAnS0oox05lgCYpvamD4SrH7wzA+SQc62gJJVdBmUrgKHhTCrbs4cixb6AqjbKOiKJ2BN7iwIVkwNaQ1ZK3sDUq5ejDTqrQ1Enec6y7lRkXQ0wfdeDiOQr568aFeKTmxF731me559oI+cSDEuHSqEAaJxP2CT8nafREz3K1u03sPxf493eClUD197pI4AGlP6SyjhN6f+qprmkoKKjbNDv0NBrjSf552gyz8avHGd9u33xZxxiMHexdGtAcC0ujrU/cpmqNBGPNm0p/6ewsSgvpx0r3ZrAk6unlkZqSh70npQU+tTwbD0fup+LhyH0dj+4VuSshjlx7bnZmo+/yhonNWFmSXbqoO0TVq3HRQTmOVUiynY30xCYH7wEAAP//oPiJ/g=="
}