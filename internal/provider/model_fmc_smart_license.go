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

type SmartLicense struct {
	Id                 types.String `tfsdk:"id"`
	Domain             types.String `tfsdk:"domain"`
	RegistrationType   types.String `tfsdk:"registration_type"`
	Token              types.String `tfsdk:"token"`
	RegistrationStatus types.String `tfsdk:"registration_status"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SmartLicense) getPath() string {
	return "/api/fmc_platform/v1/license/smartlicenses"
}

// End of section. //template:end getPath

func (data SmartLicense) toBody(ctx context.Context, state SmartLicense) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.RegistrationType.IsNull() {
		body, _ = sjson.Set(body, "registrationType", data.RegistrationType.ValueString())
	}
	if !data.Token.IsNull() {
		body, _ = sjson.Set(body, "token", data.Token.ValueString())
	}
	return body
}

func (data *SmartLicense) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("registrationType"); value.Exists() {
		data.RegistrationType = types.StringValue(value.String())
	}
	if value := res.Get("token"); value.Exists() {
		data.Token = types.StringValue(value.String())
	}
	if value := res.Get("regStatus"); value.Exists() {
		data.RegistrationStatus = types.StringValue(value.String())
	} else {
		data.RegistrationStatus = types.StringNull()
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SmartLicense) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("registrationType"); value.Exists() && !data.RegistrationType.IsNull() {
		data.RegistrationType = types.StringValue(value.String())
	} else {
		data.RegistrationType = types.StringNull()
	}
	if value := res.Get("token"); value.Exists() && !data.Token.IsNull() {
		data.Token = types.StringValue(value.String())
	} else {
		data.Token = types.StringNull()
	}
	if value := res.Get("regStatus"); value.Exists() && !data.RegistrationStatus.IsNull() {
		data.RegistrationStatus = types.StringValue(value.String())
	} else {
		data.RegistrationStatus = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SmartLicense) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
