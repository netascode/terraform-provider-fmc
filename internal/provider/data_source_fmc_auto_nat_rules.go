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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AutoNatRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &AutoNatRulesDataSource{}
)

func NewAutoNatRulesDataSource() datasource.DataSource {
	return &AutoNatRulesDataSource{}
}

type AutoNatRulesDataSource struct {
	client *fmc.Client
}

func (d *AutoNatRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auto_nat_rules"
}

func (d *AutoNatRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Auto Nat Rules.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"nat_policy_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the associated nat policy",
				Required:            true,
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"interface_ipv6": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"fall_through": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"dns": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"route_lookup": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"no_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"net_to_net": schema.BoolAttribute{
				MarkdownDescription: "The type of the rule",
				Computed:            true,
			},
			"original_network_id": schema.StringAttribute{
				MarkdownDescription: "ID of original network object",
				Computed:            true,
			},
			"translated_network_id": schema.StringAttribute{
				MarkdownDescription: "ID of translated network object",
				Computed:            true,
			},
			"source_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of source security zone",
				Computed:            true,
			},
			"destination_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of destination security zone",
				Computed:            true,
			},
			"interface_pat": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"include_reverse": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"round_robin": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"extended_pat": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"flat_port_range": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"block_allocation": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"pat_pool_id": schema.StringAttribute{
				MarkdownDescription: "ID of pat host object",
				Computed:            true,
			},
			"pat_pool_name": schema.StringAttribute{
				MarkdownDescription: "Name of pat host object",
				Computed:            true,
			},
		},
	}
}

func (d *AutoNatRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *AutoNatRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AutoNatRules

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath()+"/"+url.QueryEscape(config.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
