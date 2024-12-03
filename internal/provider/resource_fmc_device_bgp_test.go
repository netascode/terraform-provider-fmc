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
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "ipv4_address_family_type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_learned_route_map_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_external_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_internal_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_local_distance", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_forward_packets_over_multipath_ibgp", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_forward_packets_over_multipath_ebgp", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_generate_default_route_map", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_routes_advertise_map", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_routes_advertise_exist_nonexist_map", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_authentication_password", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_send_community_attribute", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_nexthop_self", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_disable_connection_verification", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_tcp_mtu_path_discovery", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_max_hop_count", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_tcp_transport_mode", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_weight", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_version", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_customized_no_prepend", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_customized_replace_as", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.ipv4_neighbor_customized_accept_both_as", ""))

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
	config += `	ipv4_learned_route_map_id = ""` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceBGPConfig_all() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ipv4_learned_route_map_id = ""` + "\n"
	config += `	ipv4_external_distance = ` + "\n"
	config += `	ipv4_internal_distance = ` + "\n"
	config += `	ipv4_local_distance = ` + "\n"
	config += `	ipv4_forward_packets_over_multipath_ibgp = ` + "\n"
	config += `	ipv4_forward_packets_over_multipath_ebgp = ` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `		ipv4_neighbor_filter_access_lists = [{` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_neighbor_filter_route_map_lists = [{` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_neighbor_filter_prefix_lists = [{` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_neighbor_filter_as_path_lists = [{` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_neighbor_generate_default_route_map = ""` + "\n"
	config += `		ipv4_neighbor_routes_advertise_map = ""` + "\n"
	config += `		ipv4_neighbor_routes_advertise_exist_nonexist_map = ""` + "\n"
	config += `		ipv4_neighbor_authentication_password = ""` + "\n"
	config += `		ipv4_neighbor_send_community_attribute = ` + "\n"
	config += `		ipv4_neighbor_nexthop_self = ` + "\n"
	config += `		ipv4_neighbor_disable_connection_verification = ` + "\n"
	config += `		ipv4_neighbor_tcp_mtu_path_discovery = ` + "\n"
	config += `		ipv4_neighbor_max_hop_count = ` + "\n"
	config += `		ipv4_neighbor_tcp_transport_mode = ` + "\n"
	config += `		ipv4_neighbor_weight = ` + "\n"
	config += `		ipv4_neighbor_version = ""` + "\n"
	config += `		ipv4_neighbor_customized_no_prepend = ` + "\n"
	config += `		ipv4_neighbor_customized_replace_as = ` + "\n"
	config += `		ipv4_neighbor_customized_accept_both_as = ` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_aggregate_addresses = [{` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_filterings = [{` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_networks = [{` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_redistributions = [{` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_route_injections = [{` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
