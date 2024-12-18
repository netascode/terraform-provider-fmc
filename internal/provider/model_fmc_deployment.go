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

type Deployment struct {
	Id             types.String `tfsdk:"id"`
	Domain         types.String `tfsdk:"domain"`
	Version        types.String `tfsdk:"version"`
	ForceDeploy    types.Bool   `tfsdk:"force_deploy"`
	DeviceList     types.List   `tfsdk:"device_list"`
	DeploymentNote types.String `tfsdk:"deployment_note"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Deployment) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deploymentrequests"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Deployment) toBody(ctx context.Context, state Deployment) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DeploymentRequest")
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, "version", data.Version.ValueString())
	}
	if !data.ForceDeploy.IsNull() {
		body, _ = sjson.Set(body, "forceDeploy", data.ForceDeploy.ValueBool())
	}
	body, _ = sjson.Set(body, "ignoreWarning", true)
	if !data.DeviceList.IsNull() {
		var values []string
		data.DeviceList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "deviceList", values)
	}
	if !data.DeploymentNote.IsNull() {
		body, _ = sjson.Set(body, "deploymentNote", data.DeploymentNote.ValueString())
	}
	body, _ = sjson.Set(body, "save", true)
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Deployment) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("version"); value.Exists() {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
	}
	if value := res.Get("forceDeploy"); value.Exists() {
		data.ForceDeploy = types.BoolValue(value.Bool())
	} else {
		data.ForceDeploy = types.BoolNull()
	}
	if value := res.Get("deviceList"); value.Exists() {
		data.DeviceList = helpers.GetStringList(value.Array())
	} else {
		data.DeviceList = types.ListNull(types.StringType)
	}
	if value := res.Get("deploymentNote"); value.Exists() {
		data.DeploymentNote = types.StringValue(value.String())
	} else {
		data.DeploymentNote = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *Deployment) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("version"); value.Exists() && !data.Version.IsNull() {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
	}
	if value := res.Get("forceDeploy"); value.Exists() && !data.ForceDeploy.IsNull() {
		data.ForceDeploy = types.BoolValue(value.Bool())
	} else {
		data.ForceDeploy = types.BoolNull()
	}
	if value := res.Get("deviceList"); value.Exists() && !data.DeviceList.IsNull() {
		data.DeviceList = helpers.GetStringList(value.Array())
	} else {
		data.DeviceList = types.ListNull(types.StringType)
	}
	if value := res.Get("deploymentNote"); value.Exists() && !data.DeploymentNote.IsNull() {
		data.DeploymentNote = types.StringValue(value.String())
	} else {
		data.DeploymentNote = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Deployment) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk