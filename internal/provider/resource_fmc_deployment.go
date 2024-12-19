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

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &DeploymentResource{}
)

func NewDeploymentResource() resource.Resource {
	return &DeploymentResource{}
}

type DeploymentResource struct {
	client *fmc.Client
}

func (r *DeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deployment"
}

func (r *DeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Deployment.").String,

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
			"version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version to which the deployment should be done. By default set to current unix timestamp.").String,
				Optional:            true,
				Computed:            true,
			},
			"force_deploy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force deployment (even if there are no configuration changes).").String,
				Optional:            true,
			},
			"device_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of device ids to be deployed. If forceDeploy is not set to true, only devices that are deployable will be deployed.").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"deployment_note": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("User note related to deployment.").String,
				Optional:            true,
			},
			"deploy": schema.BoolAttribute{
				MarkdownDescription: "This attribute is only used internally.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *DeploymentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

func (r *DeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Deployment

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

	// If version not provided, set it to the current unix timestamp
	if plan.Version.IsUnknown() {
		plan.Version = types.StringValue(strconv.FormatInt(time.Now().UnixMilli(), 10))
	}

	origDeviceList := plan.DeviceList

	if !plan.ForceDeploy.ValueBool() {
		// Get list of deployable devices
		resDeployable, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deployabledevices?expanded=true", reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to obtain list of deployable devices object (GET), got error: %s, %s", err, resDeployable.String()))
			return
		}
		deployableDevices := resDeployable.Get("items.#.device.id")

		var deployableDeviceIds []string
		deployableDevices.ForEach(func(_, value gjson.Result) bool {
			deployableDeviceIds = append(deployableDeviceIds, value.String())
			return true
		})

		var deviceIdsList []string
		plan.DeviceList.ElementsAs(ctx, &deviceIdsList, false)

		devicesToDeploy := filterForDeployableDevices(deviceIdsList, deployableDeviceIds)

		plan.DeviceList = helpers.GetStringListFromStringSlice(devicesToDeploy)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	var deviceIdsSlice []string
	plan.DeviceList.ElementsAs(ctx, &deviceIdsSlice, false)

	if len(deviceIdsSlice) != 0 {
		// Create object
		body := plan.toBody(ctx, Deployment{})
		body, _ = sjson.Delete(body, "deploy")
		res, err := r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	plan.DeviceList = origDeviceList
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.State.SetAttribute(ctx, path.Root("deploy"), false)
}

func (r *DeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan Deployment

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

	// If version not provided, set it to the current unix timestamp
	if plan.Version.IsUnknown() {
		plan.Version = types.StringValue(strconv.FormatInt(time.Now().UnixMilli(), 10))
	}

	origDeviceList := plan.DeviceList

	if !plan.ForceDeploy.ValueBool() {
		// Get list of deployable devices
		resDeployable, err := r.client.Get("/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deployabledevices?expanded=true", reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to obtain list of deployable devices object (GET), got error: %s, %s", err, resDeployable.String()))
			return
		}
		deployableDevices := resDeployable.Get("items.#.device.id")

		var deployableDeviceIds []string
		deployableDevices.ForEach(func(_, value gjson.Result) bool {
			deployableDeviceIds = append(deployableDeviceIds, value.String())
			return true
		})

		var deviceIdsList []string
		plan.DeviceList.ElementsAs(ctx, &deviceIdsList, false)

		devicesToDeploy := filterForDeployableDevices(deviceIdsList, deployableDeviceIds)

		plan.DeviceList = helpers.GetStringListFromStringSlice(devicesToDeploy)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	var deviceIdsSlice []string
	plan.DeviceList.ElementsAs(ctx, &deviceIdsSlice, false)

	if len(deviceIdsSlice) != 0 {
		// Create object
		body := plan.toBody(ctx, Deployment{})
		body, _ = sjson.Delete(body, "deploy")
		res, err := r.client.Post(plan.getPath(), body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	plan.DeviceList = origDeviceList
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Deployment

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Checks what devices from those that were requested for deployment are actually deployable
// Function returns list of requested device ids that are deployable
func filterForDeployableDevices(listOfDevices, deployableDevices []string) []string {
	devicesMap := make(map[string]bool)
	for _, str := range listOfDevices {
		devicesMap[str] = true
	}

	var result []string
	for _, str := range deployableDevices {
		if devicesMap[str] {
			result = append(result, str)
			// Remove the element to avoid duplicates in the result
			delete(devicesMap, str)
		}
	}

	return result
}
