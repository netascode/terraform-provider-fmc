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
type SecurityZone struct {
	Id            types.String `tfsdk:"id"`
	Domain        types.String `tfsdk:"domain"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Overridable   types.Bool   `tfsdk:"overridable"`
	InterfaceMode types.String `tfsdk:"interface_mode"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SecurityZone) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/securityzones"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SecurityZone) toBody(ctx context.Context, state SecurityZone) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Overridable.IsNull() {
		body, _ = sjson.Set(body, "overridable", data.Overridable.ValueBool())
	}
	if !data.InterfaceMode.IsNull() {
		body, _ = sjson.Set(body, "interfaceMode", data.InterfaceMode.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SecurityZone) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("interfaceMode"); value.Exists() {
		data.InterfaceMode = types.StringValue(value.String())
	} else {
		data.InterfaceMode = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SecurityZone) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("interfaceMode"); value.Exists() && !data.InterfaceMode.IsNull() {
		data.InterfaceMode = types.StringValue(value.String())
	} else {
		data.InterfaceMode = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SecurityZone) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Overridable.IsNull() {
		return false
	}
	if !data.InterfaceMode.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull