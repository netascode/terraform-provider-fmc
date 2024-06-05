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
	"github.com/netascode/terraform-provider-fmc/internal/ftd"
)

// End of section. //template:end imports

func TestAccDataSourceFmcDevice(t *testing.T) {
	ftdAddr := os.Getenv("FTD_ADDR")
	if ftdAddr == "" {
		t.Skip("skipping test, set environment variable FTD_ADDR")
	}
	// Prepare Terraform's var.ftd_addr
	os.Setenv("TF_VAR_ftd_addr", ftdAddr)
	// Prepare Terraform's var.reg_key
	os.Setenv("TF_VAR_reg_key", ftd.MustRandomizeKey())

	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "name", "device1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "license_caps.0", "BASE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device.test", "type", "Device"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: ftd.MustInitFromEnv,
				Config:    testAccDataSourceFmcDevicePrerequisitesConfig + testAccDataSourceFmcDeviceConfig(),
				Check:     resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceFmcDevicePrerequisitesConfig = `
resource "fmc_access_control_policy" "minimum" {
  name = "POLICY1"
  default_action = "BLOCK"
}

resource "fmc_access_control_policy" "test" {
  name = "POLICY2"
  default_action = "PERMIT"
}

variable "ftd_addr" {} // tests will set $TF_VAR_ftd_addr

variable "reg_key" {} // tests will set $TF_VAR_reg_key

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcDeviceConfig() string {
	config := `resource "fmc_device" "test" {` + "\n"
	config += `	name = "device1"` + "\n"
	config += `	host_name = var.ftd_addr` + "\n"
	config += `	license_caps = ["BASE"]` + "\n"
	config += `	reg_key = var.reg_key` + "\n"
	config += `	type = "Device"` + "\n"
	config += `	access_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	nat_policy_id = null` + "\n"
	config += `	prohibit_packet_transfer = false` + "\n"
	config += `	performance_tier = "FTDv5"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device" "test" {
			id = fmc_device.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
