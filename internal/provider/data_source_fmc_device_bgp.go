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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceBGPDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceBGPDataSource{}
)

func NewDeviceBGPDataSource() datasource.DataSource {
	return &DeviceBGPDataSource{}
}

type DeviceBGPDataSource struct {
	client *fmc.Client
}

func (d *DeviceBGPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bgp"
}

func (d *DeviceBGPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device BGP.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the parent device (fmc_device.example.id).",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_address_family_type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_learned_route_map_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_default_information_orginate": schema.BoolAttribute{
				MarkdownDescription: "Generate default routes",
				Computed:            true,
			},
			"ipv4_auto_aummary": schema.BoolAttribute{
				MarkdownDescription: "Summarize subnet routes into network level routes",
				Computed:            true,
			},
			"ipv4_bgp_supress_inactive": schema.BoolAttribute{
				MarkdownDescription: "Suppresing advertise inactive routes",
				Computed:            true,
			},
			"ipv4_synchronization": schema.BoolAttribute{
				MarkdownDescription: "Synchronize between BGP and IGP systems",
				Computed:            true,
			},
			"ipv4_bgp_redistribute_internal": schema.BoolAttribute{
				MarkdownDescription: "Redistribute IBGP into IGP. (Use filtering to limit the number of prefixes that are redistributed)",
				Computed:            true,
			},
			"ipv4_external_distance": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_internal_distance": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_local_distance": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_forward_packets_over_multipath_ibgp": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_forward_packets_over_multipath_ebgp": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_neighbor_address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pv4_neighbor_romote_as": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_bfd": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_update_source_interface_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_address_family_ipv4": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_shutdown": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_description": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_filter_access_lists": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"ipv4_neighbor_filter_route_map_lists": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"route_map_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"ipv4_neighbor_filter_prefix_lists": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"route_map_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"ipv4_neighbor_filter_as_path_lists": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"update_direction": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"as_path_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"ipv4_neighbor_filter_max_prefix": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_filter_threshold_value": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_filter_restart_interval": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_routes_advertisement_interval": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_routes_remove_private_as": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_generate_default_route_map": schema.StringAttribute{
							MarkdownDescription: "Generate default routes - Route Map",
							Computed:            true,
						},
						"ipv4_neighbor_routes_advertise_map_use_exist": schema.BoolAttribute{
							MarkdownDescription: "Use Exist Map or Non-Exist Map",
							Computed:            true,
						},
						"ipv4_neighbor_routes_advertise_map": schema.StringAttribute{
							MarkdownDescription: "Specified route maps are advertised when the prefix exists in the Advertise Map and Exist Map.",
							Computed:            true,
						},
						"ipv4_neighbor_routes_advertise_exist_nonexist_map": schema.StringAttribute{
							MarkdownDescription: "Specified route maps are advertised when the prefix exists only in the Advertise Map.",
							Computed:            true,
						},
						"ipv4_neighbor_keepalive_interval": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_hold_time": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_min_hold_time": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_authentication_password": schema.StringAttribute{
							MarkdownDescription: "Setting password enables authentication.",
							Computed:            true,
						},
						"ipv4_neighbor_send_community_attribute": schema.BoolAttribute{
							MarkdownDescription: "Send Community attribute to this neighbor",
							Computed:            true,
						},
						"ipv4_neighbor_nexthop_self": schema.BoolAttribute{
							MarkdownDescription: "Use itself as next hop for this neighbor",
							Computed:            true,
						},
						"ipv4_neighbor_disable_connection_verification": schema.BoolAttribute{
							MarkdownDescription: "Disable Connection Verification",
							Computed:            true,
						},
						"ipv4_neighbor_tcp_mtu_path_discovery": schema.BoolAttribute{
							MarkdownDescription: "Use TCP path MTU discovery.",
							Computed:            true,
						},
						"ipv4_neighbor_max_hop_count": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_tcp_transport_mode": schema.BoolAttribute{
							MarkdownDescription: "True set it to active, False to passive.",
							Computed:            true,
						},
						"ipv4_neighbor_weight": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_version": schema.StringAttribute{
							MarkdownDescription: "0 - default, 4 - IPv4",
							Computed:            true,
						},
						"ipv4_neighbor_customized_local_as_number": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_neighbor_customized_no_prepend": schema.BoolAttribute{
							MarkdownDescription: "Do not prepend local AS number to routes received from neighbor",
							Computed:            true,
						},
						"ipv4_neighbor_customized_replace_as": schema.BoolAttribute{
							MarkdownDescription: "Replace real AS number with localAS number in routes received from neighbor",
							Computed:            true,
						},
						"ipv4_neighbor_customized_accept_both_as": schema.BoolAttribute{
							MarkdownDescription: "Replace real AS number with localAS number in routes received from neighbor",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"generate_as": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"filter": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"advertise_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"attribute_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"suppress_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_filterings": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"network_direction": schema.StringAttribute{
							MarkdownDescription: "Possible values - incomingroutefilter, outgoingroutefilter",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"prorocol_process": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_networks": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_redistributions": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"metric": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"process_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"source_protocol": schema.StringAttribute{
							MarkdownDescription: "Possible values - RedistributeConnected, RedistributeStatic, RedistributeOSPF, RedistributeEIGRP",
							Computed:            true,
						},
						"match_external1": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"match_external2": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"match_internal": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"match_nssa_external1": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"match_nssa_external2": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_route_injections": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inject_route_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"exist_route_map_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *DeviceBGPDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceBGPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceBGPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceBGP

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
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
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
