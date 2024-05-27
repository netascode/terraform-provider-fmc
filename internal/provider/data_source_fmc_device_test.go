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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceFmcDevice(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "name", "device1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "host_name", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "license_caps.0", "BASE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "type", "Device"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "access_policy_id", "fmc_access_control_policy.test.id"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDevicePrerequisitesConfig + testAccDataSourceFmcDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceFmcDevicePrerequisitesConfig = `
resource "fmc_access_control_policy" "test" {
  name = "POLICY1"
  default_action = "BLOCK"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcDeviceConfig() string {
	config := `resource "fmc_device" "test" {` + "\n"
	config += `	name = "device1"` + "\n"
	config += `	host_name = "10.0.0.1"` + "\n"
	config += `	license_caps = ["BASE"]` + "\n"
	config += `	reg_key = "key1"` + "\n"
	config += `	type = "Device"` + "\n"
	config += `	access_policy_id = "fmc_access_control_policy.test.id"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device" "test" {
			id = fmc_device.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
