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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcDeviceBGPGenerelSettings(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp_generel_settings.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp_generel_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "as_number", "65535"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "router_id", "AUTOMATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "scanning_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "as_no_in_path_attribute", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "aggregate_timer", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "default_local_preference", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "keepalive_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "hold_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "min_hold_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp_generel_settings.test", "next_hop_delay_interval", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceBGPGenerelSettingsPrerequisitesConfig + testAccFmcDeviceBGPGenerelSettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceBGPGenerelSettingsPrerequisitesConfig + testAccFmcDeviceBGPGenerelSettingsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcDeviceBGPGenerelSettingsPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceBGPGenerelSettingsConfig_minimum() string {
	config := `resource "fmc_device_bgp_generel_settings" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "65535"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceBGPGenerelSettingsConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll