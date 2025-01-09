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

func TestAccDataSourceFmcHosts(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_hosts.test", "items.hosts_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_hosts.test", "items.hosts_1.description", "My Host 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_hosts.test", "items.hosts_1.overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_hosts.test", "items.hosts_1.ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_hosts.test", "items.hosts_1.type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcHostsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcHostsConfig() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = { "hosts_1" = {` + "\n"
	config += `		description = "My Host 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		ip = "10.1.1.1"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_hosts" "test" {
			depends_on = [fmc_hosts.test]
			items = {
				"hosts_1" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
