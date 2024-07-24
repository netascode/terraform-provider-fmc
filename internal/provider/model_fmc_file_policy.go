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

type FilePolicy struct {
	Id                     types.String `tfsdk:"id"`
	Domain                 types.String `tfsdk:"domain"`
	Name                   types.String `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	ArchiveDepth           types.String `tfsdk:"archive_depth"`
	ArchiveDepthAction     types.Bool   `tfsdk:"archive_depth_action"`
	BlockEncryptedArchives types.Bool   `tfsdk:"block_encrypted_archives"`
	CleanList              types.Bool   `tfsdk:"clean_list"`
	CustomDetectionList    types.Bool   `tfsdk:"custom_detection_list"`
	FirstTimeFileAnalysis  types.Bool   `tfsdk:"first_time_file_analysis"`
	InspectArchives        types.Bool   `tfsdk:"inspect_archives"`
	ThreatScore            types.String `tfsdk:"threat_score"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FilePolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/filepolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FilePolicy) toBody(ctx context.Context, state FilePolicy) string {
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
	if !data.ArchiveDepth.IsNull() {
		body, _ = sjson.Set(body, "archiveDepth", data.ArchiveDepth.ValueString())
	}
	if !data.ArchiveDepthAction.IsNull() {
		body, _ = sjson.Set(body, "archiveDepthAction", data.ArchiveDepthAction.ValueBool())
	}
	if !data.BlockEncryptedArchives.IsNull() {
		body, _ = sjson.Set(body, "blockEncryptedArchives", data.BlockEncryptedArchives.ValueBool())
	}
	if !data.CleanList.IsNull() {
		body, _ = sjson.Set(body, "cleanList", data.CleanList.ValueBool())
	}
	if !data.CustomDetectionList.IsNull() {
		body, _ = sjson.Set(body, "customDetectionList", data.CustomDetectionList.ValueBool())
	}
	if !data.FirstTimeFileAnalysis.IsNull() {
		body, _ = sjson.Set(body, "firstTimeFileAnalysis", data.FirstTimeFileAnalysis.ValueBool())
	}
	if !data.InspectArchives.IsNull() {
		body, _ = sjson.Set(body, "inspectArchives", data.InspectArchives.ValueBool())
	}
	if !data.ThreatScore.IsNull() {
		body, _ = sjson.Set(body, "threatScore", data.ThreatScore.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FilePolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("archiveDepth"); value.Exists() {
		data.ArchiveDepth = types.StringValue(value.String())
	} else {
		data.ArchiveDepth = types.StringNull()
	}
	if value := res.Get("archiveDepthAction"); value.Exists() {
		data.ArchiveDepthAction = types.BoolValue(value.Bool())
	} else {
		data.ArchiveDepthAction = types.BoolNull()
	}
	if value := res.Get("blockEncryptedArchives"); value.Exists() {
		data.BlockEncryptedArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockEncryptedArchives = types.BoolNull()
	}
	if value := res.Get("cleanList"); value.Exists() {
		data.CleanList = types.BoolValue(value.Bool())
	} else {
		data.CleanList = types.BoolNull()
	}
	if value := res.Get("customDetectionList"); value.Exists() {
		data.CustomDetectionList = types.BoolValue(value.Bool())
	} else {
		data.CustomDetectionList = types.BoolNull()
	}
	if value := res.Get("firstTimeFileAnalysis"); value.Exists() {
		data.FirstTimeFileAnalysis = types.BoolValue(value.Bool())
	} else {
		data.FirstTimeFileAnalysis = types.BoolNull()
	}
	if value := res.Get("inspectArchives"); value.Exists() {
		data.InspectArchives = types.BoolValue(value.Bool())
	} else {
		data.InspectArchives = types.BoolNull()
	}
	if value := res.Get("threatScore"); value.Exists() {
		data.ThreatScore = types.StringValue(value.String())
	} else {
		data.ThreatScore = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FilePolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("archiveDepth"); value.Exists() && !data.ArchiveDepth.IsNull() {
		data.ArchiveDepth = types.StringValue(value.String())
	} else {
		data.ArchiveDepth = types.StringNull()
	}
	if value := res.Get("archiveDepthAction"); value.Exists() && !data.ArchiveDepthAction.IsNull() {
		data.ArchiveDepthAction = types.BoolValue(value.Bool())
	} else {
		data.ArchiveDepthAction = types.BoolNull()
	}
	if value := res.Get("blockEncryptedArchives"); value.Exists() && !data.BlockEncryptedArchives.IsNull() {
		data.BlockEncryptedArchives = types.BoolValue(value.Bool())
	} else {
		data.BlockEncryptedArchives = types.BoolNull()
	}
	if value := res.Get("cleanList"); value.Exists() && !data.CleanList.IsNull() {
		data.CleanList = types.BoolValue(value.Bool())
	} else {
		data.CleanList = types.BoolNull()
	}
	if value := res.Get("customDetectionList"); value.Exists() && !data.CustomDetectionList.IsNull() {
		data.CustomDetectionList = types.BoolValue(value.Bool())
	} else {
		data.CustomDetectionList = types.BoolNull()
	}
	if value := res.Get("firstTimeFileAnalysis"); value.Exists() && !data.FirstTimeFileAnalysis.IsNull() {
		data.FirstTimeFileAnalysis = types.BoolValue(value.Bool())
	} else {
		data.FirstTimeFileAnalysis = types.BoolNull()
	}
	if value := res.Get("inspectArchives"); value.Exists() && !data.InspectArchives.IsNull() {
		data.InspectArchives = types.BoolValue(value.Bool())
	} else {
		data.InspectArchives = types.BoolNull()
	}
	if value := res.Get("threatScore"); value.Exists() && !data.ThreatScore.IsNull() {
		data.ThreatScore = types.StringValue(value.String())
	} else {
		data.ThreatScore = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FilePolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
