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
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AutoNatRules struct {
	Id                     types.String `tfsdk:"id"`
	Domain                 types.String `tfsdk:"domain"`
	NatPolicyId            types.String `tfsdk:"nat_policy_id"`
	NatType                types.String `tfsdk:"nat_type"`
	InterfaceIpv6          types.Bool   `tfsdk:"interface_ipv6"`
	FallThrough            types.Bool   `tfsdk:"fall_through"`
	Dns                    types.Bool   `tfsdk:"dns"`
	RouteLookup            types.Bool   `tfsdk:"route_lookup"`
	NoProxyArp             types.Bool   `tfsdk:"no_proxy_arp"`
	NetToNet               types.Bool   `tfsdk:"net_to_net"`
	OriginalNetworkId      types.String `tfsdk:"original_network_id"`
	TranslatedNetworkId    types.String `tfsdk:"translated_network_id"`
	SourceInterfaceId      types.String `tfsdk:"source_interface_id"`
	DestinationInterfaceId types.String `tfsdk:"destination_interface_id"`
	InterfacePat           types.String `tfsdk:"interface_pat"`
	IncludeReverse         types.String `tfsdk:"include_reverse"`
	RoundRobin             types.String `tfsdk:"round_robin"`
	ExtendedPat            types.String `tfsdk:"extended_pat"`
	FlatPortRange          types.String `tfsdk:"flat_port_range"`
	BlockAllocation        types.String `tfsdk:"block_allocation"`
	PatPoolId              types.String `tfsdk:"pat_pool_id"`
	PatPoolName            types.String `tfsdk:"pat_pool_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AutoNatRules) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdnatpolicies/%v/autonatrules", url.QueryEscape(data.NatPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AutoNatRules) toBody(ctx context.Context, state AutoNatRules) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.NatType.IsNull() {
		body, _ = sjson.Set(body, "natType", data.NatType.ValueString())
	}
	if !data.InterfaceIpv6.IsNull() {
		body, _ = sjson.Set(body, "interfaceIpv6", data.InterfaceIpv6.ValueBool())
	}
	if !data.FallThrough.IsNull() {
		body, _ = sjson.Set(body, "fallThrough", data.FallThrough.ValueBool())
	}
	if !data.Dns.IsNull() {
		body, _ = sjson.Set(body, "dns", data.Dns.ValueBool())
	}
	if !data.RouteLookup.IsNull() {
		body, _ = sjson.Set(body, "routeLookup", data.RouteLookup.ValueBool())
	}
	if !data.NoProxyArp.IsNull() {
		body, _ = sjson.Set(body, "noProxyArp", data.NoProxyArp.ValueBool())
	}
	if !data.NetToNet.IsNull() {
		body, _ = sjson.Set(body, "netToNet", data.NetToNet.ValueBool())
	}
	if !data.OriginalNetworkId.IsNull() {
		body, _ = sjson.Set(body, "originalNetwork.id", data.OriginalNetworkId.ValueString())
	}
	if !data.TranslatedNetworkId.IsNull() {
		body, _ = sjson.Set(body, "translatedNetwork.id", data.TranslatedNetworkId.ValueString())
	}
	if !data.SourceInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "sourceInterface.id", data.SourceInterfaceId.ValueString())
	}
	if !data.DestinationInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "destinationInterface.id", data.DestinationInterfaceId.ValueString())
	}
	if !data.InterfacePat.IsNull() {
		body, _ = sjson.Set(body, "patOptions.interfacePat", data.InterfacePat.ValueString())
	}
	if !data.IncludeReverse.IsNull() {
		body, _ = sjson.Set(body, "patOptions.includeReverse", data.IncludeReverse.ValueString())
	}
	if !data.RoundRobin.IsNull() {
		body, _ = sjson.Set(body, "patOptions.roundRobin", data.RoundRobin.ValueString())
	}
	if !data.ExtendedPat.IsNull() {
		body, _ = sjson.Set(body, "patOptions.extendedPat", data.ExtendedPat.ValueString())
	}
	if !data.FlatPortRange.IsNull() {
		body, _ = sjson.Set(body, "patOptions.flatPortRange", data.FlatPortRange.ValueString())
	}
	if !data.BlockAllocation.IsNull() {
		body, _ = sjson.Set(body, "patOptions.blockAllocation", data.BlockAllocation.ValueString())
	}
	if !data.PatPoolId.IsNull() {
		body, _ = sjson.Set(body, "patOptions.patPoolAddress.id", data.PatPoolId.ValueString())
	}
	if !data.PatPoolName.IsNull() {
		body, _ = sjson.Set(body, "patOptions.patPoolAddress.name", data.PatPoolName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AutoNatRules) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("natType"); value.Exists() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("interfaceIpv6"); value.Exists() {
		data.InterfaceIpv6 = types.BoolValue(value.Bool())
	} else {
		data.InterfaceIpv6 = types.BoolNull()
	}
	if value := res.Get("fallThrough"); value.Exists() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("dns"); value.Exists() {
		data.Dns = types.BoolValue(value.Bool())
	} else {
		data.Dns = types.BoolNull()
	}
	if value := res.Get("routeLookup"); value.Exists() {
		data.RouteLookup = types.BoolValue(value.Bool())
	} else {
		data.RouteLookup = types.BoolNull()
	}
	if value := res.Get("noProxyArp"); value.Exists() {
		data.NoProxyArp = types.BoolValue(value.Bool())
	} else {
		data.NoProxyArp = types.BoolNull()
	}
	if value := res.Get("netToNet"); value.Exists() {
		data.NetToNet = types.BoolValue(value.Bool())
	} else {
		data.NetToNet = types.BoolNull()
	}
	if value := res.Get("originalNetwork.id"); value.Exists() {
		data.OriginalNetworkId = types.StringValue(value.String())
	} else {
		data.OriginalNetworkId = types.StringNull()
	}
	if value := res.Get("translatedNetwork.id"); value.Exists() {
		data.TranslatedNetworkId = types.StringValue(value.String())
	} else {
		data.TranslatedNetworkId = types.StringNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("patOptions.interfacePat"); value.Exists() {
		data.InterfacePat = types.StringValue(value.String())
	} else {
		data.InterfacePat = types.StringNull()
	}
	if value := res.Get("patOptions.includeReverse"); value.Exists() {
		data.IncludeReverse = types.StringValue(value.String())
	} else {
		data.IncludeReverse = types.StringNull()
	}
	if value := res.Get("patOptions.roundRobin"); value.Exists() {
		data.RoundRobin = types.StringValue(value.String())
	} else {
		data.RoundRobin = types.StringNull()
	}
	if value := res.Get("patOptions.extendedPat"); value.Exists() {
		data.ExtendedPat = types.StringValue(value.String())
	} else {
		data.ExtendedPat = types.StringNull()
	}
	if value := res.Get("patOptions.flatPortRange"); value.Exists() {
		data.FlatPortRange = types.StringValue(value.String())
	} else {
		data.FlatPortRange = types.StringNull()
	}
	if value := res.Get("patOptions.blockAllocation"); value.Exists() {
		data.BlockAllocation = types.StringValue(value.String())
	} else {
		data.BlockAllocation = types.StringNull()
	}
	if value := res.Get("patOptions.patPoolAddress.id"); value.Exists() {
		data.PatPoolId = types.StringValue(value.String())
	} else {
		data.PatPoolId = types.StringNull()
	}
	if value := res.Get("patOptions.patPoolAddress.name"); value.Exists() {
		data.PatPoolName = types.StringValue(value.String())
	} else {
		data.PatPoolName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AutoNatRules) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("natType"); value.Exists() && !data.NatType.IsNull() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("interfaceIpv6"); value.Exists() && !data.InterfaceIpv6.IsNull() {
		data.InterfaceIpv6 = types.BoolValue(value.Bool())
	} else {
		data.InterfaceIpv6 = types.BoolNull()
	}
	if value := res.Get("fallThrough"); value.Exists() && !data.FallThrough.IsNull() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("dns"); value.Exists() && !data.Dns.IsNull() {
		data.Dns = types.BoolValue(value.Bool())
	} else {
		data.Dns = types.BoolNull()
	}
	if value := res.Get("routeLookup"); value.Exists() && !data.RouteLookup.IsNull() {
		data.RouteLookup = types.BoolValue(value.Bool())
	} else {
		data.RouteLookup = types.BoolNull()
	}
	if value := res.Get("noProxyArp"); value.Exists() && !data.NoProxyArp.IsNull() {
		data.NoProxyArp = types.BoolValue(value.Bool())
	} else {
		data.NoProxyArp = types.BoolNull()
	}
	if value := res.Get("netToNet"); value.Exists() && !data.NetToNet.IsNull() {
		data.NetToNet = types.BoolValue(value.Bool())
	} else {
		data.NetToNet = types.BoolNull()
	}
	if value := res.Get("originalNetwork.id"); value.Exists() && !data.OriginalNetworkId.IsNull() {
		data.OriginalNetworkId = types.StringValue(value.String())
	} else {
		data.OriginalNetworkId = types.StringNull()
	}
	if value := res.Get("translatedNetwork.id"); value.Exists() && !data.TranslatedNetworkId.IsNull() {
		data.TranslatedNetworkId = types.StringValue(value.String())
	} else {
		data.TranslatedNetworkId = types.StringNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() && !data.DestinationInterfaceId.IsNull() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("patOptions.interfacePat"); value.Exists() && !data.InterfacePat.IsNull() {
		data.InterfacePat = types.StringValue(value.String())
	} else {
		data.InterfacePat = types.StringNull()
	}
	if value := res.Get("patOptions.includeReverse"); value.Exists() && !data.IncludeReverse.IsNull() {
		data.IncludeReverse = types.StringValue(value.String())
	} else {
		data.IncludeReverse = types.StringNull()
	}
	if value := res.Get("patOptions.roundRobin"); value.Exists() && !data.RoundRobin.IsNull() {
		data.RoundRobin = types.StringValue(value.String())
	} else {
		data.RoundRobin = types.StringNull()
	}
	if value := res.Get("patOptions.extendedPat"); value.Exists() && !data.ExtendedPat.IsNull() {
		data.ExtendedPat = types.StringValue(value.String())
	} else {
		data.ExtendedPat = types.StringNull()
	}
	if value := res.Get("patOptions.flatPortRange"); value.Exists() && !data.FlatPortRange.IsNull() {
		data.FlatPortRange = types.StringValue(value.String())
	} else {
		data.FlatPortRange = types.StringNull()
	}
	if value := res.Get("patOptions.blockAllocation"); value.Exists() && !data.BlockAllocation.IsNull() {
		data.BlockAllocation = types.StringValue(value.String())
	} else {
		data.BlockAllocation = types.StringNull()
	}
	if value := res.Get("patOptions.patPoolAddress.id"); value.Exists() && !data.PatPoolId.IsNull() {
		data.PatPoolId = types.StringValue(value.String())
	} else {
		data.PatPoolId = types.StringNull()
	}
	if value := res.Get("patOptions.patPoolAddress.name"); value.Exists() && !data.PatPoolName.IsNull() {
		data.PatPoolName = types.StringValue(value.String())
	} else {
		data.PatPoolName = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AutoNatRules) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
