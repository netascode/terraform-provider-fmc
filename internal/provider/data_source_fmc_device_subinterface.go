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
	_ datasource.DataSource              = &DeviceSubinterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceSubinterfaceDataSource{}
)

func NewDeviceSubinterfaceDataSource() datasource.DataSource {
	return &DeviceSubinterfaceDataSource{}
}

type DeviceSubinterfaceDataSource struct {
	client *fmc.Client
}

func (d *DeviceSubinterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_subinterface"
}

func (d *DeviceSubinterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device Subinterface.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the parent device (fmc_device.example.id).",
				Required:            true,
			},
			"parent_id": schema.StringAttribute{
				MarkdownDescription: "Link to the parent interface (fmc_device_physical_interface.example.link). The attribute is actually used solely for dependency propagation and its value is ignored: you would still need to set the `parent_name`.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable the subinterface.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"parent_name": schema.StringAttribute{
				MarkdownDescription: "Name of the parent interface. As the fmc_device_physical_interface.example.name does not propagate dependency adequately on Terraform, you additionally have to use `parent_id` attribute.",
				Computed:            true,
			},
			"index": schema.Int64Attribute{
				MarkdownDescription: "The numerical id of this subinterface, unique on the parent interface.",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN identifier, unique per the parent interface. Must be non-empty in order to set `logical_name`, `security_zone_id`, `mtu`.",
				Computed:            true,
			},
			"logical_name": schema.StringAttribute{
				MarkdownDescription: "Customizable logical name of the subinterface, unique on the device. Should not contain whitespace or slash characters. Can only be set when vlan_id is set. Must be non-empty in order to set `security_zone_id` or `mtu`.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"management_only": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether this subinterface limits traffic to management traffic; when true, through-the-box traffic is disallowed. Value true conflicts with mode INLINE, PASSIVE, TAP, ERSPAN, or with security_zone_id.",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Maximum transmission unit. Can only be used when logical_name is set on the parent interface.",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority 0-65535. Can only be set for routed subinterfaces.",
				Computed:            true,
			},
			"security_zone_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the assigned security zone (fmc_security_zone.example.id). Can only be used when logical_name is set.",
				Computed:            true,
			},
			"ipv4_static_address": schema.StringAttribute{
				MarkdownDescription: "Static IPv4 address. Conflicts with mode INLINE, PASSIVE, TAP, ERSPAN.",
				Computed:            true,
			},
			"ipv4_static_netmask": schema.StringAttribute{
				MarkdownDescription: "Netmask (width) for ipv4_static_address.",
				Computed:            true,
			},
			"ipv4_dhcp_obtain_route": schema.BoolAttribute{
				MarkdownDescription: "Any non-null value here indicates to enable DHCPv4. Value `false` indicates to enable DHCPv4 without obtaining from there the default IPv4 route but anyway requires also ipv4_dhcp_route_metric to be set to exactly 1. Value `true` indicates to enable DHCPv4 and obtain the route and also requires ipv4_dhcp_route_metric to be non-null. The ipv4_dhcp_obtain_route must be null when using ipv4_static_netmask.",
				Computed:            true,
			},
			"ipv4_dhcp_route_metric": schema.Int64Attribute{
				MarkdownDescription: "The metric for ipv4_dhcp_obtain_route. Any non-null value enables DHCP as a side effect. Must be null when using ipv4_static_netmask.",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6.",
				Computed:            true,
			},
			"ipv6_enforce_eui": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
				Computed:            true,
			},
			"ipv6_enable_auto_config": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6 autoconfiguration.",
				Computed:            true,
			},
			"ipv6_enable_dhcp_address": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable DHCPv6 for address config.",
				Computed:            true,
			},
			"ipv6_enable_dhcp_nonaddress": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable DHCPv6 for non-address config.",
				Computed:            true,
			},
			"ipv6_enable_ra": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable IPv6 router advertisement (RA).",
				Computed:            true,
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 address without a slash and prefix.",
							Computed:            true,
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: "Prefix width for the IPv6 address.",
							Computed:            true,
						},
						"enforce_eui": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether to enforce IPv6 Extended Unique Identifier (EUI64 from RFC2373).",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DeviceSubinterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *DeviceSubinterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceSubinterface

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
