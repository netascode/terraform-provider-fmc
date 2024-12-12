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

func TestAccDataSourceFmcSGTs(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sgts.test", "items.sgt_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sgts.test", "items.sgt_1.name", "SGT1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sgts.test", "items.sgt_1.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sgts.test", "items.sgt_1.description", "My SGT object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sgts.test", "items.sgt_1.tag", "11"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSGTsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSGTsConfig() string {
	config := `resource "fmc_sgts" "test" {` + "\n"
	config += `	items = { "sgt_1" = {` + "\n"
	config += `		name = "SGT1"` + "\n"
	config += `		description = "My SGT object"` + "\n"
	config += `		tag = "11"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_sgts" "test" {
			depends_on = [fmc_sgts.test]
			items = {
				"sgt_1" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
