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
	"encoding/json"
	"fmt"

	// "net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"

	// "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

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
				MarkdownDescription: helpers.NewAttributeDescription("Epoch unix time stamp (13 digits).").String,
				Optional:            true,
			},
			"jobid": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("JobId of deployment.").String,
				Optional:            true,
			},
			"force_deploy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force deployment (even if there are no configuration changes).").String,
				Optional:            true,
			},
			"ignore_warning": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore warnings during deployment.").String,
				Optional:            true,
			},
			"device_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of device ids to be deployed.").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"deployment_note": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("User note related to deployment.").String,
				Optional:            true,
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

// End of section. //template:end model

func (r *DeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Deployment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	myInterfaceUUID := strings.ReplaceAll(plan.DeviceList.Elements()[0].String(), `"`, "")

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Deployment{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}

	// BoCODE
	// As there are no GET api call for /api/fmc_config/v1/domain/{domainUUID}/deployment/deploymentrequests
	// We need to create abstraction state which reflects deployemnt
	// We can use /api/fmc_config/v1/domain/{domainUUID}/deployment/jobhistories API
	// and /api/fmc_config/v1/domain/{domainUUID}/job/taskstatuses/{objectId} (optionally, no jobID in taskstauses)
	// First after POST /api/fmc_config/v1/domain/{domainUUID}/deployment/deploymentrequests
	// we have to moinitor task status which available in response body of POST api call
	// then, when deployment task is finished we need to store Deployment job into tf state

	urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/jobhistories?expanded=true"
	res, err = r.client.Get(urlPath, reqMods...)
	jsonString := res.String()
	var resMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &resMap)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error parsing JSON data (jobhistories)", ""))
	}

	// Variable to store the matching item
	var matchingItem map[string]interface{}
	// Slice to collect matching data
	var matchingData []interface{}

	// Access "items"
	items, ok := resMap["items"].([]interface{})
	if !ok {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error: 'items' is not a valid array", ""))
	}

	// Iterate over items to find the JobID based on devices UUID
JonbIdLookup:
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// // Check if the ID matches
		// if id, ok := itemMap["id"].(string); ok && id == my_job_id {
		// 	matchingItem = itemMap
		// 	break // Exit the loop as we found the match
		// }

		if itemMap["jobType"].(interface{}).(string) != "DEPLOYMENT" {
			continue
		}

		deviceList, ok := itemMap["deviceList"].([]interface{})

		jobId, ok := itemMap["id"].(string)
		if !ok {
			continue
		}
		println(jobId)

		// Iterate over deviceData to check deviceUUID
		for _, device := range deviceList {
			deviceMap, ok := device.(map[string]interface{})
			if !ok {
				continue
			}

			// deviceDetails, ok := deviceMap["data"].(map[string]interface{})
			// if !ok {
			// 	continue
			// }
			deviceUUID, ok := deviceMap["deviceUUID"].(interface{}).(string)
			if !ok {
				continue
			}

			if deviceUUID == myInterfaceUUID {
				plan.Id = types.StringValue(jobId)
				matchingData = append(matchingData, deviceMap)
				break JonbIdLookup
			}
		}
	}

	// Print the matching item
	if matchingItem != nil {
		tflog.Debug(ctx, fmt.Sprintf("Matching item: %+v\n", matchingItem))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("No matching item found. %+v\n", ""))
	}

	// EoCODE

	// plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	// BoCODE

	// Target job ID to match
	jobId := state.Id.ValueString()
	// state.Jobid = types.StringValue(jobId)

	urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/jobhistories/" + jobId
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

	// Update res with required key (tfsate)
	jsonString := res.String()
	var resMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &resMap)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error parsing JSON data (jobhistories)", ""))
	}
	resMap["version"] = "1732274866000"
	// resMap["ignore_warning"] = "true"
	resMapJSON, err := json.Marshal(resMap)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal JSON: %v", err))
	}
	res = gjson.Parse(string(resMapJSON))

	// Read device list
	resDeviceId := res.Get("deviceList.0.deviceUUID").String()
	// resIgnoreWarning := res.Get("ignoreWarning").Bool()
	// resVersion := res.Get("version").String()

	// Define the type of elements in the list
	elementType := types.StringType
	// Define the list values
	values := []attr.Value{
		types.StringValue(resDeviceId),
	}
	// Create the ListValue
	deviceList, diags := types.ListValue(elementType, values)
	if diags.HasError() {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error creating deviceList (jobhistories)", ""))
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	state.DeviceList = deviceList
	// state.IgnoreWarning = types.BoolValue(resIgnoreWarning)
	// state.Version = types.StringValue(resVersion)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *DeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Deployment

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

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

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
