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

func TestAccFmcPorts(t *testing.T) {
	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.ports_1.id"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.port", "1234"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.description", "port 1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_2.port", "1235"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_3.port", "1236"))

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_ports.test", "items.ports_1.id"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.port", "1233"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_1.description", "port 1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_2.port", "1235"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_ports.test", "items.ports_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_ports.test", "items.ports_4.port", "1239"))

	var steps []resource.TestStep

	steps = append(steps, resource.TestStep{
		Config: testAccFmcPortsConfig_step01(),
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPortsConfig_step02(),
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

func testAccFmcPortsConfig_step01() string {
	config := `resource "fmc_ports" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"ports_1" = {` + "\n"
	config += `			port = "1234",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `			description = "port 1"` + "\n"
	config += `			overridable = true` + "\n"
	config += `		},` + "\n"
	config += `		"ports_2" = {` + "\n"
	config += `			port = "1235",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `		},` + "\n"
	config += `		"ports_3" = {` + "\n"
	config += `			port = "1236",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}

func testAccFmcPortsConfig_step02() string {
	config := `resource "fmc_ports" "test" {` + "\n"
	config += `	items = {` + "\n"
	config += `		"ports_1" = {` + "\n"
	config += `			port = "1233",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `			description = "port 1 new description"` + "\n"
	config += `		},` + "\n"
	config += `		"ports_2" = {` + "\n"
	config += `			port = "1235",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `		},` + "\n"
	config += `		"ports_4" = {` + "\n"
	config += `			port = "1239",` + "\n"
	config += `			protocol = "TCP",` + "\n"
	config += `		},` + "\n"
	config += `	} ` + "\n"
	config += `}` + "\n"
	return config
}
