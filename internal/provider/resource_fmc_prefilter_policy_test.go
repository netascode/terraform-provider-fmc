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

func TestAccFmcPrefilterPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "description", "My prefilter policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action", "BLOCK_TUNNELS"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_log_end", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.action", "FASTPATH"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.rule_type", "PREFILTER"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_network_literals.0.value", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_interfaces.0.id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_interfaces.0.type", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_interfaces.0.id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_interfaces.0.type", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.tunnel_zone.0.id", "0050568A-7F57-0ed3-0000-004294975576"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.encapsulation_ports_gre", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.encapsulation_ports_in_in_ip", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.encapsulation_ports_ipv6_in_ip", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.encapsulation_ports_teredo", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPrefilterPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPrefilterPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_prefilter_policy.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPrefilterPolicyConfig_minimum() string {
	config := `resource "fmc_prefilter_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPrefilterPolicyConfig_all() string {
	config := `resource "fmc_prefilter_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My prefilter policy"` + "\n"
	config += `	default_action = "BLOCK_TUNNELS"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = false` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	rules = [{` + "\n"
	config += `		name = "rule1"` + "\n"
	config += `		action = "FASTPATH"` + "\n"
	config += `		rule_type = "PREFILTER"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		bidirectional = true` + "\n"
	config += `		log_begin = true` + "\n"
	config += `		log_end = true` + "\n"
	config += `		send_events_to_fmc = true` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `			type = fmc_network.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `			type = fmc_host.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		source_interfaces = [{` + "\n"
	config += `			id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `			type = "ROUTED"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_interfaces = [{` + "\n"
	config += `			id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `			type = "ROUTED"` + "\n"
	config += `		}]` + "\n"
	config += `		tunnel_zone = [{` + "\n"
	config += `			id = "0050568A-7F57-0ed3-0000-004294975576"` + "\n"
	config += `		}]` + "\n"
	config += `		encapsulation_ports_gre = false` + "\n"
	config += `		encapsulation_ports_in_in_ip = false` + "\n"
	config += `		encapsulation_ports_ipv6_in_ip = false` + "\n"
	config += `		encapsulation_ports_teredo = false` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
