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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceHAPair struct {
	Id                            types.String `tfsdk:"id"`
	Domain                        types.String `tfsdk:"domain"`
	Name                          types.String `tfsdk:"name"`
	PrimaryDeviceId               types.String `tfsdk:"primary_device_id"`
	SecondaryDeviceId             types.String `tfsdk:"secondary_device_id"`
	IsEncryptionEnabled           types.Bool   `tfsdk:"is_encryption_enabled"`
	UseSameLinkForFailovers       types.Bool   `tfsdk:"use_same_link_for_failovers"`
	SharedKey                     types.String `tfsdk:"shared_key"`
	EncKeyGenerationScheme        types.String `tfsdk:"enc_key_generation_scheme"`
	LanFailoverStandbyIp          types.String `tfsdk:"lan_failover_standby_ip"`
	LanFailoverActiveIp           types.String `tfsdk:"lan_failover_active_ip"`
	LanFailoverName               types.String `tfsdk:"lan_failover_name"`
	LanFailoverSubnetMask         types.String `tfsdk:"lan_failover_subnet_mask"`
	LanFailoverIpv6Addr           types.Bool   `tfsdk:"lan_failover_ipv6_addr"`
	LanFailoverInterfaceName      types.String `tfsdk:"lan_failover_interface_name"`
	LanFailoverInterfaceId        types.String `tfsdk:"lan_failover_interface_id"`
	LanFailoverInterfaceType      types.String `tfsdk:"lan_failover_interface_type"`
	StatefulFailoverStandbyIp     types.String `tfsdk:"stateful_failover_standby_ip"`
	StatefulFailoverActiveIp      types.String `tfsdk:"stateful_failover_active_ip"`
	StatefulFailoverName          types.String `tfsdk:"stateful_failover_name"`
	StatefulFailoverSubnetMask    types.String `tfsdk:"stateful_failover_subnet_mask"`
	StatefulFailoverIpv6Addr      types.Bool   `tfsdk:"stateful_failover_ipv6_addr"`
	StatefulFailoverInterfaceName types.String `tfsdk:"stateful_failover_interface_name"`
	StatefulFailoverInterfaceId   types.String `tfsdk:"stateful_failover_interface_id"`
	StatefulFailoverInterfaceType types.String `tfsdk:"stateful_failover_interface_type"`
	Action                        types.String `tfsdk:"action"`
	ForceBreak                    types.Bool   `tfsdk:"force_break"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceHAPair) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicehapairs/ftddevicehapairs"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceHAPair) toBody(ctx context.Context, state DeviceHAPair) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DeviceHAPair")
	if !data.PrimaryDeviceId.IsNull() {
		body, _ = sjson.Set(body, "primary.id", data.PrimaryDeviceId.ValueString())
	}
	if !data.SecondaryDeviceId.IsNull() {
		body, _ = sjson.Set(body, "secondary.id", data.SecondaryDeviceId.ValueString())
	}
	if !data.IsEncryptionEnabled.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.isEncryptionEnabled", data.IsEncryptionEnabled.ValueBool())
	}
	if !data.UseSameLinkForFailovers.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.useSameLinkForFailovers", data.UseSameLinkForFailovers.ValueBool())
	}
	if !data.SharedKey.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.sharedKey", data.SharedKey.ValueString())
	}
	if !data.EncKeyGenerationScheme.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.encKeyGenerationScheme", data.EncKeyGenerationScheme.ValueString())
	}
	if !data.LanFailoverStandbyIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.standbyIP", data.LanFailoverStandbyIp.ValueString())
	}
	if !data.LanFailoverActiveIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.activeIP", data.LanFailoverActiveIp.ValueString())
	}
	if !data.LanFailoverName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.logicalName", data.LanFailoverName.ValueString())
	}
	if !data.LanFailoverSubnetMask.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.subnetMask", data.LanFailoverSubnetMask.ValueString())
	}
	if !data.LanFailoverIpv6Addr.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.useIPv6Address", data.LanFailoverIpv6Addr.ValueBool())
	}
	if !data.LanFailoverInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.name", data.LanFailoverInterfaceName.ValueString())
	}
	if !data.LanFailoverInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.id", data.LanFailoverInterfaceId.ValueString())
	}
	if !data.LanFailoverInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.lanFailover.interfaceObject.type", data.LanFailoverInterfaceType.ValueString())
	}
	if !data.StatefulFailoverStandbyIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.standbyIP", data.StatefulFailoverStandbyIp.ValueString())
	}
	if !data.StatefulFailoverActiveIp.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.activeIP", data.StatefulFailoverActiveIp.ValueString())
	}
	if !data.StatefulFailoverName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.logicalName", data.StatefulFailoverName.ValueString())
	}
	if !data.StatefulFailoverSubnetMask.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.subnetMask", data.StatefulFailoverSubnetMask.ValueString())
	}
	if !data.StatefulFailoverIpv6Addr.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.useIPv6Address", data.StatefulFailoverIpv6Addr.ValueBool())
	}
	if !data.StatefulFailoverInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.name", data.StatefulFailoverInterfaceName.ValueString())
	}
	if !data.StatefulFailoverInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.id", data.StatefulFailoverInterfaceId.ValueString())
	}
	if !data.StatefulFailoverInterfaceType.IsNull() {
		body, _ = sjson.Set(body, "ftdHABootstrap.statefulFailover.interfaceObject.type", data.StatefulFailoverInterfaceType.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, "action", data.Action.ValueString())
	}
	if !data.ForceBreak.IsNull() {
		body, _ = sjson.Set(body, "forceBreak", data.ForceBreak.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceHAPair) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("primary.id"); value.Exists() {
		data.PrimaryDeviceId = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceId = types.StringNull()
	}
	if value := res.Get("secondary.id"); value.Exists() {
		data.SecondaryDeviceId = types.StringValue(value.String())
	} else {
		data.SecondaryDeviceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.isEncryptionEnabled"); value.Exists() {
		data.IsEncryptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsEncryptionEnabled = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.sharedKey"); value.Exists() {
		data.SharedKey = types.StringValue(value.String())
	} else {
		data.SharedKey = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.encKeyGenerationScheme"); value.Exists() {
		data.EncKeyGenerationScheme = types.StringValue(value.String())
	} else {
		data.EncKeyGenerationScheme = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.standbyIP"); value.Exists() {
		data.LanFailoverStandbyIp = types.StringValue(value.String())
	} else {
		data.LanFailoverStandbyIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.activeIP"); value.Exists() {
		data.LanFailoverActiveIp = types.StringValue(value.String())
	} else {
		data.LanFailoverActiveIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.logicalName"); value.Exists() {
		data.LanFailoverName = types.StringValue(value.String())
	} else {
		data.LanFailoverName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.subnetMask"); value.Exists() {
		data.LanFailoverSubnetMask = types.StringValue(value.String())
	} else {
		data.LanFailoverSubnetMask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.interfaceObject.name"); value.Exists() {
		data.LanFailoverInterfaceName = types.StringValue(value.String())
	} else {
		data.LanFailoverInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.standbyIP"); value.Exists() {
		data.StatefulFailoverStandbyIp = types.StringValue(value.String())
	} else {
		data.StatefulFailoverStandbyIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.activeIP"); value.Exists() {
		data.StatefulFailoverActiveIp = types.StringValue(value.String())
	} else {
		data.StatefulFailoverActiveIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.logicalName"); value.Exists() {
		data.StatefulFailoverName = types.StringValue(value.String())
	} else {
		data.StatefulFailoverName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.subnetMask"); value.Exists() {
		data.StatefulFailoverSubnetMask = types.StringValue(value.String())
	} else {
		data.StatefulFailoverSubnetMask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.useIPv6Address"); value.Exists() {
		data.StatefulFailoverIpv6Addr = types.BoolValue(value.Bool())
	} else {
		data.StatefulFailoverIpv6Addr = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.name"); value.Exists() {
		data.StatefulFailoverInterfaceName = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.id"); value.Exists() {
		data.StatefulFailoverInterfaceId = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.type"); value.Exists() {
		data.StatefulFailoverInterfaceType = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceType = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("forceBreak"); value.Exists() {
		data.ForceBreak = types.BoolValue(value.Bool())
	} else {
		data.ForceBreak = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceHAPair) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("primary.id"); value.Exists() && !data.PrimaryDeviceId.IsNull() {
		data.PrimaryDeviceId = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceId = types.StringNull()
	}
	if value := res.Get("secondary.id"); value.Exists() && !data.SecondaryDeviceId.IsNull() {
		data.SecondaryDeviceId = types.StringValue(value.String())
	} else {
		data.SecondaryDeviceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.isEncryptionEnabled"); value.Exists() && !data.IsEncryptionEnabled.IsNull() {
		data.IsEncryptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsEncryptionEnabled = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.sharedKey"); value.Exists() && !data.SharedKey.IsNull() {
		data.SharedKey = types.StringValue(value.String())
	} else {
		data.SharedKey = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.encKeyGenerationScheme"); value.Exists() && !data.EncKeyGenerationScheme.IsNull() {
		data.EncKeyGenerationScheme = types.StringValue(value.String())
	} else {
		data.EncKeyGenerationScheme = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.standbyIP"); value.Exists() && !data.LanFailoverStandbyIp.IsNull() {
		data.LanFailoverStandbyIp = types.StringValue(value.String())
	} else {
		data.LanFailoverStandbyIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.activeIP"); value.Exists() && !data.LanFailoverActiveIp.IsNull() {
		data.LanFailoverActiveIp = types.StringValue(value.String())
	} else {
		data.LanFailoverActiveIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.logicalName"); value.Exists() && !data.LanFailoverName.IsNull() {
		data.LanFailoverName = types.StringValue(value.String())
	} else {
		data.LanFailoverName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.subnetMask"); value.Exists() && !data.LanFailoverSubnetMask.IsNull() {
		data.LanFailoverSubnetMask = types.StringValue(value.String())
	} else {
		data.LanFailoverSubnetMask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.lanFailover.interfaceObject.name"); value.Exists() && !data.LanFailoverInterfaceName.IsNull() {
		data.LanFailoverInterfaceName = types.StringValue(value.String())
	} else {
		data.LanFailoverInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.standbyIP"); value.Exists() && !data.StatefulFailoverStandbyIp.IsNull() {
		data.StatefulFailoverStandbyIp = types.StringValue(value.String())
	} else {
		data.StatefulFailoverStandbyIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.activeIP"); value.Exists() && !data.StatefulFailoverActiveIp.IsNull() {
		data.StatefulFailoverActiveIp = types.StringValue(value.String())
	} else {
		data.StatefulFailoverActiveIp = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.logicalName"); value.Exists() && !data.StatefulFailoverName.IsNull() {
		data.StatefulFailoverName = types.StringValue(value.String())
	} else {
		data.StatefulFailoverName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.subnetMask"); value.Exists() && !data.StatefulFailoverSubnetMask.IsNull() {
		data.StatefulFailoverSubnetMask = types.StringValue(value.String())
	} else {
		data.StatefulFailoverSubnetMask = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.useIPv6Address"); value.Exists() && !data.StatefulFailoverIpv6Addr.IsNull() {
		data.StatefulFailoverIpv6Addr = types.BoolValue(value.Bool())
	} else {
		data.StatefulFailoverIpv6Addr = types.BoolNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.name"); value.Exists() && !data.StatefulFailoverInterfaceName.IsNull() {
		data.StatefulFailoverInterfaceName = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceName = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.id"); value.Exists() && !data.StatefulFailoverInterfaceId.IsNull() {
		data.StatefulFailoverInterfaceId = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceId = types.StringNull()
	}
	if value := res.Get("ftdHABootstrap.statefulFailover.interfaceObject.type"); value.Exists() && !data.StatefulFailoverInterfaceType.IsNull() {
		data.StatefulFailoverInterfaceType = types.StringValue(value.String())
	} else {
		data.StatefulFailoverInterfaceType = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("forceBreak"); value.Exists() && !data.ForceBreak.IsNull() {
		data.ForceBreak = types.BoolValue(value.Bool())
	} else {
		data.ForceBreak = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceHAPair) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
