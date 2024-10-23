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

func TestAccFmcNetworks(t *testing.T) {
	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.prefix", "10.20.30.0/24"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.description", "network 1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_2.prefix", "10.20.40.0/24"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_3.prefix", "10.20.50.0/24"))

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.prefix", "10.20.30.0/24"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_1.description", "network 1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_2.prefix", "10.20.40.0/24"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_networks.test", "items.networks_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_networks.test", "items.networks_4.prefix", "10.20.60.0/24"))

	var steps []resource.TestStep

	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworksConfig_step01(),
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworksConfig_step02(),
		Check:  resource.ComposeTestCheckFunc(checks_step02...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcNetworksConfig_minimum() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworksConfig_all() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = { "networks_1" = {` + "\n"
	config += `		description = "My Network 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		prefix = "10.1.1.1"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func testAccFmcNetworksConfig_step01() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"networks_1" = {` + "\n"
	config += `			prefix = "10.20.30.0/24"` + "\n"
	config += `			description = "network 1"` + "\n"
	config += `			overridable = true` + "\n"
	config += `		},` + "\n"
	config += `		"networks_2" = {` + "\n"
	config += `			prefix = "10.20.40.0/24"` + "\n"
	config += `		},` + "\n"
	config += `		"networks_3" = {` + "\n"
	config += `			prefix = "10.20.50.0/24"` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}

func testAccFmcNetworksConfig_step02() string {
	config := `resource "fmc_networks" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"networks_1" = {` + "\n"
	config += `			prefix = "10.20.30.0/24"` + "\n"
	config += `			description = "network 1 new description"` + "\n"
	config += `		},` + "\n"
	config += `		"networks_2" = {` + "\n"
	config += `			prefix = "10.20.40.0/24"` + "\n"
	config += `		},` + "\n"
	config += `		"networks_4" = {` + "\n"
	config += `			prefix = "10.20.60.0/24"` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}
