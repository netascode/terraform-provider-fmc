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

func TestAccFmcDeviceVTEPPolicy(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_vtep_policy.test", "vteps.0.nve_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_vtep_policy.test", "vteps.0.neighbor_discovery", "STATIC_PEER_IP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_vtep_policy.test", "vteps.0.neighbor_address_literal", "192.168.0.1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceVTEPPolicyPrerequisitesConfig + testAccFmcDeviceVTEPPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceVTEPPolicyPrerequisitesConfig + testAccFmcDeviceVTEPPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcDeviceVTEPPolicyPrerequisitesConfig = `
resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = "GigabitEthernet0/1"
  mode         = "NONE"
  logical_name = "myinterface-0-1"
  mtu          = 9000
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceVTEPPolicyConfig_minimum() string {
	config := `resource "fmc_device_vtep_policy" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	nve_enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceVTEPPolicyConfig_all() string {
	config := `resource "fmc_device_vtep_policy" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	nve_enabled = true` + "\n"
	config += `	vteps = [{` + "\n"
	config += `		source_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		nve_number = 1` + "\n"
	config += `		neighbor_discovery = "STATIC_PEER_IP"` + "\n"
	config += `		neighbor_address_literal = "192.168.0.1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
