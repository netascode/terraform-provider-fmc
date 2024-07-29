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
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkGroups struct {
	Id     types.String                  `tfsdk:"id"`
	Domain types.String                  `tfsdk:"domain"`
	Items  map[string]NetworkGroupsItems `tfsdk:"items"`
}

type NetworkGroupsItems struct {
	Id          types.String                 `tfsdk:"id"`
	Description types.String                 `tfsdk:"description"`
	Overridable types.Bool                   `tfsdk:"overridable"`
	GroupNames  types.Set                    `tfsdk:"group_names"`
	Objects     []NetworkGroupsItemsObjects  `tfsdk:"objects"`
	Literals    []NetworkGroupsItemsLiterals `tfsdk:"literals"`
}

type NetworkGroupsItemsObjects struct {
	Id types.String `tfsdk:"id"`
}
type NetworkGroupsItemsLiterals struct {
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroups) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/networkgroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkGroups) toBody(ctx context.Context, state NetworkGroups) string {
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
			if !item.GroupNames.IsNull() {
				var values []string
				item.GroupNames.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "group_names", values)
			}
			if len(item.Objects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "objects", []interface{}{})
				for _, childItem := range item.Objects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "objects.-1", itemChildBody)
				}
			}
			if len(item.Literals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "literals", []interface{}{})
				for _, childItem := range item.Literals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "literals.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroups) fromBody(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkGroups) fromBodyPartial(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkGroups) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
					if val.Id.ValueString() == v.Get("id").String() {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if data.Items[i].Id.IsUnknown() {
			v := data.Items[i]
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

func (data *NetworkGroups) Clone() NetworkGroups {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}
