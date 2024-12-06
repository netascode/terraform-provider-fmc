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

	// myInterfaceUUID := strings.ReplaceAll(plan.DeviceList.Elements()[0].String(), `"`, "")
	interfaceUUIDList, err := extractDeviceList(plan.DeviceList)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting device list", ""))
		return
	}

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

	// ############################################
	// As there are no GET api call for /api/fmc_config/v1/domain/{domainUUID}/deployment/deploymentrequests
	// We need to create abstraction state which reflects deployemnt
	// We can use /api/fmc_config/v1/domain/{domainUUID}/deployment/jobhistories API
	// and /api/fmc_config/v1/domain/{domainUUID}/job/taskstatuses/{objectId} (optionally, no jobID in taskstauses)
	// First after POST /api/fmc_config/v1/domain/{domainUUID}/deployment/deploymentrequests
	// we have to moinitor task status which available in response body of POST api call
	// then, when deployment task is finished we need to store Deployment job into tf state

	// Check deployemnt task status and waint until is finished
	// Read task id from deployment response
	resDeployment, err := resJson2Map(res)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting task url", ""))
		return
	}
	deploymentTaskId := resDeployment["metadata"].(map[string]interface{})["task"].(map[string]interface{})["id"].(interface{}).(string)

	// Get task status
	for {
		urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/" + deploymentTaskId
		res, err = r.client.Get(urlPath, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read job task status (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}
		resTaskstaus, err := resJson2Map(res)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("%s: Error getting task status", ""))
			return
		}
		if resTaskstaus["status"].(interface{}).(string) == "Deployed" {
			break
		}
	}

	// Read deployment history
	urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/jobhistories?expanded=true"
	res, err = r.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read deployment history (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	resMap, err := resJson2Map(res)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting job histories", ""))
		return
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

	// Iterate over items to find latest JobID based on devices UUID
JobIdLookup:
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		if itemMap["jobType"].(interface{}).(string) != "DEPLOYMENT" {
			continue
		}

		// *** It has to be fixed to add device list rather than single device
		deviceList, ok := itemMap["deviceList"].([]interface{})

		jobId, ok := itemMap["id"].(string)
		if !ok {
			continue
		}

		// Iterate over deviceData to check deviceUUID belongs to deployment
		for _, device := range deviceList {
			deviceMap, ok := device.(map[string]interface{})
			if !ok {
				continue
			}

			deviceUUID, ok := deviceMap["deviceUUID"].(interface{}).(string)
			if !ok {
				continue
			}

			// *** It has to be fixed to add device list rather than single device
			if contains(interfaceUUIDList, deviceUUID) {
				plan.Id = types.StringValue(jobId)
				matchingData = append(matchingData, deviceMap)
				break JobIdLookup
			}
		}
	}

	// Print the matching item
	if matchingItem != nil {
		tflog.Debug(ctx, fmt.Sprintf("Matching item: %+v\n", matchingItem))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("No matching item found. %+v\n", ""))
	}

	// ############################################

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Function to extract device UUIDs from plan
func extractDeviceList(deviceList types.List) ([]string, error) {
	var extractedList []string

	// Iterate through each element in the ListValue
	for _, elem := range deviceList.Elements() {
		// Convert the element to a StringValue and extract its value
		if stringValue, ok := elem.(types.String); ok {
			extractedList = append(extractedList, stringValue.ValueString())
		} else {
			return nil, fmt.Errorf("element is not a StringValue")
		}
	}

	return extractedList, nil
}

// Function to convert gjson.Result to a map
func resJson2Map(res gjson.Result) (map[string]interface{}, error) {
	// Convert gjson.Result to string
	jsonString := res.String()

	// Declare a map to hold the unmarshaled data
	var resMap map[string]interface{}

	// Unmarshal the JSON string into the map
	err := json.Unmarshal([]byte(jsonString), &resMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON data: %w", err)
	}

	// Return the map and nil error
	return resMap, nil
}

// Function to check if item belongs to list
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// Function to heck if all elements of list2 are in list1 (ignoring order)
func containsAllElements(list1, list2 []string) bool {
	lookup := make(map[string]bool)
	for _, item := range list1 {
		lookup[item] = true
	}
	for _, item := range list2 {
		if !lookup[item] {
			return false
		}
	}
	return true
}

func extractDeviceUUIDs(res string) ([]string, error) {
	// Parse the JSON into a map
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(res), &resMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Extract deviceList
	deviceList, ok := resMap["deviceList"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("deviceList is not an array or is missing")
	}

	// Extract deviceUUIDs
	var deviceUUIDs []string
	for _, item := range deviceList {
		device, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("deviceList item is not an object")
		}

		uuid, ok := device["deviceUUID"].(string)
		if !ok {
			return nil, fmt.Errorf("deviceUUID is not a string or is missing")
		}

		deviceUUIDs = append(deviceUUIDs, uuid)
	}

	return deviceUUIDs, nil
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

	// ############################################

	// Read list of deployable devices
	urlPathDeployable := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/deployabledevices?expanded=true"
	resDeployable, err := r.client.Get(urlPathDeployable, reqMods...)
	jsonStringDeployable := resDeployable.String()
	var resMapDeployable map[string]interface{}
	err = json.Unmarshal([]byte(jsonStringDeployable), &resMapDeployable)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error parsing JSON data (deployabledevices)", ""))
	}

	// Access "items"
	items, ok := resMapDeployable["items"].([]interface{})
	var deviceIdDeployable []string
	if !ok {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error: 'items' is not a valid array", ""))
	}
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		deviceIdDeployable = append(deviceIdDeployable, itemMap["device"].(map[string]interface{})["id"].(interface{}).(string))
	}

	// EO Read List of deployable devices

	// Target job ID to match
	jobId := state.Id.ValueString()

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

	// Read device list
	resDeviceIdList, err := extractDeviceUUIDs(res.String())
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting device list", ""))
		return
	}

	// Check if deployemnt hasn't been rolbacked and device list is not showing up as deployable device
	// Deployabale device list check
	if containsAllElements(deviceIdDeployable, resDeviceIdList) {
		// if contains(deviceIdDeployable, resDeviceId) {
		var resMap map[string]interface{}
		resMap, err = resJson2Map(res)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("%s: Error getting job histories", ""))
			return
		}

		resMap["deviceList"] = []interface{}{}

		resMapJSON, err := json.Marshal(resMap)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal JSON: %v", err))
		}
		res = gjson.Parse(string(resMapJSON))

		// After `terraform import` we switch to a full read.
		if imp {
			state.fromBody(ctx, res)
		} else {
			state.fromBodyPartial(ctx, res)
		}
	} else {
		// Update res with required key (tfsate) becuse we are using different API call to read device state
		var resMap map[string]interface{}
		resMap, err = resJson2Map(res)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("%s: Error parsing JSON data (jobhistories)", ""))
		}

		// Add hash values to res which are not exists in res but exists in state
		resMap["version"] = state.Version.ValueString()
		resMap["ignoreWarning"] = state.IgnoreWarning.ValueBool()
		resMapJSON, err := json.Marshal(resMap)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal JSON: %v", err))
		}
		res = gjson.Parse(string(resMapJSON))

		// Define the type of elements in the list
		elementType := types.StringType

		// Define the list values
		var values []attr.Value

		// Populate the values from resDeviceIdList
		for _, deviceId := range resDeviceIdList {
			values = append(values, types.StringValue(deviceId))
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
	}

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
	// res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
	// 	return
	// }

	// As there are no DELETE api call for /api/fmc_config/v1/domain/{domainUUID}/deployment/deploymentrequests
	// we can only roblback deployemnt to previous state
	// we can use /api/fmc_config/v1/domain/{domainUUID}/deployment/rollbackrequests to execute rolback
	// first we need to get deployemnt id and device list from state
	// next get deployemnt job id to rollback to, it can be retrieved from deployemnt job history
	// /api/fmc_config/v1/domain/{domainUUID}/deployment/jobhistories API
	// next initiate rollback api and wait until the job is finished

	// Get data from state
	stateDeviceList, err := extractDeviceList(state.DeviceList)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting device list from state", ""))
		return
	}
	stateDeploymentJobId := state.Id.ValueString()
	var rollbackToDeploymentJobId string

	// Read deployment history
	urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/jobhistories?expanded=true"
	res, err := r.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read deployment history (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	resMap, err := resJson2Map(res)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error getting job histories", ""))
		return
	}

	// Access "items"
	items, ok := resMap["items"].([]interface{})
	if !ok {
		tflog.Debug(ctx, fmt.Sprintf("%s: Error: 'items' is not a valid array", ""))
	}

	// Iterate over items to find JobID for rollback based on latest deployment job id and devices UUID
	// stateDeploymentJobId
JobIdLookup:
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		if itemMap["jobType"].(interface{}).(string) != "DEPLOYMENT" {
			continue
		}

		if (itemMap["id"].(interface{}).(string) != stateDeploymentJobId) && (rollbackToDeploymentJobId == "") {
			continue
		}

		deviceList, ok := itemMap["deviceList"].([]interface{})
		if !ok {
			continue
		}

		// Iterate over deviceData to check deviceUUID belongs to deployment
		for _, device := range deviceList {
			deviceMap, ok := device.(map[string]interface{})
			if !ok {
				continue
			}

			// *** It has to be fixed to add device list rather than single device
			deviceUUID, ok := deviceMap["deviceUUID"].(interface{}).(string)
			if !ok {
				continue
			}

			// *** It has to be fixed to add device list rather than single device
			if contains(stateDeviceList, deviceUUID) {
				// set rollbackToDeploymentJobId to be equal fo stateDeploymentJobId
				// and continuse to iterate to find previous deployemnt
				if rollbackToDeploymentJobId == "" {
					rollbackToDeploymentJobId = stateDeploymentJobId
					continue
				} else {
					rollbackToDeploymentJobId = itemMap["id"].(interface{}).(string)
					break JobIdLookup
				}
			}
		}
	}

	// Trigger rollback
	if rollbackToDeploymentJobId != "" && rollbackToDeploymentJobId != stateDeploymentJobId {

		urlPath = "/api/fmc_config/v1/domain/{DOMAIN_UUID}/deployment/rollbackrequests"
		body := `{ `
		body += `  "type": "RollbackRequest",`
		body += `  "rollbackDeviceList": [`
		body += `    {`
		body += `      "deploymentJobId": "` + rollbackToDeploymentJobId + `",`
		// *** It has to be fixed to add device list rather than single device
		body += `      "deviceList": [`
		body += `        "` + stateDeviceList[0] + `"`
		body += `      ]`
		body += `    }`
		body += `  ]`
		body += `}`

		res, err = r.client.Post(urlPath, body, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
			return
		}

		resDeployment, err := resJson2Map(res)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("%s: Error getting task url", ""))
			return
		}
		deploymentTaskId := resDeployment["metadata"].(map[string]interface{})["task"].(map[string]interface{})["id"].(interface{}).(string)

		// Get task status
		for {
			urlPath := "/api/fmc_config/v1/domain/{DOMAIN_UUID}/job/taskstatuses/" + deploymentTaskId
			res, err = r.client.Get(urlPath, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read job task status (POST/PUT), got error: %s, %s", err, res.String()))
				return
			}
			resTaskstaus, err := resJson2Map(res)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("%s: Error getting task status", ""))
				return
			}
			if resTaskstaus["status"].(interface{}).(string) == "Deployed" {
				break
			}
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))
		resp.State.RemoveResource(ctx)
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: Delete failed. Could not trigger rollback", state.Id.ValueString()))
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
