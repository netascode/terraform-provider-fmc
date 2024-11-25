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
	_ resource.Resource                = &DecryptionPolicyResource{}
	_ resource.ResourceWithImportState = &DecryptionPolicyResource{}
)

func NewDecryptionPolicyResource() resource.Resource {
	return &DecryptionPolicyResource{}
}

type DecryptionPolicyResource struct {
	client *fmc.Client
}

func (r *DecryptionPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_decryption_policy"
}

func (r *DecryptionPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Decryption Policy.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("User-created name of the resource.").String,
				Required:            true,
			},
			"decryption_errors": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether decryption errors are allowed.").AddStringEnumDescription("BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"compressed_session": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the session is compressed.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"ssl_v2_session": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether SSLv2 is enabled.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"unknown_cipher_suite": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether unknown cipher suites are allowed.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"unsupported_cipher_suite": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether unsupported cipher suites are allowed.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"session_not_cached": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether sessions are cached.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"handshake_errors": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether handshake errors are allowed.").AddStringEnumDescription("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("INHERIT_DEFAULT_ACTION", "DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"policy_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates the default action: PERMIT or DENY.").AddStringEnumDescription("DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DO_NOT_DECRYPT", "BLOCK", "BLOCK_WITH_RESET"),
				},
			},
			"event_log_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates the default action: PERMIT or DENY.").AddStringEnumDescription("LOG_FLOW_END", "LOG_NONE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("LOG_FLOW_END", "LOG_NONE"),
				},
			},
			"disallow_untrusted_issuer_resign": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether untrusted issuers are disallowed.").String,
				Optional:            true,
			},
			"strip_http3": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether HTTP/3 is stripped.").String,
				Optional:            true,
			},
			"tls13_decryption": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether TLS 1.3 decryption is enabled.").String,
				Optional:            true,
			},
			"quic_decryption": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether QUIC decryption is enabled.").String,
				Optional:            true,
			},
			"adaptive_probe": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether adaptive probing is enabled.").String,
				Optional:            true,
			},
			"block_extensions": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether extensions are blocked.").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"log_end": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will log events at the end of the connection.").String,
				Optional:            true,
			},
			"send_events": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the device will send events.").String,
				Optional:            true,
			},
			"trusted_cas": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of trusted CAs.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ca_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the trusted CA.").String,
							Optional:            true,
						},
						"ca_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the trusted CA.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DecryptionPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DecryptionPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DecryptionPolicy

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
	body := plan.toBody(ctx, DecryptionPolicy{})
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

func (r *DecryptionPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DecryptionPolicy

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

func (r *DecryptionPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DecryptionPolicy

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

func (r *DecryptionPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DecryptionPolicy

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

func (r *DecryptionPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
