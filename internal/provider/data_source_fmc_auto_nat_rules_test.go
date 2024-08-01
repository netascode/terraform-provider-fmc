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

func TestAccDataSourceFmcAutoNatRules(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_auto_nat_rules.test", "nat_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_auto_nat_rules.test", "interface_ipv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_auto_nat_rules.test", "fall_through", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAutoNatRulesPrerequisitesConfig + testAccDataSourceFmcAutoNatRulesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcAutoNatRulesPrerequisitesConfig = `
resource "fmc_nat_policy" "_nat_policy_test" {
  name = "nat_rule_test"
  description = "My nat policy Test"
}

resource "fmc_network" "_original_network_test" {
  name        = "NET1"
  description = "My network object"
  prefix      = "10.1.2.0/24"
}

resource "fmc_network" "_translated_network_test" {
  name        = "NET2"
  description = "My network object"
  prefix      = "20.1.2.0/24"
}

resource "fmc_security_zone" "_src_zone_test" {
  name           = "SecZone1"
  interface_mode = "ROUTED"
}

resource "fmc_security_zone" "_dest_zone_test" {
  name           = "SecZone2"
  interface_mode = "ROUTED"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcAutoNatRulesConfig() string {
	config := `resource "fmc_auto_nat_rules" "test" {` + "\n"
	config += `	nat_policy_id = fmc_nat_policy._nat_policy_test.id` + "\n"
	config += `	nat_type = "STATIC"` + "\n"
	config += `	interface_ipv6 = false` + "\n"
	config += `	fall_through = false` + "\n"
	config += `	original_network_id = fmc_network._original_network_test.id` + "\n"
	config += `	translated_network_id = fmc_network._translated_network_test.id` + "\n"
	config += `	source_interface_id = fmc_security_zone._src_zone_test.id` + "\n"
	config += `	destination_interface_id = fmc_security_zone._dest_zone_test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_auto_nat_rules" "test" {
			id = fmc_auto_nat_rules.test.id
			nat_policy_id = fmc_nat_policy._nat_policy_test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
