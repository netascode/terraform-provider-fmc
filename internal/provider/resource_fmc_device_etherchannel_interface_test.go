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

func TestAccFmcDeviceEtherChannelInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" && os.Getenv("TF_VAR_physical_interface_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id or TF_VAR_physical_interface_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "logical_name", "myinterface-0-1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mode", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ether_channel_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "ipv6_enable_ra", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_etherchannel_interface.test", "auto_negotiation", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccFmcDeviceEtherChannelInterfaceConfig_all(),
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

const testAccFmcDeviceEtherChannelInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "physical_interface_id" { default = null } // tests will set $TF_VAR_physical_interface_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceEtherChannelInterfaceConfig_minimum() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "iface_minimum"` + "\n"
	config += `	management_only = true` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceEtherChannelInterfaceConfig_all() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "myinterface-0-1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	management_only = false` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	name = ""` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `	nve_only = false` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_enable_ra = ` + "\n"
	config += `	auto_negotiation = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll