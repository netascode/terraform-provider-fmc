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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &SmartLicenseResource{}
)

func NewSmartLicenseResource() resource.Resource {
	return &SmartLicenseResource{}
}

type SmartLicenseResource struct {
	client *fmc.Client
}

func (r *SmartLicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smart_license"
}

func (r *SmartLicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Smart License.").String,

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
			"registration_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Action to be executed on the smart license.").AddStringEnumDescription("REGISTER", "EVALUATION").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REGISTER", "EVALUATION"),
				},
			},
			"token": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Registration token.").String,
				Optional:            true,
			},
			"registration_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Status of a smart license.").String,
				Computed:            true,
			},
			"force": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set to true to re-register smart license.").String,
				Optional:            true,
			},
		},
	}
}

func (r *SmartLicenseResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

func (r *SmartLicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SmartLicense
	var state SmartLicense

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

	// Read state before create
	res, err := r.client.Get(state.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	state.fromBody(ctx, res.Get("items.0"))

	if state.RegistrationStatus.ValueString() == "EVALUATION" && plan.RegistrationType.ValueString() == "EVALUATION" {
		plan.RegistrationStatus = state.RegistrationStatus
		plan.Id = types.StringValue("")
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	if plan.RegistrationType.ValueString() == "REGISTER" && plan.Token.ValueString() == "" {
		resp.Diagnostics.AddError("Provider Error", "Token required for registration")
		return
	}

	// Create object
	body := plan.toBody(ctx, SmartLicense{})
	res, err = r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	// Read state after create
	res, err = r.client.Get(state.getPath(), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	state.fromBody(ctx, res.Get("items.0"))
	plan.RegistrationStatus = state.RegistrationStatus

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *SmartLicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SmartLicense

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

	res, err := r.client.Get(state.getPath(), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	state.fromBody(ctx, res.Get("items.0"))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *SmartLicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SmartLicense

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

	if state.RegistrationStatus.ValueString() == "EVALUATION" && plan.RegistrationType.ValueString() == "EVALUATION" {
		plan.RegistrationStatus = state.RegistrationStatus
		plan.Id = types.StringValue("")
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	if plan.RegistrationType.ValueString() == "REGISTER" && plan.Token.ValueString() == "" {
		resp.Diagnostics.AddError("Provider Error", "Token required for registration")
		return
	}

	if plan.Force.ValueBool() && state.RegistrationStatus.ValueString() != "UNREGISTERED" {
		res, err := r.deregisterSmartLicense(ctx, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
		res, err = r.client.Get(state.getPath(), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		state.fromBody(ctx, res.Get("items.0"))
	}

	if state.RegistrationStatus.ValueString() != "REGISTERED" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

		body := plan.toBody(ctx, SmartLicense{})
		res, err := r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
		plan.Id = types.StringValue(res.Get("id").String())

		tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

		// Read state after update
		res, err = r.client.Get(state.getPath(), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		state.fromBody(ctx, res.Get("items.0"))
	}

	plan.RegistrationStatus = state.RegistrationStatus
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SmartLicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SmartLicense

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
	res, err := r.deregisterSmartLicense(ctx, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

func (r *SmartLicenseResource) deregisterSmartLicense(ctx context.Context, reqMods ...func(*fmc.Req)) (gjson.Result, error) {
	plan := SmartLicense{
		RegistrationType: types.StringValue("DEREGISTER"),
	}
	body := plan.toBody(ctx, SmartLicense{})
	return r.client.Post(plan.getPath(), body, reqMods...)
}
