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

func TestAccFmcPolicyAssignments(t *testing.T) {
	if os.Getenv("TF_VAR_target_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_target_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_policy_assignments.test", "targets.0.type", "DeviceGroup"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPolicyAssignmentsPrerequisitesConfig + testAccFmcPolicyAssignmentsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPolicyAssignmentsPrerequisitesConfig + testAccFmcPolicyAssignmentsConfig_all(),
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

const testAccFmcPolicyAssignmentsPrerequisitesConfig = `
resource "fmc_access_control_policy" "example" {
    categories                        = []
    default_action                    = "BLOCK"
    default_action_log_begin          = false
    default_action_log_end            = false
    default_action_send_events_to_fmc = false
    default_action_send_syslog        = false
    name                              = "policy-example-test"
    rules                             = []
}

variable "target_id" { default = null } // tests will set $TF_VAR_target_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPolicyAssignmentsConfig_minimum() string {
	config := `resource "fmc_policy_assignments" "test" {` + "\n"
	config += `	policy_id = fmc_access_control_policy.example.id` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = var.target_id` + "\n"
	config += `		type = "DeviceGroup"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPolicyAssignmentsConfig_all() string {
	config := `resource "fmc_policy_assignments" "test" {` + "\n"
	config += `	policy_id = fmc_access_control_policy.example.id` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = var.target_id` + "\n"
	config += `		type = "DeviceGroup"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
