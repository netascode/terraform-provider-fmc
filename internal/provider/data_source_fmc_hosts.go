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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &HostsDataSource{}
	_ datasource.DataSourceWithConfigure = &HostsDataSource{}
)

func NewHostsDataSource() datasource.DataSource {
	return &HostsDataSource{}
}

type HostsDataSource struct {
	client *fmc.Client
}

func (d *HostsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hosts"
}

func (d *HostsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Hosts.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"items": schema.MapNestedAttribute{
				MarkdownDescription: "Map of hosts. The key of the map is the name of the individual Host. Renaming Hosts in bulk is not yet implemented.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID of the managed Host.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Optional user-created description.",
							Computed:            true,
						},
						"overridable": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether object values can be overridden.",
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: "IP of the host.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the object; this value is always 'Host'.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *HostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

func (d *HostsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Hosts

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

	// In one request, up to 1000 elements can be read
	// Here all elements from FMC are requested and locally only required are processed
	offset := 0
	limit := 1000
	fullOutput := `{"items":[]}`
	for page := 1; ; page++ {
		urlPath := config.getPath() + fmt.Sprintf("?expanded=true&limit=%d&offset=%d", limit, offset)
		res, err := d.client.Get(urlPath, reqMods...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET) from %s, got error: %s, %s", urlPath, err, res.String()))
			return
		}

		// Merge output of current request to overall list
		resItems := res.Get("items").String()
		fullOutput, _ = sjson.SetRaw(fullOutput, "items.-1", resItems[1:len(resItems)-1])

		// If there are no more pages to get, break the loop
		if !res.Get("paging.next.0").Exists() {
			break
		}

		// Increate offset to get next bulk of data
		offset += limit
	}

	config.fromBody(ctx, gjson.Parse(fullOutput))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}