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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &PolicyAssignmentResource{}
	_ resource.ResourceWithImportState = &PolicyAssignmentResource{}
)

func NewPolicyAssignmentResource() resource.Resource {
	return &PolicyAssignmentResource{}
}

type PolicyAssignmentResource struct {
	client *fmc.Client
}

func (r *PolicyAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_assignment"
}

func (r *PolicyAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Assignment.").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the object; this is always 'PolicyAssignment'").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policy to be assigned.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the policy to be assigned.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"policy_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the policy to be assigned.").AddStringEnumDescription("FTDNatPolicy", "HealthPolicy", "AccessPolicy").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("FTDNatPolicy", "HealthPolicy", "AccessPolicy"),
				},
			},
			"after_destroy_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Policy to be assigned after this policy assignment is destroyed. Applicable for Health and Access Control Policies only.").String,
				Optional:            true,
			},
			"targets": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of devices to which policy should be attached").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the device to which policy should be attached").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the device to which policy should be attached").AddStringEnumDescription("Device", "DeviceHAPair", "DeviceGroup").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Device", "DeviceHAPair", "DeviceGroup"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *PolicyAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *PolicyAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyAssignment

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

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	// Check if policy assignment already exists
	// For Update requests it may not exist in case all devices are re-assigned to other policies (i.a. other TF resource)
	res, err := r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.PolicyId.ValueString()), reqMods...)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 500")) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment does not exist", plan.Id.ValueString()))

		// Create object (POST)
		body := plan.toBody(ctx, PolicyAssignment{})
		body, _ = sjson.Delete(body, "dummy_after_destroy_policy_id")

		res, err = r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else if err != nil {
		// Failed to retrieve policy assignments
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	} else {
		// Policy assignment already exists - need to update it
		tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", plan.Id.ValueString()))

		// Update object (PUT)
		updateBody := res.String()
		updateBody, _ = sjson.Delete(updateBody, "links")

		// Update target list
		for _, target := range plan.Targets {
			if !res.Get(fmt.Sprintf("targets.#(id==%s)", target.Id.ValueString())).Exists() {
				updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
					"id":   target.Id.ValueString(),
					"type": target.Type.ValueString(),
				})
			}
		}

		// Update object
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.PolicyId.ValueString()), updateBody, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update policy assignment (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *PolicyAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyAssignment

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

func (r *PolicyAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update assigns policy to device once on the list
	var plan, state PolicyAssignment

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

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 500")) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment does not exist", plan.Id.ValueString()))

		// Create object (POST)
		body := plan.toBody(ctx, PolicyAssignment{})
		body, _ = sjson.Delete(body, "dummy_after_destroy_policy_id")

		res, err = r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else if err != nil {
		// Failed to retrieve policy assignments
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	} else {
		// Policy assignment already exists - need to update it (PUT)
		tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", plan.Id.ValueString()))

		updateBody := res.String()
		updateBody, _ = sjson.Delete(updateBody, "links")

		if state.PolicyType.ValueString() == "AccessPolicy" || state.PolicyType.ValueString() == "HealthPolicy" {
			// Update target list
			for _, target := range plan.Targets {
				if !res.Get(fmt.Sprintf("targets.#(id==%s)", target.Id.ValueString())).Exists() {
					updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
						"id":   target.Id.ValueString(),
						"type": target.Type.ValueString(),
					})
				}
			}
		} else {
			// Get list of IDs from devices from state
			var stateTargets = make([]string, len(state.Targets))
			for _, target := range state.Targets {
				stateTargets = append(stateTargets, target.Id.ValueString())
			}

			// Update target list - remove all devices from state
			existingTargets := gjson.Get(updateBody, "targets").Array()
			updateBody, _ = sjson.Set(updateBody, "targets", []interface{}{})
			for _, target := range existingTargets {
				if !helpers.Contains(stateTargets, target.Get("id").String()) {
					updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
						"id":   target.Get("id").String(),
						"type": target.Get("type").String(),
					})
				}
			}

			// temporary variable for JSON lookups
			tmpUpdateBody := gjson.Parse(updateBody)

			// Add devices from plan
			for _, target := range plan.Targets {
				tflog.Debug(ctx, fmt.Sprintf("target: %s", target.Id.ValueString()))
				if !tmpUpdateBody.Get(fmt.Sprintf("targets.#(id==%s)", target.Id.String())).Exists() {
					updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
						"id":   target.Id.ValueString(),
						"type": target.Type.ValueString(),
					})
				}
			}
		}

		// Update object
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.PolicyId.ValueString()), updateBody, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update policy assignment (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyAssignment

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

	// make sure only one policy assignment is updated at a time
	policyMu.Lock()
	defer policyMu.Unlock()

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	if state.PolicyType.ValueString() == "HealthPolicy" {
		state.PolicyId = state.AfterDestroyPolicyId

		// Create object (POST)
		body := state.toBody(ctx, PolicyAssignment{})
		body, _ = sjson.Delete(body, "dummy_after_destroy_policy_id")

		res, err := r.client.Post(state.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else if state.PolicyType.ValueString() == "AccessPolicy" {
		res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.AfterDestroyPolicyId.ValueString()), reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			tflog.Debug(ctx, fmt.Sprintf("%s: After destroy policy assignment does not exist", state.Id.ValueString()))

			state.PolicyId = state.AfterDestroyPolicyId

			// Create object (POST)
			body := state.toBody(ctx, PolicyAssignment{})
			body, _ = sjson.Delete(body, "dummy_after_destroy_policy_id")

			res, err = r.client.Post(state.getPath(), body, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
				return
			}
		} else if err != nil {
			// Failed to retrieve policy assignments
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		} else {
			// Policy assignment already exists - need to update it
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment already exists", state.Id.ValueString()))

			// Update object (PUT)
			updateBody := res.String()
			updateBody, _ = sjson.Delete(updateBody, "links")

			// Update target list
			for _, target := range state.Targets {
				if !res.Get(fmt.Sprintf("targets.#(id==%s)", target.Id.ValueString())).Exists() {
					updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
						"id":   target.Id.ValueString(),
						"type": target.Type.ValueString(),
					})
				}
			}

			// Update object
			res, err = r.client.Put(state.getPath()+"/"+url.QueryEscape(state.AfterDestroyPolicyId.ValueString()), updateBody, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update policy assignment (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.PolicyId.ValueString()), reqMods...)
		if err == nil {
			// Policy assignment already exists - need to update it (remove configured devices, but keep non-managed ones)
			tflog.Debug(ctx, fmt.Sprintf("%s: Policy assignment exists", state.Id.ValueString()))

			// Get list of IDs from devices from state
			var stateTargets = make([]string, len(state.Targets))
			for _, target := range state.Targets {
				stateTargets = append(stateTargets, target.Id.ValueString())
			}

			// Update object (PUT)
			updateBody := res.String()
			updateBody, _ = sjson.Delete(updateBody, "links")

			// Update target list - keep non-managed devices
			existingTargets := gjson.Get(updateBody, "targets").Array()
			updateBody, _ = sjson.Set(updateBody, "targets", []interface{}{})
			for _, target := range existingTargets {
				if !helpers.Contains(stateTargets, target.Get("id").String()) {
					updateBody, _ = sjson.Set(updateBody, "targets.-1", map[string]string{
						"id":   target.Get("id").String(),
						"type": target.Get("type").String(),
					})
				}
			}

			// There is no DELETE API endpoint for policy assignment, so we need to update the policy assignment (even with empty targets)
			res, err = r.client.Put(state.getPath()+"/"+url.QueryEscape(state.PolicyId.ValueString()), updateBody, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update policy assignment (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *PolicyAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
