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

func TestAccDataSourceFmcDeviceHAPairs(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "name", "FTD_HA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "primary_device_id", "<FTD1_ID>"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "secondary_device_id", "<FTD2_ID>"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "is_encryption_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "use_same_link_for_failovers", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "shared_key", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "enc_key_generation_scheme", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_standby_ip", "1.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_active_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_name", "LAN-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_ipv6_addr", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_interface_id", "<interface uuid>"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "lan_failover_interface_type", "PhysicalInterface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_standby_ip", "1.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_active_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_name", "Stateful-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_ipv6_addr", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_interface_id", "<interface uuid>"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "stateful_failover_interface_type", "PhysicalInterface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "action", "SWITCH"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pairs.test", "force_break", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceHAPairsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcDeviceHAPairsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceHAPairsConfig() string {
	config := `resource "fmc_device_ha_pairs" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = "<FTD1_ID>"` + "\n"
	config += `	secondary_device_id = "<FTD2_ID>"` + "\n"
	config += `	is_encryption_enabled = false` + "\n"
	config += `	use_same_link_for_failovers = false` + "\n"
	config += `	shared_key = "cisco123"` + "\n"
	config += `	enc_key_generation_scheme = "CUSTOM"` + "\n"
	config += `	lan_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	lan_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	lan_failover_name = "LAN-INTERFACE"` + "\n"
	config += `	lan_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	lan_failover_ipv6_addr = false` + "\n"
	config += `	lan_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	lan_failover_interface_id = "<interface uuid>"` + "\n"
	config += `	lan_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	stateful_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	stateful_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	stateful_failover_name = "Stateful-INTERFACE"` + "\n"
	config += `	stateful_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	stateful_failover_ipv6_addr = false` + "\n"
	config += `	stateful_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	stateful_failover_interface_id = "<interface uuid>"` + "\n"
	config += `	stateful_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	action = "SWITCH"` + "\n"
	config += `	force_break = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pairs" "test" {
			id = fmc_device_ha_pairs.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceHAPairsConfig() string {
	config := `resource "fmc_device_ha_pairs" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = "<FTD1_ID>"` + "\n"
	config += `	secondary_device_id = "<FTD2_ID>"` + "\n"
	config += `	is_encryption_enabled = false` + "\n"
	config += `	use_same_link_for_failovers = false` + "\n"
	config += `	shared_key = "cisco123"` + "\n"
	config += `	enc_key_generation_scheme = "CUSTOM"` + "\n"
	config += `	lan_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	lan_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	lan_failover_name = "LAN-INTERFACE"` + "\n"
	config += `	lan_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	lan_failover_ipv6_addr = false` + "\n"
	config += `	lan_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	lan_failover_interface_id = "<interface uuid>"` + "\n"
	config += `	lan_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	stateful_failover_standby_ip = "1.1.1.2"` + "\n"
	config += `	stateful_failover_active_ip = "1.1.1.1"` + "\n"
	config += `	stateful_failover_name = "Stateful-INTERFACE"` + "\n"
	config += `	stateful_failover_subnet_mask = "255.255.255.0"` + "\n"
	config += `	stateful_failover_ipv6_addr = false` + "\n"
	config += `	stateful_failover_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	stateful_failover_interface_id = "<interface uuid>"` + "\n"
	config += `	stateful_failover_interface_type = "PhysicalInterface"` + "\n"
	config += `	action = "SWITCH"` + "\n"
	config += `	force_break = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pairs" "test" {
			name = fmc_device_ha_pairs.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
