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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccFmcAccessControlPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "description", "My access control policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_send_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "categories.0.name", "cat1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcAccessControlPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcAccessControlPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_access_control_policy.test",
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
func testAccFmcAccessControlPolicyConfig_minimum() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccFmcAccessControlPolicyConfig_all() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My access control policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = true` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_send_syslog = true` + "\n"
	config += `	categories = [{` + "\n"
	config += `	  name = "cat1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  action = "ALLOW"` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  category_name = "cat1"` + "\n"
	config += `	  enabled = true` + "\n"
	config += `	  source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestNewValidAccessControlPolicy(t *testing.T) {

	steps := []resource.TestStep{{
		Config: `resource fmc_access_control_policy step1 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	}, {
		Config: `resource fmc_access_control_policy step2 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	}, {
		Config: `resource fmc_access_control_policy step3 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "catd" },` + "\n" +
			`		{ name = "catm", section = "mandatory" }` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`"catm" must be somewhere above category "catd"`),
	}, {
		Config: `resource fmc_access_control_policy step4 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [{ name = "cat1", section = "" }]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`value must be one of`),
	}, {
		Config: `resource fmc_access_control_policy step5 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "mandatory" },` + "\n" +
			`		{ name = "cat2", section = "mandatory" },` + "\n" +
			`		{ name = "cat3", section = "default"   },` + "\n" +
			`		{ name = "cat4", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat1",      name = "r1", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat3",      name = "r3", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat4",      name = "r4", action = "ALLOW"},` + "\n" +
			`		{ section = "default",         name = "r5", action = "ALLOW"},` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	}, {
		Config: `resource fmc_access_control_policy step6 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "mandatory" },` + "\n" +
			`		{ name = "cat2", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",    name = "step6r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",     name = "step6r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step7 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",      name = "r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step8 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",      name = "r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat1",      name = "r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step9 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{                              name = "step9r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat1",      name = "step9r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step10 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [` + "\n" +
			`		{                              name = "step10r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "step10r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step11 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat1",      name = "step11r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "step11r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step12 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [{ name = "cat1", section = "default" }]` + "\n" +
			`	rules = [{ category_name = "cat1", section = "default", name = "r1", action = "ALLOW"}]` + "\n" +
			`}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
		ExpectError:        regexp.MustCompile(`Cannot use section together with category_name`),
	}}

	resource.Test(t, resource.TestCase{
		// If you see "Step 7/x, expected an error" look above for the name "step7".
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
