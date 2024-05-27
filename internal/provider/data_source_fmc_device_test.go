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

func TestAccDataSourceFmcDevice(t *testing.T) {
	if os.Getenv("FMC_DEV") == "" {
		t.Skip("skipping test, set environment variable FMC_DEV")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "name", "device1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "license_caps.0", "BASE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "type", "Device"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDevicePrerequisitesConfig + testAccDataSourceFmcDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"ssh": {
				VersionConstraint: "2.7.0",
				Source:            "loafoe/ssh",
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceFmcDevicePrerequisitesConfig = `
resource "fmc_access_control_policy" "test" {
  name = "POLICY1"
  default_action = "BLOCK"

#   FTDv does not handle this ssh method: "Failed to upload script"
#   connection {
#     type = "ssh"
#     user = "admin"
#     password = "${var.dev_password}"
#     host = "${var.dev}"
#     # agent = false
#     timeout = "30s"
#   }

#   provisioner "remote-exec" {
#     inline = [
#       "configure manager delete 10.50.202.44",
#       # "configure manager add 10.50.202.44 key1",
#     ]
#   }
}

// Set all these with TF_VAR_varname:
variable "dev" {}
variable "dev_password" {}

resource "ssh_resource" "preconfigured_device" {
  host         = "${var.dev}"
  user         = "admin"
  password     = "${var.dev_password}"
  agent        = false

  commands = [
    "configure manager delete 10.50.202.44",
    "configure manager add 10.50.202.44 key1",
  "show managers verbose",
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcDeviceConfig() string {
	config := `resource "fmc_device" "test" {` + "\n"
	config += `	name = "device1"` + "\n"
	config += `	host_name = var.dev` + "\n"
	config += `	license_caps = ["BASE"]` + "\n"
	config += `	reg_key = "key1"` + "\n"
	config += `	type = "Device"` + "\n"
	config += `	access_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device" "test" {
			id = fmc_device.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
