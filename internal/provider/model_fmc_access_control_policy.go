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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AccessControlPolicy struct {
	Id                           types.String                    `tfsdk:"id"`
	Domain                       types.String                    `tfsdk:"domain"`
	Name                         types.String                    `tfsdk:"name"`
	Description                  types.String                    `tfsdk:"description"`
	DefaultAction                types.String                    `tfsdk:"default_action"`
	DefaultActionId              types.String                    `tfsdk:"default_action_id"`
	DefaultActionLogBegin        types.Bool                      `tfsdk:"default_action_log_begin"`
	DefaultActionLogEnd          types.Bool                      `tfsdk:"default_action_log_end"`
	DefaultActionSendEventsToFmc types.Bool                      `tfsdk:"default_action_send_events_to_fmc"`
	DefaultActionSendSyslog      types.Bool                      `tfsdk:"default_action_send_syslog"`
	Categories                   []AccessControlPolicyCategories `tfsdk:"categories"`
	Rules                        []AccessControlPolicyRules      `tfsdk:"rules"`
}

type AccessControlPolicyCategories struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

type AccessControlPolicyRules struct {
	Id                    types.String                                    `tfsdk:"id"`
	Action                types.String                                    `tfsdk:"action"`
	Name                  types.String                                    `tfsdk:"name"`
	CategoryName          types.String                                    `tfsdk:"category_name"`
	Enabled               types.Bool                                      `tfsdk:"enabled"`
	SourceNetworkLiterals []AccessControlPolicyRulesSourceNetworkLiterals `tfsdk:"source_network_literals"`
	SourceNetworkObjects  types.List                                      `tfsdk:"source_network_objects"`
}

type AccessControlPolicyRulesSourceNetworkLiterals struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AccessControlPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AccessControlPolicy) toBody(ctx context.Context, state AccessControlPolicy) string {
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
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.action", data.DefaultAction.ValueString())
	}
	if state.DefaultActionId.ValueString() != "" {
		body, _ = sjson.Set(body, "defaultAction.id", state.DefaultActionId.ValueString())
	}
	if !data.DefaultActionLogBegin.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logBegin", data.DefaultActionLogBegin.ValueBool())
	}
	if !data.DefaultActionLogEnd.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logEnd", data.DefaultActionLogEnd.ValueBool())
	}
	if !data.DefaultActionSendEventsToFmc.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.sendEventsToFMC", data.DefaultActionSendEventsToFmc.ValueBool())
	}
	if !data.DefaultActionSendSyslog.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.enableSyslog", data.DefaultActionSendSyslog.ValueBool())
	}
	if len(data.Categories) > 0 {
		body, _ = sjson.Set(body, "dummy_categories", []interface{}{})
		for _, item := range data.Categories {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_categories.-1", itemBody)
		}
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "dummy_rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.CategoryName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "category_name", item.CategoryName.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []interface{}{})
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals.-1", itemChildBody)
				}
			}
			if !item.SourceNetworkObjects.IsNull() {
				var values []string
				item.SourceNetworkObjects.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", values)
			}
			body, _ = sjson.SetRaw(body, "dummy_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AccessControlPolicy) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultAction.action"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogBegin = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogEnd = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendEventsToFmc = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.enableSyslog"); value.Exists() {
		data.DefaultActionSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendSyslog = types.BoolValue(false)
	}
	if value := res.Get("dummy_categories"); value.Exists() {
		data.Categories = make([]AccessControlPolicyCategories, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessControlPolicyCategories{}
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
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
			data.Categories = append(data.Categories, item)
			return true
		})
	}
	if value := res.Get("dummy_rules"); value.Exists() {
		data.Rules = make([]AccessControlPolicyRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessControlPolicyRules{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			} else {
				item.Action = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("category_name"); cValue.Exists() {
				item.CategoryName = types.StringValue(cValue.String())
			} else {
				item.CategoryName = types.StringNull()
			}
			if cValue := v.Get("enabled"); cValue.Exists() {
				item.Enabled = types.BoolValue(cValue.Bool())
			} else {
				item.Enabled = types.BoolValue(true)
			}
			if cValue := v.Get("sourceNetworks.literals"); cValue.Exists() {
				item.SourceNetworkLiterals = make([]AccessControlPolicyRulesSourceNetworkLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceNetworkLiterals{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.SourceNetworkLiterals = append(item.SourceNetworkLiterals, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceNetworks.objects"); cValue.Exists() {
				item.SourceNetworkObjects = helpers.GetStringList(cValue.Array())
			} else {
				item.SourceNetworkObjects = types.ListNull(types.StringType)
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

func (data *AccessControlPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultAction.action"); value.Exists() && !data.DefaultAction.IsNull() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() && !data.DefaultActionLogBegin.IsNull() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else if data.DefaultActionLogBegin.ValueBool() != false {
		data.DefaultActionLogBegin = types.BoolNull()
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() && !data.DefaultActionLogEnd.IsNull() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else if data.DefaultActionLogEnd.ValueBool() != false {
		data.DefaultActionLogEnd = types.BoolNull()
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() && !data.DefaultActionSendEventsToFmc.IsNull() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else if data.DefaultActionSendEventsToFmc.ValueBool() != false {
		data.DefaultActionSendEventsToFmc = types.BoolNull()
	}
	if value := res.Get("defaultAction.enableSyslog"); value.Exists() && !data.DefaultActionSendSyslog.IsNull() {
		data.DefaultActionSendSyslog = types.BoolValue(value.Bool())
	} else if data.DefaultActionSendSyslog.ValueBool() != false {
		data.DefaultActionSendSyslog = types.BoolNull()
	}

	resLen := len(res.Get("dummy_categories").Array())
	for i := len(data.Categories); i < resLen; i++ {
		data.Categories = append(data.Categories, AccessControlPolicyCategories{})
	}
	if len(data.Categories) > resLen {
		data.Categories = data.Categories[:resLen]
	}

	for i := range data.Categories {
		r := res.Get(fmt.Sprintf("dummy_categories.%d", i))
		if value := r.Get("id"); value.Exists() {
			data.Categories[i].Id = types.StringValue(value.String())
		} else {
			data.Categories[i].Id = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.Categories[i].Name.IsNull() {
			data.Categories[i].Name = types.StringValue(value.String())
		} else {
			data.Categories[i].Name = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.Categories[i].Description.IsNull() {
			data.Categories[i].Description = types.StringValue(value.String())
		} else {
			data.Categories[i].Description = types.StringNull()
		}
		// if value := r.Get("mandatory_section"); value.Exists() && !data.Categories[i].MandatorySection.IsNull() {
		// 	data.Categories[i].MandatorySection = types.BoolValue(value.Bool())
		// } else if data.Categories[i].MandatorySection.ValueBool() != true {
		// 	data.Categories[i].MandatorySection = types.BoolNull()
		// }
	}

	resLen = len(res.Get("dummy_rules").Array())
	for i := len(data.Rules); i < resLen; i++ {
		data.Rules = append(data.Rules, AccessControlPolicyRules{})
	}
	if len(data.Rules) > resLen {
		data.Rules = data.Rules[:resLen]
	}

	for i := range data.Rules {
		r := res.Get(fmt.Sprintf("dummy_rules.%d", i))
		if value := r.Get("action"); value.Exists() && !data.Rules[i].Action.IsNull() {
			data.Rules[i].Action = types.StringValue(value.String())
		} else {
			data.Rules[i].Action = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.Rules[i].Name.IsNull() {
			data.Rules[i].Name = types.StringValue(value.String())
		} else {
			data.Rules[i].Name = types.StringNull()
		}
		if value := r.Get("enabled"); value.Exists() && !data.Rules[i].Enabled.IsNull() {
			data.Rules[i].Enabled = types.BoolValue(value.Bool())
		} else if data.Rules[i].Enabled.ValueBool() != true {
			data.Rules[i].Enabled = types.BoolNull()
		}
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AccessControlPolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.DefaultAction.IsNull() {
		return false
	}
	if !data.DefaultActionId.IsNull() {
		return false
	}
	if !data.DefaultActionLogBegin.IsNull() {
		return false
	}
	if !data.DefaultActionLogEnd.IsNull() {
		return false
	}
	if !data.DefaultActionSendEventsToFmc.IsNull() {
		return false
	}
	if !data.DefaultActionSendSyslog.IsNull() {
		return false
	}
	if len(data.Categories) > 0 {
		return false
	}
	if len(data.Rules) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
