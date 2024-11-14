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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &PrefilterPolicyResource{}
	_ resource.ResourceWithImportState = &PrefilterPolicyResource{}
)

func NewPrefilterPolicyResource() resource.Resource {
	return &PrefilterPolicyResource{}
}

type PrefilterPolicyResource struct {
	client *fmc.Client
}

func (r *PrefilterPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_prefilter_policy"
}

func (r *PrefilterPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Prefilter Policy.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("The name of the prefilter policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the default action to take when none of the rules meet the conditions.").AddStringEnumDescription("BLOCK_TUNNELS", "ANALYZE_TUNNELS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BLOCK_TUNNELS", "ANALYZE_TUNNELS"),
				},
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
			"default_action_syslog_config_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the syslog config. Can be set only when either default_action_log_begin or default_action_log_end is true.").String,
				Optional:            true,
			},
			"default_action_snmp_config_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID of the SNMP alert. Can be set only when either default_action_log_begin or default_action_log_end is true.").String,
				Optional:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of rules. Rules must be sorted in the order of the corresponding categories, if they have `category_name`. Uncategorized non-mandatory rules must be below all other rules. The first matching rule is selected. Except for MONITOR rules, the system does not continue to evaluate traffic against additional rules after that traffic matches a rule.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Unique identifier (UUID) of the prefilter rule.").String,
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("What to do when the conditions defined by the rule are met.").AddStringEnumDescription("FASTPATH", "ANALYZE", "BLOCK").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("FASTPATH", "ANALYZE", "BLOCK"),
							},
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the rule is prefilter rule or tunnel rule.").AddStringEnumDescription("PREFILTER", "TUNNEL").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PREFILTER", "TUNNEL"),
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the access rule is in effect (true) or not (false). Default is true.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"bidirectional": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the rule is bidirectional.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
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
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtMost(1000),
				},
			},
		},
	}
}

func (r *PrefilterPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *PrefilterPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PrefilterPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
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
	body := plan.toBody(ctx, PrefilterPolicy{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *PrefilterPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PrefilterPolicy

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *PrefilterPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PrefilterPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *PrefilterPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PrefilterPolicy

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

func (r *PrefilterPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
