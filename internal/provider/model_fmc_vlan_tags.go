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
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VLANTags struct {
	Id     types.String             `tfsdk:"id"`
	Domain types.String             `tfsdk:"domain"`
	Items  map[string]VLANTagsItems `tfsdk:"items"`
}

type VLANTagsItems struct {
	Id          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Overridable types.Bool   `tfsdk:"overridable"`
	StartTag    types.String `tfsdk:"start_tag"`
	EndTag      types.String `tfsdk:"end_tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VLANTags) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/vlantags"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VLANTags) toBody(ctx context.Context, state VLANTags) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Overridable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overridable", item.Overridable.ValueBool())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "VlanTag")
			if !item.StartTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "data.startTag", item.StartTag.ValueString())
			}
			if !item.EndTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "data.endTag", item.EndTag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VLANTags) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
		}
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("overridable"); value.Exists() {
			data.Overridable = types.BoolValue(value.Bool())
		} else {
			data.Overridable = types.BoolNull()
		}
		if value := res.Get("data.startTag"); value.Exists() {
			data.StartTag = types.StringValue(value.String())
		} else {
			data.StartTag = types.StringNull()
		}
		if value := res.Get("data.endTag"); value.Exists() {
			data.EndTag = types.StringValue(value.String())
		} else {
			data.EndTag = types.StringNull()
		}
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VLANTags) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
			data.Overridable = types.BoolValue(value.Bool())
		} else {
			data.Overridable = types.BoolNull()
		}
		if value := res.Get("data.startTag"); value.Exists() && !data.StartTag.IsNull() {
			data.StartTag = types.StringValue(value.String())
		} else {
			data.StartTag = types.StringNull()
		}
		if value := res.Get("data.endTag"); value.Exists() && !data.EndTag.IsNull() {
			data.EndTag = types.StringValue(value.String())
		} else {
			data.EndTag = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VLANTags) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *VLANTags) Clone() VLANTags {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data VLANTags) toBodyNonBulk(ctx context.Context, state VLANTags) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to chage. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk
