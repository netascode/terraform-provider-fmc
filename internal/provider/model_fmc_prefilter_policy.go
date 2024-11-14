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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type PrefilterPolicy struct {
	Id                           types.String           `tfsdk:"id"`
	Domain                       types.String           `tfsdk:"domain"`
	Name                         types.String           `tfsdk:"name"`
	Description                  types.String           `tfsdk:"description"`
	DefaultAction                types.String           `tfsdk:"default_action"`
	DefaultActionLogBegin        types.Bool             `tfsdk:"default_action_log_begin"`
	DefaultActionLogEnd          types.Bool             `tfsdk:"default_action_log_end"`
	DefaultActionSendEventsToFmc types.Bool             `tfsdk:"default_action_send_events_to_fmc"`
	DefaultActionSyslogConfigId  types.String           `tfsdk:"default_action_syslog_config_id"`
	DefaultActionSnmpConfigId    types.String           `tfsdk:"default_action_snmp_config_id"`
	Rules                        []PrefilterPolicyRules `tfsdk:"rules"`
}

type PrefilterPolicyRules struct {
	Id              types.String `tfsdk:"id"`
	Action          types.String `tfsdk:"action"`
	RuleType        types.String `tfsdk:"rule_type"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	Bidirectional   types.Bool   `tfsdk:"bidirectional"`
	LogBegin        types.Bool   `tfsdk:"log_begin"`
	LogEnd          types.Bool   `tfsdk:"log_end"`
	SendEventsToFmc types.Bool   `tfsdk:"send_events_to_fmc"`
	SendSyslog      types.Bool   `tfsdk:"send_syslog"`
	SyslogConfigId  types.String `tfsdk:"syslog_config_id"`
	SyslogSeverity  types.String `tfsdk:"syslog_severity"`
	SnmpConfigId    types.String `tfsdk:"snmp_config_id"`
	Description     types.String `tfsdk:"description"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PrefilterPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/prefilterpolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PrefilterPolicy) toBody(ctx context.Context, state PrefilterPolicy) string {
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
	if !data.DefaultActionLogBegin.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logBegin", data.DefaultActionLogBegin.ValueBool())
	}
	if !data.DefaultActionLogEnd.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logEnd", data.DefaultActionLogEnd.ValueBool())
	}
	if !data.DefaultActionSendEventsToFmc.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.sendEventsToFMC", data.DefaultActionSendEventsToFmc.ValueBool())
	}
	if !data.DefaultActionSyslogConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.syslogConfig.id", data.DefaultActionSyslogConfigId.ValueString())
	}
	if !data.DefaultActionSnmpConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.snmpConfig.id", data.DefaultActionSnmpConfigId.ValueString())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "dummy_rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.RuleType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ruleType", item.RuleType.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Bidirectional.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bidirectional", item.Bidirectional.ValueBool())
			}
			if !item.LogBegin.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logBegin", item.LogBegin.ValueBool())
			}
			if !item.LogEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logEnd", item.LogEnd.ValueBool())
			}
			if !item.SendEventsToFmc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendEventsToFMC", item.SendEventsToFmc.ValueBool())
			}
			if !item.SendSyslog.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSyslog", item.SendSyslog.ValueBool())
			}
			if !item.SyslogConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogConfig.id", item.SyslogConfigId.ValueString())
			}
			if !item.SyslogSeverity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogSeverity", item.SyslogSeverity.ValueString())
			}
			if !item.SnmpConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "snmpConfig.id", item.SnmpConfigId.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PrefilterPolicy) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	if value := res.Get("dummy_rules"); value.Exists() {
		data.Rules = make([]PrefilterPolicyRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PrefilterPolicyRules{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("ruleType"); value.Exists() {
				data.RuleType = types.StringValue(value.String())
			} else {
				data.RuleType = types.StringNull()
			}
			if value := res.Get("enabled"); value.Exists() {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolValue(true)
			}
			if value := res.Get("bidirectional"); value.Exists() {
				data.Bidirectional = types.BoolValue(value.Bool())
			} else {
				data.Bidirectional = types.BoolValue(true)
			}
			if value := res.Get("logBegin"); value.Exists() {
				data.LogBegin = types.BoolValue(value.Bool())
			} else {
				data.LogBegin = types.BoolValue(false)
			}
			if value := res.Get("logEnd"); value.Exists() {
				data.LogEnd = types.BoolValue(value.Bool())
			} else {
				data.LogEnd = types.BoolValue(false)
			}
			if value := res.Get("sendEventsToFMC"); value.Exists() {
				data.SendEventsToFmc = types.BoolValue(value.Bool())
			} else {
				data.SendEventsToFmc = types.BoolValue(false)
			}
			if value := res.Get("enableSyslog"); value.Exists() {
				data.SendSyslog = types.BoolValue(value.Bool())
			} else {
				data.SendSyslog = types.BoolValue(false)
			}
			if value := res.Get("syslogConfig.id"); value.Exists() {
				data.SyslogConfigId = types.StringValue(value.String())
			} else {
				data.SyslogConfigId = types.StringNull()
			}
			if value := res.Get("syslogSeverity"); value.Exists() {
				data.SyslogSeverity = types.StringValue(value.String())
			} else {
				data.SyslogSeverity = types.StringNull()
			}
			if value := res.Get("snmpConfig.id"); value.Exists() {
				data.SnmpConfigId = types.StringValue(value.String())
			} else {
				data.SnmpConfigId = types.StringNull()
			}
			(*parent).Rules = append((*parent).Rules, data)
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
func (data *PrefilterPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() && !data.DefaultActionSyslogConfigId.IsNull() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() && !data.DefaultActionSnmpConfigId.IsNull() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	{
		l := len(res.Get("dummy_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_rules array resizing from %d to %d", len(data.Rules), l))
		for i := len(data.Rules); i < l; i++ {
			data.Rules = append(data.Rules, PrefilterPolicyRules{})
		}
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("dummy_rules.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("ruleType"); value.Exists() && !data.RuleType.IsNull() {
			data.RuleType = types.StringValue(value.String())
		} else {
			data.RuleType = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else if data.Enabled.ValueBool() != true {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("bidirectional"); value.Exists() && !data.Bidirectional.IsNull() {
			data.Bidirectional = types.BoolValue(value.Bool())
		} else if data.Bidirectional.ValueBool() != true {
			data.Bidirectional = types.BoolNull()
		}
		if value := res.Get("logBegin"); value.Exists() && !data.LogBegin.IsNull() {
			data.LogBegin = types.BoolValue(value.Bool())
		} else if data.LogBegin.ValueBool() != false {
			data.LogBegin = types.BoolNull()
		}
		if value := res.Get("logEnd"); value.Exists() && !data.LogEnd.IsNull() {
			data.LogEnd = types.BoolValue(value.Bool())
		} else if data.LogEnd.ValueBool() != false {
			data.LogEnd = types.BoolNull()
		}
		if value := res.Get("sendEventsToFMC"); value.Exists() && !data.SendEventsToFmc.IsNull() {
			data.SendEventsToFmc = types.BoolValue(value.Bool())
		} else if data.SendEventsToFmc.ValueBool() != false {
			data.SendEventsToFmc = types.BoolNull()
		}
		if value := res.Get("enableSyslog"); value.Exists() && !data.SendSyslog.IsNull() {
			data.SendSyslog = types.BoolValue(value.Bool())
		} else if data.SendSyslog.ValueBool() != false {
			data.SendSyslog = types.BoolNull()
		}
		if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogConfigId.IsNull() {
			data.SyslogConfigId = types.StringValue(value.String())
		} else {
			data.SyslogConfigId = types.StringNull()
		}
		if value := res.Get("syslogSeverity"); value.Exists() && !data.SyslogSeverity.IsNull() {
			data.SyslogSeverity = types.StringValue(value.String())
		} else {
			data.SyslogSeverity = types.StringNull()
		}
		if value := res.Get("snmpConfig.id"); value.Exists() && !data.SnmpConfigId.IsNull() {
			data.SnmpConfigId = types.StringValue(value.String())
		} else {
			data.SnmpConfigId = types.StringNull()
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PrefilterPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i := range data.Rules {
		r := res.Get(fmt.Sprintf("dummy_rules.%d", i))
		if v := data.Rules[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Rules[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
