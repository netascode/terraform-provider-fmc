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

type DynamicObject struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	ObjectType  types.String `tfsdk:"object_type"`
	AgentId     types.String `tfsdk:"agent_id"`
	TopicName   types.String `tfsdk:"topic_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DynamicObject) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/dynamicobjects"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DynamicObject) toBody(ctx context.Context, state DynamicObject) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "DynamicObject")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ObjectType.IsNull() {
		body, _ = sjson.Set(body, "objectType", data.ObjectType.ValueString())
	}
	if !data.AgentId.IsNull() {
		body, _ = sjson.Set(body, "agentId", data.AgentId.ValueString())
	}
	if !data.TopicName.IsNull() {
		body, _ = sjson.Set(body, "topicName", data.TopicName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DynamicObject) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("objectType"); value.Exists() {
		data.ObjectType = types.StringValue(value.String())
	} else {
		data.ObjectType = types.StringNull()
	}
	if value := res.Get("agentId"); value.Exists() {
		data.AgentId = types.StringValue(value.String())
	} else {
		data.AgentId = types.StringNull()
	}
	if value := res.Get("topicName"); value.Exists() {
		data.TopicName = types.StringValue(value.String())
	} else {
		data.TopicName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DynamicObject) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("objectType"); value.Exists() && !data.ObjectType.IsNull() {
		data.ObjectType = types.StringValue(value.String())
	} else {
		data.ObjectType = types.StringNull()
	}
	if value := res.Get("agentId"); value.Exists() && !data.AgentId.IsNull() {
		data.AgentId = types.StringValue(value.String())
	} else {
		data.AgentId = types.StringNull()
	}
	if value := res.Get("topicName"); value.Exists() && !data.TopicName.IsNull() {
		data.TopicName = types.StringValue(value.String())
	} else {
		data.TopicName = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DynamicObject) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
