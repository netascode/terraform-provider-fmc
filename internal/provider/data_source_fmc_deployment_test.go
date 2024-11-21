// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeployment(t *testing.T) {
	if os.Getenv("TF_VAR_timestamp") == "" && os.Getenv("TF_VAR_device_id_list") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_timestamp or TF_VAR_device_id_list")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeploymentPrerequisitesConfig + testAccDataSourceFmcDeploymentConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeploymentPrerequisitesConfig = `
variable "timestamp" { default = null } // tests will set $TF_VAR_timestamp
variable "device_id_list" {             // tests will set $TF_VAR_device_id_list
  type = list(string)
  default = null 
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeploymentConfig() string {
	config := `resource "fmc_deployment" "test" {` + "\n"
	config += `	version = var.timestamp` + "\n"
	config += `	ignore_warning = true` + "\n"
	config += `	device_list = var.device_id_list` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_deployment" "test" {
			id = fmc_deployment.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
