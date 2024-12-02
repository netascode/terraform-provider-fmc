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

func TestAccFmcDeviceBGP(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "as_number"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "address_family_type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "learned_route_map_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "external_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "internal_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "local_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "forward_packets_over_multipath_ibgp", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "forward_packets_over_multipath_ebgp", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "maximum_paths.0.value", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceBGPPrerequisitesConfig + testAccFmcDeviceBGPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceBGPPrerequisitesConfig + testAccFmcDeviceBGPConfig_all(),
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

const testAccFmcDeviceBGPPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceBGPConfig_minimum() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	learned_route_map_id = ""` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceBGPConfig_all() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	learned_route_map_id = ""` + "\n"
	config += `	external_distance = ` + "\n"
	config += `	internal_distance = ` + "\n"
	config += `	local_distance = ` + "\n"
	config += `	forward_packets_over_multipath_ibgp = ` + "\n"
	config += `	forward_packets_over_multipath_ebgp = ` + "\n"
	config += `	neighbors = [{` + "\n"
	config += `	}]` + "\n"
	config += `	maximum_paths = [{` + "\n"
	config += `		value = ` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
