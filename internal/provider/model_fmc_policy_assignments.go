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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type PolicyAssignments struct {
	Id       types.String               `tfsdk:"id"`
	Domain   types.String               `tfsdk:"domain"`
	Name     types.String               `tfsdk:"name"`
	PolicyId types.String               `tfsdk:"policy_id"`
	Targets  []PolicyAssignmentsTargets `tfsdk:"targets"`
}

type PolicyAssignmentsTargets struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PolicyAssignments) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/assignment/policyassignments"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PolicyAssignments) toBody(ctx context.Context, state PolicyAssignments) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "PolicyAssignment")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "policy.name", data.Name.ValueString())
	}
	if !data.PolicyId.IsNull() {
		body, _ = sjson.Set(body, "policy.id", data.PolicyId.ValueString())
	}
	if len(data.Targets) > 0 {
		body, _ = sjson.Set(body, "targets", []interface{}{})
		for _, item := range data.Targets {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "targets.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PolicyAssignments) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("policy.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policy.id"); value.Exists() {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	if value := res.Get("targets"); value.Exists() {
		data.Targets = make([]PolicyAssignmentsTargets, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyAssignmentsTargets{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).Targets = append((*parent).Targets, data)
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
func (data *PolicyAssignments) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("policy.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policy.id"); value.Exists() && !data.PolicyId.IsNull() {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	for i := 0; i < len(data.Targets); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Targets[i].Id.ValueString()}

		parent := &data
		data := (*parent).Targets[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("targets").ForEach(
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
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Targets[%d] = %+v",
				i,
				(*parent).Targets[i],
			))
			(*parent).Targets = slices.Delete((*parent).Targets, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).Targets[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PolicyAssignments) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
