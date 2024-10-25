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

func TestAccDataSourceFmcPolicyAssignments(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "name", "AccessPolicyName"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "type", "AccessPolicy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "policy_id", "0050568A-2561-0ed3-0000-004294972836"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "targets.0.id", "9862719c-8d5f-11ef-99a6-aef0794da1c1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "targets.0.type", "DeviceGroup"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignments.test", "targets.0.name", "FTD_Device1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcPolicyAssignmentsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcPolicyAssignmentsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcPolicyAssignmentsConfig() string {
	config := `resource "fmc_policy_assignments" "test" {` + "\n"
	config += `	name = "AccessPolicyName"` + "\n"
	config += `	type = "AccessPolicy"` + "\n"
	config += `	policy_id = "0050568A-2561-0ed3-0000-004294972836"` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = "9862719c-8d5f-11ef-99a6-aef0794da1c1"` + "\n"
	config += `		type = "DeviceGroup"` + "\n"
	config += `		name = "FTD_Device1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_assignments" "test" {
			id = fmc_policy_assignments.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcPolicyAssignmentsConfig() string {
	config := `resource "fmc_policy_assignments" "test" {` + "\n"
	config += `	name = "AccessPolicyName"` + "\n"
	config += `	type = "AccessPolicy"` + "\n"
	config += `	policy_id = "0050568A-2561-0ed3-0000-004294972836"` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = "9862719c-8d5f-11ef-99a6-aef0794da1c1"` + "\n"
	config += `		type = "DeviceGroup"` + "\n"
	config += `		name = "FTD_Device1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_assignments" "test" {
			name = fmc_policy_assignments.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
