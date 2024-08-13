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
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DecryptionPolicy struct {
	Id                            types.String                 `tfsdk:"id"`
	Domain                        types.String                 `tfsdk:"domain"`
	Name                          types.String                 `tfsdk:"name"`
	DecryptionErrors              types.String                 `tfsdk:"decryption_errors"`
	CompressedSession             types.String                 `tfsdk:"compressed_session"`
	SslV2Session                  types.String                 `tfsdk:"ssl_v2_session"`
	UnknownCipherSuite            types.String                 `tfsdk:"unknown_cipher_suite"`
	UnsupportedCipherSuite        types.String                 `tfsdk:"unsupported_cipher_suite"`
	SessionNotCached              types.String                 `tfsdk:"session_not_cached"`
	HandshakeErrors               types.String                 `tfsdk:"handshake_errors"`
	PolicyAction                  types.String                 `tfsdk:"policy_action"`
	EventLogAction                types.String                 `tfsdk:"event_log_action"`
	DisallowUntrustedIssuerResign types.Bool                   `tfsdk:"disallow_untrusted_issuer_resign"`
	StripHttp3                    types.Bool                   `tfsdk:"strip_http3"`
	Tls13Decryption               types.Bool                   `tfsdk:"tls13_decryption"`
	QuicDecryption                types.Bool                   `tfsdk:"quic_decryption"`
	AdaptiveProbe                 types.Bool                   `tfsdk:"adaptive_probe"`
	BlockExtensions               types.List                   `tfsdk:"block_extensions"`
	LogEnd                        types.Bool                   `tfsdk:"log_end"`
	SendEvents                    types.Bool                   `tfsdk:"send_events"`
	TrustedCas                    []DecryptionPolicyTrustedCas `tfsdk:"trusted_cas"`
}

type DecryptionPolicyTrustedCas struct {
	CaName types.String `tfsdk:"ca_name"`
	CaId   types.String `tfsdk:"ca_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DecryptionPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/decryptionpolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DecryptionPolicy) toBody(ctx context.Context, state DecryptionPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.DecryptionErrors.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.decryptionErrors", data.DecryptionErrors.ValueString())
	}
	if !data.CompressedSession.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.compressedSession", data.CompressedSession.ValueString())
	}
	if !data.SslV2Session.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.sslV2Session", data.SslV2Session.ValueString())
	}
	if !data.UnknownCipherSuite.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.unknownCipherSuite", data.UnknownCipherSuite.ValueString())
	}
	if !data.UnsupportedCipherSuite.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.unsupportedCipherSuite", data.UnsupportedCipherSuite.ValueString())
	}
	if !data.SessionNotCached.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.sessionNotCached", data.SessionNotCached.ValueString())
	}
	if !data.HandshakeErrors.IsNull() {
		body, _ = sjson.Set(body, "undecryptableActions.handshakeErrors", data.HandshakeErrors.ValueString())
	}
	if !data.PolicyAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.policyAction", data.PolicyAction.ValueString())
	}
	if !data.EventLogAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.eventLogAction", data.EventLogAction.ValueString())
	}
	if !data.DisallowUntrustedIssuerResign.IsNull() {
		body, _ = sjson.Set(body, "advancedOptions.disallowUntrustedIssuerResign", data.DisallowUntrustedIssuerResign.ValueBool())
	}
	if !data.StripHttp3.IsNull() {
		body, _ = sjson.Set(body, "advancedOptions.stripHTTP3", data.StripHttp3.ValueBool())
	}
	if !data.Tls13Decryption.IsNull() {
		body, _ = sjson.Set(body, "advancedOptions.tls13Decryption", data.Tls13Decryption.ValueBool())
	}
	if !data.QuicDecryption.IsNull() {
		body, _ = sjson.Set(body, "advancedOptions.quicDecryption", data.QuicDecryption.ValueBool())
	}
	if !data.AdaptiveProbe.IsNull() {
		body, _ = sjson.Set(body, "advancedOptions.adaptiveProbe", data.AdaptiveProbe.ValueBool())
	}
	if !data.BlockExtensions.IsNull() {
		var values []int64
		data.BlockExtensions.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "advancedOptions.blockExtensions", values)
	}
	if !data.LogEnd.IsNull() {
		body, _ = sjson.Set(body, "logging.logEnd", data.LogEnd.ValueBool())
	}
	if !data.SendEvents.IsNull() {
		body, _ = sjson.Set(body, "logging.sendEvents", data.SendEvents.ValueBool())
	}
	if len(data.TrustedCas) > 0 {
		body, _ = sjson.Set(body, "trustedCAs.objects", []interface{}{})
		for _, item := range data.TrustedCas {
			itemBody := ""
			if !item.CaName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.CaName.ValueString())
			}
			if !item.CaId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.CaId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "trustedCAs.objects.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DecryptionPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("undecryptableActions.decryptionErrors"); value.Exists() {
		data.DecryptionErrors = types.StringValue(value.String())
	} else {
		data.DecryptionErrors = types.StringNull()
	}
	if value := res.Get("undecryptableActions.compressedSession"); value.Exists() {
		data.CompressedSession = types.StringValue(value.String())
	} else {
		data.CompressedSession = types.StringNull()
	}
	if value := res.Get("undecryptableActions.sslV2Session"); value.Exists() {
		data.SslV2Session = types.StringValue(value.String())
	} else {
		data.SslV2Session = types.StringNull()
	}
	if value := res.Get("undecryptableActions.unknownCipherSuite"); value.Exists() {
		data.UnknownCipherSuite = types.StringValue(value.String())
	} else {
		data.UnknownCipherSuite = types.StringNull()
	}
	if value := res.Get("undecryptableActions.unsupportedCipherSuite"); value.Exists() {
		data.UnsupportedCipherSuite = types.StringValue(value.String())
	} else {
		data.UnsupportedCipherSuite = types.StringNull()
	}
	if value := res.Get("undecryptableActions.sessionNotCached"); value.Exists() {
		data.SessionNotCached = types.StringValue(value.String())
	} else {
		data.SessionNotCached = types.StringNull()
	}
	if value := res.Get("undecryptableActions.handshakeErrors"); value.Exists() {
		data.HandshakeErrors = types.StringValue(value.String())
	} else {
		data.HandshakeErrors = types.StringNull()
	}
	if value := res.Get("defaultAction.policyAction"); value.Exists() {
		data.PolicyAction = types.StringValue(value.String())
	} else {
		data.PolicyAction = types.StringNull()
	}
	if value := res.Get("defaultAction.eventLogAction"); value.Exists() {
		data.EventLogAction = types.StringValue(value.String())
	} else {
		data.EventLogAction = types.StringNull()
	}
	if value := res.Get("advancedOptions.disallowUntrustedIssuerResign"); value.Exists() {
		data.DisallowUntrustedIssuerResign = types.BoolValue(value.Bool())
	} else {
		data.DisallowUntrustedIssuerResign = types.BoolNull()
	}
	if value := res.Get("advancedOptions.stripHTTP3"); value.Exists() {
		data.StripHttp3 = types.BoolValue(value.Bool())
	} else {
		data.StripHttp3 = types.BoolNull()
	}
	if value := res.Get("advancedOptions.tls13Decryption"); value.Exists() {
		data.Tls13Decryption = types.BoolValue(value.Bool())
	} else {
		data.Tls13Decryption = types.BoolNull()
	}
	if value := res.Get("advancedOptions.quicDecryption"); value.Exists() {
		data.QuicDecryption = types.BoolValue(value.Bool())
	} else {
		data.QuicDecryption = types.BoolNull()
	}
	if value := res.Get("advancedOptions.adaptiveProbe"); value.Exists() {
		data.AdaptiveProbe = types.BoolValue(value.Bool())
	} else {
		data.AdaptiveProbe = types.BoolNull()
	}
	if value := res.Get("advancedOptions.blockExtensions"); value.Exists() {
		data.BlockExtensions = helpers.GetInt64List(value.Array())
	} else {
		data.BlockExtensions = types.ListNull(types.Int64Type)
	}
	if value := res.Get("logging.logEnd"); value.Exists() {
		data.LogEnd = types.BoolValue(value.Bool())
	} else {
		data.LogEnd = types.BoolNull()
	}
	if value := res.Get("logging.sendEvents"); value.Exists() {
		data.SendEvents = types.BoolValue(value.Bool())
	} else {
		data.SendEvents = types.BoolNull()
	}
	if value := res.Get("trustedCAs.objects"); value.Exists() {
		data.TrustedCas = make([]DecryptionPolicyTrustedCas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DecryptionPolicyTrustedCas{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.CaName = types.StringValue(cValue.String())
			} else {
				item.CaName = types.StringNull()
			}
			if cValue := v.Get("id"); cValue.Exists() {
				item.CaId = types.StringValue(cValue.String())
			} else {
				item.CaId = types.StringNull()
			}
			data.TrustedCas = append(data.TrustedCas, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DecryptionPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("undecryptableActions.decryptionErrors"); value.Exists() && !data.DecryptionErrors.IsNull() {
		data.DecryptionErrors = types.StringValue(value.String())
	} else {
		data.DecryptionErrors = types.StringNull()
	}
	if value := res.Get("undecryptableActions.compressedSession"); value.Exists() && !data.CompressedSession.IsNull() {
		data.CompressedSession = types.StringValue(value.String())
	} else {
		data.CompressedSession = types.StringNull()
	}
	if value := res.Get("undecryptableActions.sslV2Session"); value.Exists() && !data.SslV2Session.IsNull() {
		data.SslV2Session = types.StringValue(value.String())
	} else {
		data.SslV2Session = types.StringNull()
	}
	if value := res.Get("undecryptableActions.unknownCipherSuite"); value.Exists() && !data.UnknownCipherSuite.IsNull() {
		data.UnknownCipherSuite = types.StringValue(value.String())
	} else {
		data.UnknownCipherSuite = types.StringNull()
	}
	if value := res.Get("undecryptableActions.unsupportedCipherSuite"); value.Exists() && !data.UnsupportedCipherSuite.IsNull() {
		data.UnsupportedCipherSuite = types.StringValue(value.String())
	} else {
		data.UnsupportedCipherSuite = types.StringNull()
	}
	if value := res.Get("undecryptableActions.sessionNotCached"); value.Exists() && !data.SessionNotCached.IsNull() {
		data.SessionNotCached = types.StringValue(value.String())
	} else {
		data.SessionNotCached = types.StringNull()
	}
	if value := res.Get("undecryptableActions.handshakeErrors"); value.Exists() && !data.HandshakeErrors.IsNull() {
		data.HandshakeErrors = types.StringValue(value.String())
	} else {
		data.HandshakeErrors = types.StringNull()
	}
	if value := res.Get("defaultAction.policyAction"); value.Exists() && !data.PolicyAction.IsNull() {
		data.PolicyAction = types.StringValue(value.String())
	} else {
		data.PolicyAction = types.StringNull()
	}
	if value := res.Get("defaultAction.eventLogAction"); value.Exists() && !data.EventLogAction.IsNull() {
		data.EventLogAction = types.StringValue(value.String())
	} else {
		data.EventLogAction = types.StringNull()
	}
	if value := res.Get("advancedOptions.disallowUntrustedIssuerResign"); value.Exists() && !data.DisallowUntrustedIssuerResign.IsNull() {
		data.DisallowUntrustedIssuerResign = types.BoolValue(value.Bool())
	} else {
		data.DisallowUntrustedIssuerResign = types.BoolNull()
	}
	if value := res.Get("advancedOptions.stripHTTP3"); value.Exists() && !data.StripHttp3.IsNull() {
		data.StripHttp3 = types.BoolValue(value.Bool())
	} else {
		data.StripHttp3 = types.BoolNull()
	}
	if value := res.Get("advancedOptions.tls13Decryption"); value.Exists() && !data.Tls13Decryption.IsNull() {
		data.Tls13Decryption = types.BoolValue(value.Bool())
	} else {
		data.Tls13Decryption = types.BoolNull()
	}
	if value := res.Get("advancedOptions.quicDecryption"); value.Exists() && !data.QuicDecryption.IsNull() {
		data.QuicDecryption = types.BoolValue(value.Bool())
	} else {
		data.QuicDecryption = types.BoolNull()
	}
	if value := res.Get("advancedOptions.adaptiveProbe"); value.Exists() && !data.AdaptiveProbe.IsNull() {
		data.AdaptiveProbe = types.BoolValue(value.Bool())
	} else {
		data.AdaptiveProbe = types.BoolNull()
	}
	if value := res.Get("advancedOptions.blockExtensions"); value.Exists() && !data.BlockExtensions.IsNull() {
		data.BlockExtensions = helpers.GetInt64List(value.Array())
	} else {
		data.BlockExtensions = types.ListNull(types.Int64Type)
	}
	if value := res.Get("logging.logEnd"); value.Exists() && !data.LogEnd.IsNull() {
		data.LogEnd = types.BoolValue(value.Bool())
	} else {
		data.LogEnd = types.BoolNull()
	}
	if value := res.Get("logging.sendEvents"); value.Exists() && !data.SendEvents.IsNull() {
		data.SendEvents = types.BoolValue(value.Bool())
	} else {
		data.SendEvents = types.BoolNull()
	}
	for i := 0; i < len(data.TrustedCas); i++ {
		keys := [...]string{"name", "id"}
		keyValues := [...]string{data.TrustedCas[i].CaName.ValueString(), data.TrustedCas[i].CaId.ValueString()}

		var r gjson.Result
		res.Get("trustedCAs.objects").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing data.TrustedCas[%d] = %+v",
				i,
				data.TrustedCas[i],
			))
			data.TrustedCas = slices.Delete(data.TrustedCas, i, i+1)
			i--

			continue
		}
		if value := r.Get("name"); value.Exists() && !data.TrustedCas[i].CaName.IsNull() {
			data.TrustedCas[i].CaName = types.StringValue(value.String())
		} else {
			data.TrustedCas[i].CaName = types.StringNull()
		}
		if value := r.Get("id"); value.Exists() && !data.TrustedCas[i].CaId.IsNull() {
			data.TrustedCas[i].CaId = types.StringValue(value.String())
		} else {
			data.TrustedCas[i].CaId = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DecryptionPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
