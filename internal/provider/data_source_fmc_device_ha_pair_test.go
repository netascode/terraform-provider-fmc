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

func TestAccDataSourceFmcDeviceHAPair(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_device_2_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_device_2_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "name", "FTD_HA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "is_encryption_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "shared_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_standby_ip", "1.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_active_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_name", "LAN-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_ipv6_addr", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "lan_failover_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_standby_ip", "10.10.10.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_active_ip", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_name", "Stateful-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_ipv6_addr", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_interface_id", "76d24097-hj7r-7786-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair.test", "stateful_failover_interface_type", "PhysicalInterface"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceHAPairPrerequisitesConfig + testAccDataSourceFmcDeviceHAPairConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceHAPairPrerequisitesConfig + testAccNamedDataSourceFmcDeviceHAPairConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceHAPairPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceHAPairConfig() string {
	config := `resource "fmc_device_ha_pair" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = var.device_id` + "\n"
	config += `	secondary_device_id = var.device_2_id` + "\n"
	config += `	is_encryption_enabled = false` + "\n"
	config += `	use_same_link_for_failovers = false` + "\n"
	config += `	shared_key = "cisco123"` + "\n"
	config += `	enc_key_generation_scheme = "CUSTOM"` + "\n"
	config += `	lan_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	lan_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	lan_failover_name = "LAN-INTERFACE"` + "\n"
	config += `	lan_failover_netmask = "255.255.255.0"` + "\n"
	config += `	lan_failover_ipv6_addr = false` + "\n"
	config += `	lan_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	lan_failover_interface_id = "757kdgh5-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	lan_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	stateful_failover_standby_ip = "10.10.10.2"` + "\n"
	config += `	stateful_failover_active_ip = "10.10.10.1"` + "\n"
	config += `	stateful_failover_name = "Stateful-INTERFACE"` + "\n"
	config += `	stateful_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	stateful_failover_ipv6_addr = false` + "\n"
	config += `	stateful_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	stateful_failover_interface_id = "76d24097-hj7r-7786-a4d0-a8c07ac08470"` + "\n"
	config += `	stateful_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair" "test" {
			id = fmc_device_ha_pair.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceHAPairConfig() string {
	config := `resource "fmc_device_ha_pair" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = var.device_id` + "\n"
	config += `	secondary_device_id = var.device_2_id` + "\n"
	config += `	is_encryption_enabled = false` + "\n"
	config += `	use_same_link_for_failovers = false` + "\n"
	config += `	shared_key = "cisco123"` + "\n"
	config += `	enc_key_generation_scheme = "CUSTOM"` + "\n"
	config += `	lan_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	lan_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	lan_failover_name = "LAN-INTERFACE"` + "\n"
	config += `	lan_failover_netmask = "255.255.255.0"` + "\n"
	config += `	lan_failover_ipv6_addr = false` + "\n"
	config += `	lan_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	lan_failover_interface_id = "757kdgh5-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	lan_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	stateful_failover_standby_ip = "10.10.10.2"` + "\n"
	config += `	stateful_failover_active_ip = "10.10.10.1"` + "\n"
	config += `	stateful_failover_name = "Stateful-INTERFACE"` + "\n"
	config += `	stateful_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	stateful_failover_ipv6_addr = false` + "\n"
	config += `	stateful_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	stateful_failover_interface_id = "76d24097-hj7r-7786-a4d0-a8c07ac08470"` + "\n"
	config += `	stateful_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair" "test" {
			name = fmc_device_ha_pair.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
