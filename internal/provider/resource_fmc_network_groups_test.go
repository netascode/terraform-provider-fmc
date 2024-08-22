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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

func TestAccFmcNetworkGroups(t *testing.T) {
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcNetworkGroupsPrerequisitesConfig + testAccFmcNetworkGroupsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworkGroupsPrerequisitesConfig + testAccFmcNetworkGroupsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	// Do not test Import here, because Import is not yet implemented.

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcNetworkGroupsConfig_minimum() string {
	config := `resource "fmc_network_groups" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworkGroupsConfig_all() string {
	config := `resource "fmc_network_groups" "test" {` + "\n"
	config += `	items = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcNetworkGroups_GroupNames(t *testing.T) {
	steps := []resource.TestStep{{
		// step 1
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		group_names = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 2
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 3
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.99.0.0/16"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 4
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		group_names = ["g2","g3","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g3" = {` + "\n" +
			`		group_names = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g4" = {` + "\n" +
			`		group_names = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g5" = {` + "\n" +
			`		group_names = ["g3", "g4"]` + "\n" +
			`	}` + "\n" +
			`	"g6" = {` + "\n" +
			`		group_names = ["g2","g3","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 5
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		group_names = ["g2","g4","g5"]` + "\n" +
			`	}` + "\n" +
			`	"g2" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`	"g4" = {` + "\n" +
			`		group_names = ["g2"]` + "\n" +
			`	}` + "\n" +
			`	"g5" = {` + "\n" +
			`		group_names = ["g4"]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 6
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g1" = {` + "\n" +
			`		literals = [{value = "10.0.0.0/8"}]` + "\n" +
			`	}` + "\n" +
			`}}`,
	}, {
		// step 7
		Config: `resource fmc_network_groups test { items = {` + "\n" +
			`	"g2" = {` + "\n" +
			`		group_names = ["g2"]` + "\n" +
			`	}` + "\n" +
			`}}`,
		ExpectError: regexp.MustCompile(`Cycle in group_names`),
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
