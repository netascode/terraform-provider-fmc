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

func TestAccDataSourceFmcDeviceBFD(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bfd.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bfd.test", "hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bfd.test", "slow_timer", "1000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceBFDPrerequisitesConfig + testAccDataSourceFmcDeviceBFDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceBFDPrerequisitesConfig + testAccNamedDataSourceFmcDeviceBFDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceBFDPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name

resource "fmc_bfd_template" "test" {
  name = "fmc_device_bfd_bfd_template"
  hop_type = "SINGLE_HOP"
  echo = "DISABLED"
}

resource "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name = var.interface_name
  logical_name = "outside"
  mode = "NONE"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceBFDConfig() string {
	config := `resource "fmc_device_bfd" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	bfd_template_id = fmc_bfd_template.test.id` + "\n"
	config += `	interface_logical_name = fmc_device_physical_interface.test.logical_name` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	slow_timer = 1000` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bfd" "test" {
			id = fmc_device_bfd.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceBFDConfig() string {
	config := `resource "fmc_device_bfd" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	bfd_template_id = fmc_bfd_template.test.id` + "\n"
	config += `	interface_logical_name = fmc_device_physical_interface.test.logical_name` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	slow_timer = 1000` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bfd" "test" {
			device_id = var.device_id
			interface_logical_name = fmc_device_bfd.test.interface_logical_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
