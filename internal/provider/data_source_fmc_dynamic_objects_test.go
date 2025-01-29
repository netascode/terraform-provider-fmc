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

func TestAccDataSourceFmcDynamicObjects(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_dynamic_objects.test", "items.dynamic_object_1.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_dynamic_objects.test", "items.dynamic_object_1.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dynamic_objects.test", "items.dynamic_object_1.description", "My Dynamic Object 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_dynamic_objects.test", "items.dynamic_object_1.object_type", "IP"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDynamicObjectsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDynamicObjectsConfig() string {
	config := `resource "fmc_dynamic_objects" "test" {` + "\n"
	config += `	items = { "dynamic_object_1" = {` + "\n"
	config += `		description = "My Dynamic Object 1"` + "\n"
	config += `		object_type = "IP"` + "\n"
	config += `		mappings = ["10.0.0.1"]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_dynamic_objects" "test" {
			depends_on = [fmc_dynamic_objects.test]
			items = {
				"dynamic_object_1" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
