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

func TestAccDataSourceFmcManualNatRules(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "nat_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "interface_ipv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "fall_through", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "unidirectional", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_manual_nat_rules.test", "description", "description of nat rule"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcManualNatRulesPrerequisitesConfig + testAccDataSourceFmcManualNatRulesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcManualNatRulesPrerequisitesConfig = `
resource "fmc_nat_policy" "_nat_policy_test" {
  name = "nat_rule_test"
  description = "My nat policy Test"
}

resource "fmc_network" "_original_src_test" {
  name        = "NET1"
  description = "My network object"
  prefix      = "10.1.2.0/24"
}

resource "fmc_network" "_translated_src_test" {
  name        = "NET2"
  description = "My network object"
  prefix      = "20.1.2.0/24"
}

resource "fmc_network" "_original_destination_test" {
  name        = "NET3"
  description = "My network object"
  prefix      = "30.1.2.0/24"
}

resource "fmc_network" "_original_destination_test_2" {
  name        = "NET4"
  description = "My network object"
  prefix      = "40.1.2.0/24"
}

resource "fmc_security_zone" "_src_zone_test" {
  name           = "SecZone1"
  interface_mode = "ROUTED"
}

resource "fmc_security_zone" "_dest_zone_test" {
  name           = "SecZone2"
  interface_mode = "ROUTED"
}

resource "fmc_port" "_translated_src_port" {
  port        = "6012"
  name        = "tcp6012"
  protocol    = "TCP"
  description = "Port TCP/6012 (HTTPS)"
}

resource "fmc_port" "_translated_destination_port" {
  port        = "6013"
  name        = "tcp6013"
  protocol    = "TCP"
  description = "Port TCP/6012 (HTTPS)"
}

resource "fmc_port" "_original_src_port" {
  port        = "6014"
  name        = "tcp6014"
  protocol    = "TCP"
  description = "Port TCP/6012 (HTTPS)"
}

resource "fmc_port" "_original_destination_port" {
  port        = "6015"
  name        = "tcp6015"
  protocol    = "TCP"
  description = "Port TCP/6012 (HTTPS)"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcManualNatRulesConfig() string {
	config := `resource "fmc_manual_nat_rules" "test" {` + "\n"
	config += `	nat_policy_id = fmc_nat_policy._nat_policy_test.id` + "\n"
	config += `	nat_type = "STATIC"` + "\n"
	config += `	interface_ipv6 = false` + "\n"
	config += `	fall_through = false` + "\n"
	config += `	enabled = true` + "\n"
	config += `	unidirectional = true` + "\n"
	config += `	original_source_id = fmc_network._original_src_test.id` + "\n"
	config += `	original_destination_id = fmc_network._original_destination_test.id` + "\n"
	config += `	original_source_port_id = fmc_port._original_src_port.id` + "\n"
	config += `	original_destination_port_id = fmc_port._original_destination_port.id` + "\n"
	config += `	translated_destination_id = fmc_network._original_destination_test.id` + "\n"
	config += `	translated_source_id = fmc_network._translated_src_test.id` + "\n"
	config += `	translated_destination_port_id = fmc_port._translated_destination_port.id` + "\n"
	config += `	translated_source_port_id = fmc_port._translated_src_port.id` + "\n"
	config += `	source_interface_id = fmc_security_zone._src_zone_test.id` + "\n"
	config += `	destination_interface_id = fmc_security_zone._dest_zone_test.id` + "\n"
	config += `	description = "description of nat rule"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_manual_nat_rules" "test" {
			id = fmc_manual_nat_rules.test.id
			nat_policy_id = fmc_nat_policy._nat_policy_test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
