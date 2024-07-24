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

func TestAccFmcFileRules(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "action", "DETECT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "protocol", "ANY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "direction", "ANY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "file_categories.0.id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "file_categories.0.name", "Archive"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "file_types.0.id", "19"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_rules.test", "file_types.0.name", "7Z"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcFileRulesPrerequisitesConfig + testAccFmcFileRulesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFileRulesPrerequisitesConfig + testAccFmcFileRulesConfig_all(),
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

const testAccFmcFileRulesPrerequisitesConfig = `
resource "fmc_file_policy" "test" {
  name                     = "file_rule_test"
  description              = "My file policy Test"
  archive_depth            = "3"
  archive_depth_action     = true
  block_encrypted_archives = false
  clean_list               = true
  custom_detection_list    = true
  first_time_file_analysis = true
  inspect_archives         = false
  threat_score             = "High"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcFileRulesConfig_minimum() string {
	config := `resource "fmc_file_rules" "test" {` + "\n"
	config += `	file_policy_id = fmc_file_policy.test.id` + "\n"
	config += `	action = "DETECT"` + "\n"
	config += `	protocol = "ANY"` + "\n"
	config += `	direction = "ANY"` + "\n"
	config += `	file_categories = [{` + "\n"
	config += `		id = "2"` + "\n"
	config += `		name = "Archive"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFileRulesConfig_all() string {
	config := `resource "fmc_file_rules" "test" {` + "\n"
	config += `	file_policy_id = fmc_file_policy.test.id` + "\n"
	config += `	action = "DETECT"` + "\n"
	config += `	protocol = "ANY"` + "\n"
	config += `	direction = "ANY"` + "\n"
	config += `	file_categories = [{` + "\n"
	config += `		id = "2"` + "\n"
	config += `		name = "Archive"` + "\n"
	config += `	}]` + "\n"
	config += `	file_types = [{` + "\n"
	config += `		id = "19"` + "\n"
	config += `		name = "7Z"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll