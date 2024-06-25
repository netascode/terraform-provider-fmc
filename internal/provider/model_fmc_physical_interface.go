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
type PhysicalInterface struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	DeviceId    types.String `tfsdk:"device_id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Mode        types.String `tfsdk:"mode"`
	Name        types.String `tfsdk:"name"`
	LogicalName types.String `tfsdk:"logical_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PhysicalInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/physicalinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PhysicalInterface) toBody(ctx context.Context, state PhysicalInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "PhysicalInterface")
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PhysicalInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PhysicalInterface) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PhysicalInterface) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceId.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.LogicalName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
