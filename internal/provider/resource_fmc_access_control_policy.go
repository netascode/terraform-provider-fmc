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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &AccessControlPolicyResource{}
	_ resource.ResourceWithImportState = &AccessControlPolicyResource{}
)

func NewAccessControlPolicyResource() resource.Resource {
	return &AccessControlPolicyResource{}
}

type AccessControlPolicyResource struct {
	client *fmc.Client
}

func (r *AccessControlPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_control_policy"
}

func (r *AccessControlPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an Access Control Policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the access control policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the default action to take when none of the rules meet the conditions.").AddStringEnumDescription("BLOCK", "TRUST", "PERMIT", "NETWORK_DISCOVERY", "INHERIT_FROM_PARENT").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BLOCK", "TRUST", "PERMIT", "NETWORK_DISCOVERY", "INHERIT_FROM_PARENT"),
				},
			},
			"default_action_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action ID.").String,
				Computed:            true,
			},
			"default_action_log_begin": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicating whether the device will log events at the beginning of the connection.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_action_log_end": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicating whether the device will log events at the end of the connection.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_action_send_events_to_fmc": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicating whether the device will send events to the Firepower Management Center event viewer.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_action_send_syslog": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicating whether the device will send events to a syslog server.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_action_syslog_config_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the syslog config. Can be set only when default_action_send_syslog is true and either default_action_log_begin or default_action_log_end is true. If not set, the default policy syslog configuration in Access Control Logging applies.").String,
				Optional:            true,
			},
			"default_action_syslog_severity": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Override the Severity of syslog alerts.").AddStringEnumDescription("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING"),
				},
			},
			"default_action_snmp_config_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the SNMP alert. Can be set only when either default_action_log_begin or default_action_log_end is true.").String,
				Optional:            true,
			},
			"default_action_intrusion_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the existing intrusion policy (e.g. fmc_intrusion_policy.example.id). Cannot be set when default action is BLOCK, TRUST, NETWORK_DISCOVERY.").String,
				Optional:            true,
			},
			"categories": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of categories.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier of the category.").String,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User-specified unique string.").String,
							Required:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The section of the policy to which the category belongs. Categories must be ordered so that entire section 'mandatory' comes above the section 'default'. If you use inheritance, the mandatory section applies before child policy's own rules, while the default section applies after child policy's own rules.").AddStringEnumDescription("default", "mandatory").AddDefaultValueDescription("default").String,
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("default", "mandatory"),
							},
							Default: stringdefault.StaticString("default"),
						},
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of rules. Rules must be sorted in the order of the corresponding categories, if they have `category_name`. Uncategorized non-mandatory rules must be below all other rules. The first matching rule is selected. Except for MONITOR rules, the system does not continue to evaluate traffic against additional rules after that traffic matches a rule.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Unique identifier (UUID) of the access rule.").String,
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("What to do when the conditions defined by the rule are met.").AddStringEnumDescription("ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE", "BLOCK_RESET_INTERACTIVE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE", "BLOCK_RESET_INTERACTIVE"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User-specified unique string.").String,
							Required:            true,
						},
						"category_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the category that owns this rule (a `name` from `categories` list).").String,
							Optional:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The section of the policy to which the rule belongs. Can only be used when the `category_name` is null. Rules must be ordered so that entire section 'mandatory' comes above the section 'default'. Null value means 'default'. If you use inheritance, the mandatory section applies before child policy's own rules, while the default section applies after child policy's own rules.").AddStringEnumDescription("default", "mandatory").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("default", "mandatory"),
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the access rule is in effect (true) or not (false). Default is true.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"source_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_network_literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destinations of traffic (literally specified).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
								},
							},
						},
						"source_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent sources of traffic (fmc_network, fmc_host, ...).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_network.this.id, etc.).").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object (such as fmc_network.this.type, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_network_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent destinations of traffic (fmc_network, fmc_host, ...).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_network.this.id, etc.).").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object (such as fmc_network.this.type, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"source_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent dynamic sources of traffic (fmc_dynamic_object).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_dynamic_object.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_dynamic_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects that represent dynamic destinations of traffic (fmc_dynamic_object).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_dynamic_object.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"source_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing source ports associated with the rule (fmc_port or fmc_port_group).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_port.this.id, fmc_port_group.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_port_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing destination ports associated with the rule (fmc_port or fmc_port_group).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_port.this.id, fmc_port_group.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"source_security_group_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the source Security Group Tags (fmc_security_group_tag - part of the dynamic attributes).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_security_group_tag.this.id, etc.).").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object (such as fmc_security_group_tag.this.type, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_security_group_tag_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the destination Security Group Tags (fmc_security_group_tag - part of the dynamic attributes).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_security_group_tag.this.id, etc.).").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the object (such as fmc_security_group_tag.this.type, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"source_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing source security zones associated with the access rule (fmc_security_zone).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_security_zone.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"destination_zones": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing destination security zones associated with the access rule (fmc_security_zone).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_security_zone.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"url_objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the URLs associated with the rule (fmc_url or fmc_url_group).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_url.this.id, fmc_url_group.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"url_categories": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects representing the URL Categories associated with the rule (fmc_url_category).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_url_category.this.id, etc.).").String,
										Optional:            true,
									},
									"reputation": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Reputation applicable to the category.").AddStringEnumDescription("ANY_EXCEPT_UNKNOWN", "TRUSTED", "FAVORABLE", "NEUTRAL", "QUESTIONABLE", "UNTRUSTED", "ANY_AND_UNKNOWN", "TRUSTED_AND_UNKNOWN", "FAVORABLE_AND_UNKNOWN", "NEUTRAL_AND_UNKNOWN", "QUESTIONABLE_AND_UNKNOWN", "UNTRUSTED_AND_UNKNOWN").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ANY_EXCEPT_UNKNOWN", "TRUSTED", "FAVORABLE", "NEUTRAL", "QUESTIONABLE", "UNTRUSTED", "ANY_AND_UNKNOWN", "TRUSTED_AND_UNKNOWN", "FAVORABLE_AND_UNKNOWN", "NEUTRAL_AND_UNKNOWN", "QUESTIONABLE_AND_UNKNOWN", "UNTRUSTED_AND_UNKNOWN"),
										},
									},
								},
							},
						},
						"log_begin": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will log events at the beginning of the connection. If 'MONITOR' action is selected for access rule, log_begin must be false or absent.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"log_end": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will log events at the end of the connection. If 'MONITOR' action is selected for access rule, log_end must be true.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"log_files": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will log file events.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"send_events_to_fmc": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will send events to the Firepower Management Center event viewer. If 'MONITOR' action is selected for access rule, send_events_to_fmc must be true.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"send_syslog": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the alerts associated with the access rule are sent to syslog.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"syslog_config_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("UUID of the syslog config. Can be set only when send_syslog is true and either log_begin or log_end is true. If not set, the default policy syslog configuration in Access Control Logging applies.").String,
							Optional:            true,
						},
						"syslog_severity": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override the Severity of syslog alerts.").AddStringEnumDescription("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING"),
							},
						},
						"snmp_config_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("UUID of the SNMP alert associated with the access rule. Can be set only when either log_begin or log_end is true.").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User-specified string.").String,
							Optional:            true,
						},
						"file_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier (UUID) of the File Policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.").String,
							Optional:            true,
						},
						"intrusion_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier (UUID) of the fmc_intrusion_policy for the rule action. Cannot be set when action is BLOCK, BLOCK_RESET, TRUST, MONITOR.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessControlPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *AccessControlPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccessControlPolicy

	// Read plan
	plan, diags := NewValidAccessControlPolicy(ctx, req.Plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AccessControlPolicy{})
	bodyCats := gjson.Parse(body).Get("dummy_categories").Array()
	body, _ = sjson.Delete(body, "dummy_categories")
	bodyRules := gjson.Parse(body).Get("dummy_rules").Array()
	body, _ = sjson.Delete(body, "dummy_rules")
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	read, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, read.String()))

		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}
		return
	}

	plan.fromBodyUnknowns(ctx, read)

	err = r.createCatsAt(ctx, plan, bodyCats, 0, &plan, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())

		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}

		return
	}

	err = r.createRulesAt(ctx, plan, bodyRules, 0, &plan, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())

		res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}

		return
	}

	diags = resp.State.Set(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *AccessControlPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccessControlPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	resGet, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	resCats, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/categories?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	resRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/accessrules?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, resGet.String()))
		return
	}

	s := resGet.String()

	replaceCats := resCats.Get("items").String()
	if replaceCats == "" {
		replaceCats = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_categories", replaceCats)

	replaceRules := resRules.Get("items").String()
	if replaceRules == "" {
		replaceRules = "[]"
	}
	s, _ = sjson.SetRaw(s, "dummy_rules", replaceRules)

	res := gjson.Parse(s)

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}
	state.adjustFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *AccessControlPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state AccessControlPolicy

	// Read plan
	plan, diags := NewValidAccessControlPolicy(ctx, req.Plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)

	bodyCats := gjson.Parse(body).Get("dummy_categories").Array()
	body, _ = sjson.Delete(body, "dummy_categories")

	bodyRules := gjson.Parse(body).Get("dummy_rules")
	body, _ = sjson.Delete(body, "dummy_rules")

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.fromBodyUnknowns(ctx, res)

	keptCats, keptRules := r.countKept(ctx, state, plan)

	// Most of attribs are set as planned, except Rules and Categories which we'll do below.
	existingRules, existingCategories := state.Rules, state.Categories
	state = plan
	state.Rules, state.Categories = existingRules, existingCategories

	// Whether we return error or success, the `state` var will be actually persisted by terraform:
	defer func() {
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
	}()

	err = r.truncateRulesAt(ctx, &state, keptRules, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	err = r.truncateCatsAt(ctx, &state, keptCats, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	if len(plan.Categories) == 0 {
		state.Categories = plan.Categories
	}

	if len(plan.Rules) == 0 {
		state.Rules = plan.Rules
	}

	err = r.createCatsAt(ctx, plan, bodyCats, keptCats, &state, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	err = r.createRulesAt(ctx, plan, bodyRules.Array(), keptRules, &state, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))
}

// countKept compares the state with the plan starting from index 0, and returns:
//
// how many categories to keep: they remain identical as to content and order
//
// how many rules to keep:
//
// - kept rules must belong to some category that is itself kept,
//
// - and must themselves remain identical as to id, content, and order
func (r *AccessControlPolicyResource) countKept(ctx context.Context, state, plan AccessControlPolicy) (int, int) {
	return 0, 0 // TODO
}

func (r *AccessControlPolicyResource) truncateRulesAt(ctx context.Context, state *AccessControlPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	var b strings.Builder
	var bulks []string
	var counts []int
	count := 0

	for i := kept; i < len(state.Rules); i++ {
		if b.Len() != 0 {
			b.WriteString(",")
		}
		b.WriteString(state.Rules[i].Id.ValueString())
		count++
		if b.Len() >= 3700-2 {
			bulks = append(bulks, b.String())
			b.Reset()
			counts = append(counts, count)
			count = 0
		}
	}
	if b.Len() > 0 {
		bulks = append(bulks, b.String())
		counts = append(counts, count)
	}

	defer func() {
		// Apparently, the bulk DELETE has a race. Stabilize:
		time.Sleep(2 * time.Second)
	}()

	for i, bulk := range bulks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+
			"/accessrules?bulk=true&filter=ids:"+url.QueryEscape(bulk), reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to bulk-delete rules, got error: %v, %s", err, res.String())
		}

		state.Rules = slices.Delete(state.Rules, kept, kept+counts[i])
	}

	return nil
}

func (r *AccessControlPolicyResource) truncateCatsAt(ctx context.Context, state *AccessControlPolicy, kept int, reqMods ...func(*fmc.Req)) error {
	for i := len(state.Categories) - 1; i >= kept; i-- {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+
			"/categories/"+url.QueryEscape(state.Categories[i].Id.ValueString()), reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to delete a category, got error: %v, %s", err, res)
		}

		state.Categories[i] = AccessControlPolicyCategories{}
		state.Categories = state.Categories[:i]
	}
	return nil
}

func (r *AccessControlPolicyResource) createCatsAt(ctx context.Context, plan AccessControlPolicy, body []gjson.Result, startIndex int, state *AccessControlPolicy, reqMods ...func(*fmc.Req)) error {
	for i := startIndex; i < len(plan.Categories); i++ {
		cat := body[i].String()
		cat, _ = sjson.Delete(cat, "id")
		cat, _ = sjson.Delete(cat, "metadata.section")
		params := ""
		if s := plan.Categories[i].Section.ValueString(); s != "" {
			params = "?section=" + url.QueryEscape(s)
		}
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+
			"/categories"+params, cat, reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to create a category (POST), got error: %v, %s", err, res)
		}

		item := plan.Categories[i]
		item.Id = types.StringValue(res.Get("id").String())

		if len(state.Categories) <= i {
			state.Categories = append(state.Categories, item)
		} else {
			state.Categories[i] = item
		}
	}

	return nil
}

// createRulesAt creates rules plan.Rules[startIndex:] (from startIndex to the end).
// The `body` should be a marshalled `plan.Rules`.
// Whether it succeeds fully or partially, it takes whatever has been really created and saves in the `state`.
// The `state` and `&plan` might be either the same value or different.
func (r *AccessControlPolicyResource) createRulesAt(ctx context.Context, plan AccessControlPolicy, body []gjson.Result, startIndex int, state *AccessControlPolicy, reqMods ...func(*fmc.Req)) error {
	for i := startIndex; i < len(body); i++ {
		bulk := `{"dummy_rules":[]}`
		j := i
		head := plan.Rules[i]
		for ; i < len(body); i++ {
			if !head.CategoryName.Equal(plan.Rules[i].CategoryName) || head.GetSection() != plan.Rules[i].GetSection() {
				i--
				break
			}
			rule := body[i].String()
			rule, _ = sjson.Delete(rule, "id")
			rule, _ = sjson.Delete(rule, "metadata.category")
			rule, _ = sjson.Delete(rule, "metadata.section")

			bulk, _ = sjson.SetRaw(bulk, "dummy_rules.-1", rule)
		}

		param := "?bulk=true"
		if cat := head.CategoryName.ValueString(); cat != "" {
			param += "&category=" + url.QueryEscape(cat)
		} else if s := head.GetSection(); s != "default" {
			param += "&section=" + url.QueryEscape(s)
		}
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/accessrules"+param,
			gjson.Parse(bulk).Get("dummy_rules").String(),
			reqMods...)
		if err != nil {
			return fmt.Errorf("Failed to configure object (POST), got error: %s, %s", err, res.String())
		}

		for _, v := range res.Get("items").Array() {
			if v.Get("name").String() == "" {
				tflog.Debug(ctx, fmt.Sprintf("ignoring a bogus item in POST reply, it can happen: %s", v))
				continue
			}

			item := plan.Rules[j]
			item.Id = types.StringValue(v.Get("id").String())

			if len(state.Rules) <= j {
				state.Rules = append(state.Rules, item)
			} else {
				state.Rules[j] = item
			}

			j++
		}
	}

	return nil
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *AccessControlPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccessControlPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *AccessControlPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
