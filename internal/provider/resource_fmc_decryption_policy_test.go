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

func TestAccFmcDecryptionPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "name", "decryption_policy_test"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "decryption_errors", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "compressed_session", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "ssl_v2_session", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "unknown_cipher_suite", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "unsupported_cipher_suite", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "session_not_cached", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "handshake_errors", "INHERIT_DEFAULT_ACTION"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "policy_action", "DO_NOT_DECRYPT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "event_log_action", "LOG_NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "disallow_untrusted_issuer_resign", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "strip_http3", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "tls13_decryption", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "quic_decryption", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "adaptive_probe", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "block_extensions.0", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "log_end", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "send_events", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "trusted_cas.0.ca_name", "Cisco-Trusted-Authorities"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_decryption_policy.test", "trusted_cas.0.ca_id", "d3856136-b65c-49e6-ac4c-0daf5be98bc1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDecryptionPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDecryptionPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_decryption_policy.test",
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

func testAccFmcDecryptionPolicyConfig_minimum() string {
	config := `resource "fmc_decryption_policy" "test" {` + "\n"
	config += `	name = "decryption_policy_test"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDecryptionPolicyConfig_all() string {
	config := `resource "fmc_decryption_policy" "test" {` + "\n"
	config += `	name = "decryption_policy_test"` + "\n"
	config += `	decryption_errors = "BLOCK"` + "\n"
	config += `	compressed_session = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	ssl_v2_session = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	unknown_cipher_suite = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	unsupported_cipher_suite = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	session_not_cached = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	handshake_errors = "INHERIT_DEFAULT_ACTION"` + "\n"
	config += `	policy_action = "DO_NOT_DECRYPT"` + "\n"
	config += `	event_log_action = "LOG_NONE"` + "\n"
	config += `	disallow_untrusted_issuer_resign = true` + "\n"
	config += `	strip_http3 = false` + "\n"
	config += `	tls13_decryption = false` + "\n"
	config += `	quic_decryption = false` + "\n"
	config += `	adaptive_probe = true` + "\n"
	config += `	block_extensions = [1]` + "\n"
	config += `	log_end = false` + "\n"
	config += `	send_events = false` + "\n"
	config += `	trusted_cas = [{` + "\n"
	config += `		ca_name = "Cisco-Trusted-Authorities"` + "\n"
	config += `		ca_id = "d3856136-b65c-49e6-ac4c-0daf5be98bc1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
