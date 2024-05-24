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
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Device struct {
	Id             types.String `tfsdk:"id"`
	Domain         types.String `tfsdk:"domain"`
	HostName       types.String `tfsdk:"host_name"`
	LicenseCaps    types.List   `tfsdk:"license_caps"`
	RegKey         types.String `tfsdk:"reg_key"`
	Type           types.String `tfsdk:"type"`
	AccessPolicyId types.String `tfsdk:"access_policy_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Device) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Device) toBody(ctx context.Context, state Device) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, "hostName", data.HostName.ValueString())
	}
	if !data.LicenseCaps.IsNull() {
		var values []string
		data.LicenseCaps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "license_caps", values)
	}
	if !data.RegKey.IsNull() {
		body, _ = sjson.Set(body, "regKey", data.RegKey.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.AccessPolicyId.IsNull() {
		body, _ = sjson.Set(body, "accessPolicy.id", data.AccessPolicyId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("hostName"); value.Exists() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() {
		data.LicenseCaps = helpers.GetStringList(value.Array())
	} else {
		data.LicenseCaps = types.ListNull(types.StringType)
	}
	if value := res.Get("regKey"); value.Exists() {
		data.RegKey = types.StringValue(value.String())
	} else {
		data.RegKey = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringValue("Device")
	}
	if value := res.Get("accessPolicy.id"); value.Exists() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Device) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("hostName"); value.Exists() && !data.HostName.IsNull() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() && !data.LicenseCaps.IsNull() {
		data.LicenseCaps = helpers.GetStringList(value.Array())
	} else {
		data.LicenseCaps = types.ListNull(types.StringType)
	}
	if value := res.Get("regKey"); value.Exists() && !data.RegKey.IsNull() {
		data.RegKey = types.StringValue(value.String())
	} else {
		data.RegKey = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else if data.Type.ValueString() != "Device" {
		data.Type = types.StringNull()
	}
	if value := res.Get("accessPolicy.id"); value.Exists() && !data.AccessPolicyId.IsNull() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Device) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.HostName.IsNull() {
		return false
	}
	if !data.LicenseCaps.IsNull() {
		return false
	}
	if !data.RegKey.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.AccessPolicyId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
