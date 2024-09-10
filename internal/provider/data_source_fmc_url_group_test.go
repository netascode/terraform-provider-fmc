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

func TestAccDataSourceFmcUrlGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "name", "url_group_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "objects.0.id", "0050568A-4E02-1ed3-0000-004294969198"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "objects.0.name", "url_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_url_group.test", "description", "My URL group"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcUrlGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcUrlGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcUrlGroupConfig() string {
	config := `resource "fmc_url_group" "test" {` + "\n"
	config += `	name = "url_group_1"` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = "0050568A-4E02-1ed3-0000-004294969198"` + "\n"
	config += `		name = "url_1"` + "\n"
	config += `	}]` + "\n"
	config += `	description = "My URL group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url_group" "test" {
			id = fmc_url_group.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcUrlGroupConfig() string {
	config := `resource "fmc_url_group" "test" {` + "\n"
	config += `	name = "url_group_1"` + "\n"
	config += `	objects = [{` + "\n"
	config += `		id = "0050568A-4E02-1ed3-0000-004294969198"` + "\n"
	config += `		name = "url_1"` + "\n"
	config += `	}]` + "\n"
	config += `	description = "My URL group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_url_group" "test" {
			name = fmc_url_group.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
