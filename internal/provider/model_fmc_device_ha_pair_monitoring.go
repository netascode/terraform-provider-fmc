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

type DeviceHAPairMonitoring struct {
	Id                 types.String `tfsdk:"id"`
	Domain             types.String `tfsdk:"domain"`
	DeviceId           types.String `tfsdk:"device_id"`
	Type               types.String `tfsdk:"type"`
	Name               types.String `tfsdk:"name"`
	MonitorInterface   types.Bool   `tfsdk:"monitor_interface"`
	Ipv4ActiveAddress  types.String `tfsdk:"ipv4_active_address"`
	Ipv4StandbyAddress types.String `tfsdk:"ipv4_standby_address"`
	Ipv4Netmask        types.String `tfsdk:"ipv4_netmask"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceHAPairMonitoring) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicehapairs/ftddevicehapairs/%v/monitoredinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceHAPairMonitoring) toBody(ctx context.Context, state DeviceHAPairMonitoring) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.MonitorInterface.IsNull() {
		body, _ = sjson.Set(body, "monitorForFailures", data.MonitorInterface.ValueBool())
	}
	if !data.Ipv4StandbyAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4Configuration.standbyIPv4Address", data.Ipv4StandbyAddress.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceHAPairMonitoring) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("monitorForFailures"); value.Exists() {
		data.MonitorInterface = types.BoolValue(value.Bool())
	} else {
		data.MonitorInterface = types.BoolNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() {
		data.Ipv4ActiveAddress = types.StringValue(value.String())
	} else {
		data.Ipv4ActiveAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.standbyIPv4Address"); value.Exists() {
		data.Ipv4StandbyAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StandbyAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceHAPairMonitoring) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("monitorForFailures"); value.Exists() && !data.MonitorInterface.IsNull() {
		data.MonitorInterface = types.BoolValue(value.Bool())
	} else {
		data.MonitorInterface = types.BoolNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() && !data.Ipv4ActiveAddress.IsNull() {
		data.Ipv4ActiveAddress = types.StringValue(value.String())
	} else {
		data.Ipv4ActiveAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.standbyIPv4Address"); value.Exists() && !data.Ipv4StandbyAddress.IsNull() {
		data.Ipv4StandbyAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StandbyAddress = types.StringNull()
	}
	if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() && !data.Ipv4Netmask.IsNull() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceHAPairMonitoring) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Ipv4ActiveAddress.IsUnknown() {
		if value := res.Get("ipv4Configuration.activeIPv4Address"); value.Exists() {
			data.Ipv4ActiveAddress = types.StringValue(value.String())
		} else {
			data.Ipv4ActiveAddress = types.StringNull()
		}
	}
	if data.Ipv4Netmask.IsUnknown() {
		if value := res.Get("ipv4Configuration.activeIPv4Mask"); value.Exists() {
			data.Ipv4Netmask = types.StringValue(value.String())
		} else {
			data.Ipv4Netmask = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// toBodyPutDelete generates minimal required body to reset the resource to its default state.
func (data DeviceHAPairMonitoring) toBodyPutDelete(ctx context.Context, state DeviceHAPairMonitoring) string {
	body := ""
	body, _ = sjson.Set(body, "monitorForFailures", false)
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	return body
}
