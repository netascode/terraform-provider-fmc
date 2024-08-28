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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkGroupsDataSource{}
)

func NewNetworkGroupsDataSource() datasource.DataSource {
	return &NetworkGroupsDataSource{}
}

type NetworkGroupsDataSource struct {
	client *fmc.Client
}

func (d *NetworkGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_groups"
}

func (d *NetworkGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This plural data source can query a bulk of Network Groups.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of network groups. The key of the map is the name of the individual Network Group. Renaming Network Groups is not yet implemented.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID of the managed Network Group.",
							Required:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Optional user-created description.",
							Computed:            true,
						},
						"overridable": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether object values can be overridden.",
							Computed:            true,
						},
						"network_groups": schema.SetAttribute{
							MarkdownDescription: "Set of names (not UUIDs) of child Network Groups. The names must be defined in the same instance of fmc_network_groups resource. This is an auxiliary way to add a child Network Group: the suggested way is to instead add it inside `objects` by its UUID. Renaming a group contained in `network_groups` is not yet implemented, while it works in `objects`.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"objects": schema.SetNestedAttribute{
							MarkdownDescription: "Set of objects (fmc_network, fmc_host, ...).",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "UUID of the object (such as fmc_network.example.id, fmc_host.example.id, fmc_network_groups.another.items[\"example\"].id, etc.).",
										Computed:            true,
									},
								},
							},
						},
						"literals": schema.SetNestedAttribute{
							MarkdownDescription: "Set of literal values.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
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

func (d *NetworkGroupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

func (d *NetworkGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkGroups

	// Read config
	diags := req.Config.Get(ctx, &config)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, diags := readNetworkGroupsSubresources(ctx, d.client, config, reqMods...)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
