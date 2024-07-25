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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	_ resource.Resource                = &ManualNatRulesResource{}
	_ resource.ResourceWithImportState = &ManualNatRulesResource{}
)

func NewManualNatRulesResource() resource.Resource {
	return &ManualNatRulesResource{}
}

type ManualNatRulesResource struct {
	client *fmc.Client
}

func (r *ManualNatRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_manual_nat_rules"
}

func (r *ManualNatRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Manual Nat Rules.").String,

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
			"nat_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the associated nat policy").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("STATIC", "DYNAMIC"),
				},
			},
			"interface_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"fall_through": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"dns": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"route_lookup": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"no_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"net_to_net": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is enabled").String,
				Optional:            true,
			},
			"unidirectional": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is unidirectional").String,
				Optional:            true,
			},
			"original_source_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original source network object").String,
				Required:            true,
			},
			"original_destination_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original destination network object").String,
				Required:            true,
			},
			"original_source_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original source port object").String,
				Optional:            true,
			},
			"original_destination_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of original destination port object").String,
				Optional:            true,
			},
			"translated_destination_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination network object").String,
				Required:            true,
			},
			"translated_source_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated source network object").String,
				Required:            true,
			},
			"translated_destination_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination port object").String,
				Optional:            true,
			},
			"translated_source_port_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of translated source port object").String,
				Optional:            true,
			},
			"source_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone").String,
				Required:            true,
			},
			"destination_interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone").String,
				Required:            true,
			},
			"interface_pat": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"include_reverse": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"round_robin": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"extended_pat": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"flat_port_range": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"block_allocation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"pat_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of pat host object").String,
				Optional:            true,
			},
			"pat_pool_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of pat host object").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the NAT rule").String,
				Optional:            true,
			},
			"interface_in_original_destination": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").String,
				Optional:            true,
			},
		},
	}
}

func (r *ManualNatRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ManualNatRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ManualNatRules

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
	body := plan.toBody(ctx, ManualNatRules{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ManualNatRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ManualNatRules

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

	res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
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

func (r *ManualNatRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ManualNatRules

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ManualNatRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ManualNatRules

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

func (r *ManualNatRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <nat_policy_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("nat_policy_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
