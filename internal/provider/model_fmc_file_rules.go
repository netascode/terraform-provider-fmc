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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FileRules struct {
	Id             types.String              `tfsdk:"id"`
	Domain         types.String              `tfsdk:"domain"`
	FilePolicyId   types.String              `tfsdk:"file_policy_id"`
	Name           types.String              `tfsdk:"name"`
	Action         types.String              `tfsdk:"action"`
	Protocol       types.String              `tfsdk:"protocol"`
	Direction      types.String              `tfsdk:"direction"`
	FileCategories []FileRulesFileCategories `tfsdk:"file_categories"`
	FileTypes      []FileRulesFileTypes      `tfsdk:"file_types"`
}

type FileRulesFileCategories struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type FileRulesFileTypes struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FileRules) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/filepolicies/%v/filerules", url.QueryEscape(data.FilePolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FileRules) toBody(ctx context.Context, state FileRules) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, "action", data.Action.ValueString())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "protocol", data.Protocol.ValueString())
	}
	if !data.Direction.IsNull() {
		body, _ = sjson.Set(body, "direction", data.Direction.ValueString())
	}
	if len(data.FileCategories) > 0 {
		body, _ = sjson.Set(body, "fileCategories", []interface{}{})
		for _, item := range data.FileCategories {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "FileCategory")
			body, _ = sjson.SetRaw(body, "fileCategories.-1", itemBody)
		}
	}
	if len(data.FileTypes) > 0 {
		body, _ = sjson.Set(body, "fileTypes", []interface{}{})
		for _, item := range data.FileTypes {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "FileType")
			body, _ = sjson.SetRaw(body, "fileTypes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FileRules) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("direction"); value.Exists() {
		data.Direction = types.StringValue(value.String())
	} else {
		data.Direction = types.StringNull()
	}
	if value := res.Get("fileCategories"); value.Exists() {
		data.FileCategories = make([]FileRulesFileCategories, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FileRulesFileCategories{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.FileCategories = append(data.FileCategories, item)
			return true
		})
	}
	if value := res.Get("fileTypes"); value.Exists() {
		data.FileTypes = make([]FileRulesFileTypes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FileRulesFileTypes{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.FileTypes = append(data.FileTypes, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FileRules) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("direction"); value.Exists() && !data.Direction.IsNull() {
		data.Direction = types.StringValue(value.String())
	} else {
		data.Direction = types.StringNull()
	}
	for i := 0; i < len(data.FileCategories); i++ {
		keys := [...]string{"id", "name"}
		keyValues := [...]string{data.FileCategories[i].Id.ValueString(), data.FileCategories[i].Name.ValueString()}

		var r gjson.Result
		res.Get("fileCategories").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing data.FileCategories[%d] = %+v",
				i,
				data.FileCategories[i],
			))
			data.FileCategories = slices.Delete(data.FileCategories, i, i+1)
			i--

			continue
		}
		if value := r.Get("id"); value.Exists() && !data.FileCategories[i].Id.IsNull() {
			data.FileCategories[i].Id = types.StringValue(value.String())
		} else {
			data.FileCategories[i].Id = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.FileCategories[i].Name.IsNull() {
			data.FileCategories[i].Name = types.StringValue(value.String())
		} else {
			data.FileCategories[i].Name = types.StringNull()
		}
	}
	for i := 0; i < len(data.FileTypes); i++ {
		keys := [...]string{"id", "name"}
		keyValues := [...]string{data.FileTypes[i].Id.ValueString(), data.FileTypes[i].Name.ValueString()}

		var r gjson.Result
		res.Get("fileTypes").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing data.FileTypes[%d] = %+v",
				i,
				data.FileTypes[i],
			))
			data.FileTypes = slices.Delete(data.FileTypes, i, i+1)
			i--

			continue
		}
		if value := r.Get("id"); value.Exists() && !data.FileTypes[i].Id.IsNull() {
			data.FileTypes[i].Id = types.StringValue(value.String())
		} else {
			data.FileTypes[i].Id = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.FileTypes[i].Name.IsNull() {
			data.FileTypes[i].Name = types.StringValue(value.String())
		} else {
			data.FileTypes[i].Name = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FileRules) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
