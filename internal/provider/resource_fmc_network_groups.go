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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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
	_ resource.Resource                = &NetworkGroupsResource{}
	_ resource.ResourceWithImportState = &NetworkGroupsResource{}
)

func NewNetworkGroupsResource() resource.Resource {
	return &NetworkGroupsResource{}
}

type NetworkGroupsResource struct {
	client *fmc.Client
}

func (r *NetworkGroupsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_groups"
}

func (r *NetworkGroupsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This plural resource manages a bulk of Network Groups. The `terraform import` of the plural resource is not yet implemented. The FMC API supports quick bulk creation for this resource, but the deletion/modification is done one-by-one. Updating and deleting `fmc_network_groups` can thus take much more time than creating it (even >500 times more time, i.e. >50000%, depending on the change size).").String,

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
			"items": schema.MapNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Map of network groups. The key of the map is the name of the individual Network Group.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("UUID of the managed Network Group.").String,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Optional user-created description.").String,
							Optional:            true,
						},
						"overridable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether object values can be overridden.").String,
							Optional:            true,
						},
						"group_names": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of names (not UUIDs) of child Network Groups. The names must be defined in the same instance of fmc_network_groups resource. This is an auxiliary way to add a child Network Group: the suggested way is to add it inside `objects` by its UUID. Attribute `group_names` is used in managed resource only and always empty in a data-source. Modification of contents of `group_names` is not yet implemented.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"objects": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of objects (fmc_network, fmc_host, ...).").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID of the object (such as fmc_network.this.id, etc.).").String,
										Optional:            true,
									},
								},
							},
						},
						"literals": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set of literal values.").String,
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
					},
				},
			},
		},
	}
}

func (r *NetworkGroupsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (r *NetworkGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkGroups

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
	body := plan.toBody(ctx, NetworkGroups{})
	// It's a pseudo-object, no Post needed.
	plan.Id = types.StringValue("00000000-0000-0000-0000-000000000000")

	var res gjson.Result
	// Create subobjects "Items"
	// Subobjects can use bulk Post.
	{
		subbody := gjson.Parse(body).Get("items").String()
		// manual fixup
		for i := range gjson.Parse(body).Get("items").Array() {
			subbody, _ = sjson.Delete(subbody, fmt.Sprintf("%d.group_names", i))
		}

		// TODO: first create child groups, then their parents
		subres, err := r.client.Post(plan.getPath()+"?bulk=true", subbody, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, subres.String()))
			return
		}

		res = set(res, "items", subres.Get("items"))
	}

	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func set(orig gjson.Result, path string, content gjson.Result) gjson.Result {
	s, err := sjson.SetRaw(orig.String(), path, content.String())
	if err != nil {
		panic(err)
	}

	return gjson.Parse(s)
}

func (r *NetworkGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkGroups

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

	var res gjson.Result
	// Read subobjects from "items"
	{
		// TODO: >1000
		params := "?expanded=true&limit=1000"
		subres, err := r.client.Get(state.getPath()+params, reqMods...)
		if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, subres.String()))
			return
		}

		res = set(res, "items", subres)
	}

	// Pseudo-resource, no Get.
	if len(state.Items) == 0 {
		resp.State.RemoveResource(ctx)
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

	resp.Diagnostics.AddError("cannot Read", "not yet implemeneted")
	return

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *NetworkGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkGroups

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

	subbody := gjson.Parse(body).Get("items")

	// TODO handle renames:
	// keyByUUID := map[string]string{}
	// for k := range plan.Items {
	// 	keyByUUID[state.Items[k].Id.ValueString()] = k
	// }
	// for k := range plan.Items {
	// 	delete(keyByUUID, plan.Items[k].Id.ValueString())
	// }

	// Subresources to Delete.
	for k := range state.Items {
		if _, found := plan.Items[k]; found {
			continue
		}

		deleting := state.Items[k].Id.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", deleting))

		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(deleting), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}

		delete(state.Items, k)
		diags = resp.State.Set(ctx, &state)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", deleting))
	}

	// Subresources to Update.
	for _, subelem := range subbody.Array() {
		k := subelem.Get("name").String()
		if k == "" {
			resp.Diagnostics.AddError("Internal Error", "key not marshaled")
			return
		}

		if plan.Items[k].Id.IsUnknown() {
			continue
		}

		unchanged, diags := objectUnchangedAt(ctx, req, path.Root("items").AtMapKey(k))
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		if unchanged {
			continue
		}

		// adjust
		// TODO: why do we throw it away? Fail on any change?
		tmp, _ := sjson.Delete(subelem.String(), "group_names")
		subelem = gjson.Parse(tmp)

		res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(subelem.Get("id").String()), subelem.String(), reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Subresources to Create: topological sort of a DAG
	creating := `[]`
	for {
		for _, subelem := range subbody.Array() {
			k := subelem.Get("name").String()
			if k == "" {
				resp.Diagnostics.AddError("Internal Error", "key not marshaled")
				return
			}

			if !plan.Items[k].Id.IsUnknown() {
				continue
			}

			children := subelem.Get("group_names")
			known := true
			for _, child := range children.Array() {
				existing := plan.Items[child.String()].Id
				if existing.IsUnknown() {
					tflog.Debug(ctx, fmt.Sprintf("FIXME postpone %s until after %s", k, child))
					known = false
					break
				}

				obj, _ := sjson.Set("{}", "id", existing.ValueString())
				obj, _ = sjson.Set(obj, "type", "AnyNonEmptyString")
				subelem = set(subelem, "objects.-1", gjson.Parse(obj))
			}
			if known {
				tmp, _ := sjson.Delete(subelem.String(), "group_names")
				subelem = gjson.Parse(tmp)

				creating, _ = sjson.SetRaw(creating, "-1", subelem.String())
			}
		}

		if creating == `[]` {
			break
		}

		tflog.Debug(ctx, fmt.Sprintf("FIXME bulk %s", creating))

		resp.Diagnostics.AddError("cannot create new item during modification", "not yet implemeneted (cause: any uuid marked Unknown in the plan)")
		return

		// Post
		// fromBodyUnknowns
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func objectUnchangedAt(ctx context.Context, req resource.UpdateRequest, where path.Path) (bool, diag.Diagnostics) {
	var plan, state attr.Value

	diags := req.Plan.GetAttribute(ctx, where, &plan)
	if diags.HasError() {
		return false, diags
	}

	diags = req.State.GetAttribute(ctx, where, &state)
	if diags.HasError() {
		return false, nil
	}

	return state.Equal(plan), diags
}

func (r *NetworkGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkGroups

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete converting to Update", state.Id.ValueString()))

	// Convert Delete call into Update call
	uresp := &resource.UpdateResponse{
		State: resp.State,
	}
	ureq := &resource.UpdateRequest{
		Plan: tfsdk.Plan{
			Schema: req.State.Schema,
		},
		State:        req.State,
		ProviderMeta: req.ProviderMeta,
		Private:      req.Private,
	}
	diags = ureq.Plan.Set(ctx, &NetworkGroups{})
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	r.Update(ctx, *ureq, uresp)

	resp.Private = uresp.Private
	resp.State = uresp.State
	resp.Diagnostics = uresp.Diagnostics

	if resp.Diagnostics.HasError() {
		tflog.Debug(ctx, fmt.Sprintf("%s: Delete failed", state.Id.ValueString()))
		return
	}

	resp.State.RemoveResource(ctx)
	tflog.Debug(ctx, fmt.Sprintf("%s: Delete successful", state.Id.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *NetworkGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
