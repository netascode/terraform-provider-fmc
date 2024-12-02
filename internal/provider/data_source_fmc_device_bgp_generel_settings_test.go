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

func TestAccDataSourceFmcDeviceBGPGenerelSettings(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "as_number", "65535"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "router_id", "AUTOMATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "scanning_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "as_no_in_path_attribute", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "aggregate_timer", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "default_local_preference", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "keepalive_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "hold_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "min_hold_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_generel_settings.test", "next_hop_delay_interval", ""))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceBGPGenerelSettingsPrerequisitesConfig + testAccDataSourceFmcDeviceBGPGenerelSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceBGPGenerelSettingsPrerequisitesConfig + testAccNamedDataSourceFmcDeviceBGPGenerelSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceBGPGenerelSettingsPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceBGPGenerelSettingsConfig() string {
	config := `resource "fmc_device_bgp_generel_settings" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "65535"` + "\n"
	config += `	router_id = "AUTOMATIC"` + "\n"
	config += `	scanning_interval = ` + "\n"
	config += `	as_no_in_path_attribute = ` + "\n"
	config += `	aggregate_timer = ` + "\n"
	config += `	default_local_preference = ` + "\n"
	config += `	keepalive_interval = ` + "\n"
	config += `	hold_time = ` + "\n"
	config += `	min_hold_time = ` + "\n"
	config += `	next_hop_delay_interval = ` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp_generel_settings" "test" {
			id = fmc_device_bgp_generel_settings.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceBGPGenerelSettingsConfig() string {
	config := `resource "fmc_device_bgp_generel_settings" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "65535"` + "\n"
	config += `	router_id = "AUTOMATIC"` + "\n"
	config += `	scanning_interval = ` + "\n"
	config += `	as_no_in_path_attribute = ` + "\n"
	config += `	aggregate_timer = ` + "\n"
	config += `	default_local_preference = ` + "\n"
	config += `	keepalive_interval = ` + "\n"
	config += `	hold_time = ` + "\n"
	config += `	min_hold_time = ` + "\n"
	config += `	next_hop_delay_interval = ` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp_generel_settings" "test" {
			name = fmc_device_bgp_generel_settings.test.name
			device_id = var.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
