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
	"time"

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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AccessControlPolicyResource{}
var _ resource.ResourceWithImportState = &AccessControlPolicyResource{}

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
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the action to take when the conditions defined by the rule are met.").AddStringEnumDescription("BLOCK", "TRUST", "PERMIT", "NETWORK_DISCOVERY", "INHERIT_FROM_PARENT").String,
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
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of rules. The first matching rule overshadows all the matching rules which follow it.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier of the rule.").String,
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
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the access rule is in effect (true) or not (false). Default is true.").AddDefaultValueDescription("true").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"source_network_literals": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
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
						// "source_network_literals": schema.ListAttribute{
						// 	MarkdownDescription: helpers.NewAttributeDescription("List of CIDRs.").String,
						// 	ElementType:         types.StringType,
						// 	Optional:            true,
						// },
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

func (r *AccessControlPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccessControlPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	resCats, err := r.createCats(ctx, plan, bodyCats, 0, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())

		res, err = r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}

		return
	}

	resRules, err := r.createRules(ctx, plan, bodyRules, 0, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())

		res, err = r.client.Delete(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Also, cannot DELETE a hanging policy object, got error: %s, %s", err, res.String()))
		}

		return
	}

	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	replace, _ := sjson.SetRaw(res.String(), "dummy_categories", resCats.Get("items").String())
	replace, _ = sjson.SetRaw(replace, "dummy_rules", resRules.Get("items").String())
	res = gjson.Parse(replace)

	plan.updateFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AccessControlPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	resCats, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/categories?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	resRules, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/accessrules?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	replaceCats := resCats.Get("items").String()
	if replaceCats == "" {
		replaceCats = "[]"
	}
	replaceRules := resRules.Get("items").String()
	if replaceRules == "" {
		replaceRules = "[]"
	}
	replace, _ := sjson.SetRaw(res.String(), "dummy_categories", replaceCats)
	replace, _ = sjson.SetRaw(replace, "dummy_rules", replaceRules)
	res = gjson.Parse(replace)

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *AccessControlPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AccessControlPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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

	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	keptCats, keptRules := r.countKept(ctx, state, plan)

	_, err = r.truncateRules(ctx, state, keptRules, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	_, err = r.truncateCats(ctx, state, keptCats, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	resCats, err := r.createCats(ctx, plan, bodyCats, keptCats, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())

		// We've did some DELETE/POST calls, determine the current tfstate.
		readResp := resource.ReadResponse{}
		r.Read(ctx, resource.ReadRequest{
			State:        req.State,
			Private:      req.Private,
			ProviderMeta: req.ProviderMeta,
		}, &readResp)
		resp.State = readResp.State
		resp.Diagnostics.Append(readResp.Diagnostics...)
		return
	}

	resRules, err := r.createRules(ctx, plan, bodyRules.Array(), keptRules, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	replace, _ := sjson.SetRaw(res.String(), "dummy_categories", resCats.Get("items").String())
	replace, _ = sjson.SetRaw(replace, "dummy_rules", resRules.Get("items").String())
	res = gjson.Parse(replace)
	plan.updateFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
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
	return 0, 0
}

func (r *AccessControlPolicyResource) truncateRules(ctx context.Context, state AccessControlPolicy, kept int, reqMods ...func(*fmc.Req)) (*gjson.Result, error) {
	var err error
	var b strings.Builder
	var bulks []string

	for i := kept; i < len(state.Rules); i++ {
		if b.Len() != 0 {
			b.WriteString(",")
		}
		b.WriteString(state.Rules[i].Id.ValueString())
		if b.Len() > 4000 {
			bulks = append(bulks, b.String())
			b.Reset()
		}
	}
	if b.Len() > 0 {
		bulks = append(bulks, b.String())
	}

	for _, bulk := range bulks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+
			"/accessrules?bulk=true&filter=ids:"+url.QueryEscape(bulk), reqMods...)
		if err != nil {
			return nil, fmt.Errorf("Failed to bulk-delete rules, got error: %v, %s", err, res.String())
		}
	}

	// Apparently, the bulk DELETE has a race. Gather logs:
	troubleshoot, err := r.client.Get(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+"/accessrules?expanded=true&offset=0&limit=1000", reqMods...)
	if err != nil {
		return nil, fmt.Errorf("Failed to bulk-get rules, got error: %s, %s", err, troubleshoot.String())
	}

	// Apparently, the bulk DELETE has a race. Stabilize:
	time.Sleep(time.Second)

	return nil, nil
}

func (r *AccessControlPolicyResource) truncateCats(ctx context.Context, state AccessControlPolicy, kept int, reqMods ...func(*fmc.Req)) (*gjson.Result, error) {
	for i := kept; i < len(state.Categories); i++ {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString())+
			"/categories/"+url.QueryEscape(state.Categories[i].Id.ValueString()), reqMods...)
		if err != nil {
			return nil, fmt.Errorf("Failed to delete a category, got error: %v, %s", err, res)
		}
	}
	return nil, nil
}

func (r *AccessControlPolicyResource) createCats(ctx context.Context, plan AccessControlPolicy, body []gjson.Result, kept int, reqMods ...func(*fmc.Req)) (*gjson.Result, error) {
	gathered := `{"items":[]}`

	for i := kept; i < len(plan.Categories); i++ {
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+
			"/categories", body[i].String(), reqMods...)
		if err != nil {
			return nil, fmt.Errorf("Failed to create a category (POST), got error: %v, %s", err, res)
		}

		gathered, _ = sjson.SetRaw(gathered, "items.-1", res.String())
	}

	ret := gjson.Parse(gathered)
	return &ret, nil
}

func (r *AccessControlPolicyResource) createRules(ctx context.Context, plan AccessControlPolicy, body []gjson.Result, kept int, reqMods ...func(*fmc.Req)) (*gjson.Result, error) {
	gathered := `{"items":[]}`

	for i := kept; i < len(body); i++ {
		bulk := `{"dummy_rules":[]}`
		cat := body[i].Get("metadata.category").String()
		for ; i < len(body); i++ {
			if cat != body[i].Get("metadata.category").String() {
				i--
				break
			}
			rule := body[i].String()
			rule, _ = sjson.Delete(rule, "id")
			rule, _ = sjson.Delete(rule, "metadata.category")

			bulk, _ = sjson.SetRaw(bulk, "dummy_rules.-1", rule)
		}

		param := "?bulk=true"
		if cat != "" {
			param += "&category=" + url.QueryEscape(cat)
		}
		res, err := r.client.Post(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+"/accessrules"+param,
			gjson.Parse(bulk).Get("dummy_rules").String(),
			reqMods...)
		if err != nil {
			return nil, fmt.Errorf("Failed to configure object (POST), got error: %s, %s", err, res.String())
		}

		for _, rule := range res.Get("items").Array() {
			// POST usually reports a fake item with a non-existing UUID. Shake fist at the sky.
			if !rule.Get("name").Exists() {
				continue
			}

			gathered, _ = sjson.SetRaw(gathered, "items.-1", rule.String())
		}
	}

	ret := gjson.Parse(gathered)
	return &ret, nil
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
}

// End of section. //template:end import
