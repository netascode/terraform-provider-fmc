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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcFilePolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "name", "file_policy_1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "description", "My file policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "archive_depth", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "archive_depth_action", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "block_encrypted_archives", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "clean_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "custom_detection_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "first_time_file_analysis", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "inspect_archives", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "threat_score", "High"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcFilePolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFilePolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_file_policy.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcFilePolicyConfig_minimum() string {
	config := `resource "fmc_file_policy" "test" {` + "\n"
	config += `	name = "file_policy_1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFilePolicyConfig_all() string {
	config := `resource "fmc_file_policy" "test" {` + "\n"
	config += `	name = "file_policy_1"` + "\n"
	config += `	description = "My file policy"` + "\n"
	config += `	archive_depth = "3"` + "\n"
	config += `	archive_depth_action = true` + "\n"
	config += `	block_encrypted_archives = false` + "\n"
	config += `	clean_list = true` + "\n"
	config += `	custom_detection_list = true` + "\n"
	config += `	first_time_file_analysis = true` + "\n"
	config += `	inspect_archives = false` + "\n"
	config += `	threat_score = "High"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
