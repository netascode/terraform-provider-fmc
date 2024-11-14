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
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.action", "FASTPATH"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.rule_type", "PREFILTER"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.send_events_to_fmc", "true"))

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
	config += `		action = "FASTPATH"` + "\n"
	config += `		rule_type = "PREFILTER"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		bidirectional = true` + "\n"
	config += `		log_begin = true` + "\n"
	config += `		log_end = true` + "\n"
	config += `		send_events_to_fmc = true` + "\n"
	config += `		description = ""` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
