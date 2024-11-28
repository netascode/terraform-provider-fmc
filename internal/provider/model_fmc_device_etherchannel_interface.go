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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceEtherChannelInterface struct {
	Id                       types.String                                    `tfsdk:"id"`
	Domain                   types.String                                    `tfsdk:"domain"`
	DeviceId                 types.String                                    `tfsdk:"device_id"`
	Enabled                  types.Bool                                      `tfsdk:"enabled"`
	Mode                     types.String                                    `tfsdk:"mode"`
	EtherChannelId           types.String                                    `tfsdk:"ether_channel_id"`
	LogicalName              types.String                                    `tfsdk:"logical_name"`
	Description              types.String                                    `tfsdk:"description"`
	ManagementOnly           types.Bool                                      `tfsdk:"management_only"`
	Mtu                      types.Int64                                     `tfsdk:"mtu"`
	Priority                 types.Int64                                     `tfsdk:"priority"`
	SecurityZoneId           types.String                                    `tfsdk:"security_zone_id"`
	Ipv4StaticAddress        types.String                                    `tfsdk:"ipv4_static_address"`
	Ipv4StaticNetmask        types.String                                    `tfsdk:"ipv4_static_netmask"`
	Ipv4DhcpObtainRoute      types.Bool                                      `tfsdk:"ipv4_dhcp_obtain_route"`
	Ipv4DhcpRouteMetric      types.Int64                                     `tfsdk:"ipv4_dhcp_route_metric"`
	Ipv6Enable               types.Bool                                      `tfsdk:"ipv6_enable"`
	Ipv6EnforceEui           types.Bool                                      `tfsdk:"ipv6_enforce_eui"`
	Ipv6EnableAutoConfig     types.Bool                                      `tfsdk:"ipv6_enable_auto_config"`
	Ipv6EnableDhcpAddress    types.Bool                                      `tfsdk:"ipv6_enable_dhcp_address"`
	Ipv6EnableDhcpNonaddress types.Bool                                      `tfsdk:"ipv6_enable_dhcp_nonaddress"`
	Ipv6EnableRa             types.Bool                                      `tfsdk:"ipv6_enable_ra"`
	Ipv6Addresses            []DeviceEtherChannelInterfaceIpv6Addresses      `tfsdk:"ipv6_addresses"`
	NveOnly                  types.Bool                                      `tfsdk:"nve_only"`
	SelectedInterfaces       []DeviceEtherChannelInterfaceSelectedInterfaces `tfsdk:"selected_interfaces"`
}

type DeviceEtherChannelInterfaceIpv6Addresses struct {
	Address    types.String `tfsdk:"address"`
	Prefix     types.String `tfsdk:"prefix"`
	EnforceEui types.Bool   `tfsdk:"enforce_eui"`
}

type DeviceEtherChannelInterfaceSelectedInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceEtherChannelInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/etherchannelinterfaces", url.QueryEscape(data.DeviceId.ValueString()), url.QueryEscape(data.EtherChannelId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceEtherChannelInterface) toBody(ctx context.Context, state DeviceEtherChannelInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	body, _ = sjson.Set(body, "type", "EtherChannelInterface")
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ManagementOnly.IsNull() {
		body, _ = sjson.Set(body, "managementOnly", data.ManagementOnly.ValueBool())
	}
	if !data.Mtu.IsNull() {
		body, _ = sjson.Set(body, "MTU", data.Mtu.ValueInt64())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.SecurityZoneId.IsNull() {
		body, _ = sjson.Set(body, "securityZone.id", data.SecurityZoneId.ValueString())
	}
	body, _ = sjson.Set(body, "securityZone.type", "SecurityZone")
	if !data.Ipv4StaticAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.address", data.Ipv4StaticAddress.ValueString())
	}
	if !data.Ipv4StaticNetmask.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.netmask", data.Ipv4StaticNetmask.ValueString())
	}
	if !data.Ipv4DhcpObtainRoute.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.enableDefaultRouteDHCP", data.Ipv4DhcpObtainRoute.ValueBool())
	}
	if !data.Ipv4DhcpRouteMetric.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.dhcpRouteMetric", data.Ipv4DhcpRouteMetric.ValueInt64())
	}
	if !data.Ipv6Enable.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableIPV6", data.Ipv6Enable.ValueBool())
	}
	if !data.Ipv6EnforceEui.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enforceEUI64", data.Ipv6EnforceEui.ValueBool())
	}
	if !data.Ipv6EnableAutoConfig.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableAutoConfig", data.Ipv6EnableAutoConfig.ValueBool())
	}
	if !data.Ipv6EnableDhcpAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPAddrConfig", data.Ipv6EnableDhcpAddress.ValueBool())
	}
	if !data.Ipv6EnableDhcpNonaddress.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPNonAddrConfig", data.Ipv6EnableDhcpNonaddress.ValueBool())
	}
	if !data.Ipv6EnableRa.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableRA", data.Ipv6EnableRa.ValueBool())
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6.addresses", []interface{}{})
		for _, item := range data.Ipv6Addresses {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefix", item.Prefix.ValueString())
			}
			if !item.EnforceEui.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enforceEUI64", item.EnforceEui.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "ipv6.addresses.-1", itemBody)
		}
	}
	if !data.NveOnly.IsNull() {
		body, _ = sjson.Set(body, "nveOnly", data.NveOnly.ValueBool())
	}
	if len(data.SelectedInterfaces) > 0 {
		body, _ = sjson.Set(body, "selectedInterfaces", []interface{}{})
		for _, item := range data.SelectedInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "selectedInterfaces.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceEtherChannelInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("managementOnly"); value.Exists() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("MTU"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("securityZone.id"); value.Exists() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpAddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpNonaddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	if value := res.Get("ipv6.addresses"); value.Exists() {
		data.Ipv6Addresses = make([]DeviceEtherChannelInterfaceIpv6Addresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEtherChannelInterfaceIpv6Addresses{}
			if value := res.Get("address"); value.Exists() {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("prefix"); value.Exists() {
				data.Prefix = types.StringValue(value.String())
			} else {
				data.Prefix = types.StringNull()
			}
			if value := res.Get("enforceEUI64"); value.Exists() {
				data.EnforceEui = types.BoolValue(value.Bool())
			} else {
				data.EnforceEui = types.BoolNull()
			}
			(*parent).Ipv6Addresses = append((*parent).Ipv6Addresses, data)
			return true
		})
	}
	if value := res.Get("nveOnly"); value.Exists() {
		data.NveOnly = types.BoolValue(value.Bool())
	} else {
		data.NveOnly = types.BoolValue(false)
	}
	if value := res.Get("selectedInterfaces"); value.Exists() {
		data.SelectedInterfaces = make([]DeviceEtherChannelInterfaceSelectedInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEtherChannelInterfaceSelectedInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringValue("PhysicalInterface")
			}
			(*parent).SelectedInterfaces = append((*parent).SelectedInterfaces, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceEtherChannelInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("ifname"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("managementOnly"); value.Exists() && !data.ManagementOnly.IsNull() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("MTU"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("securityZone.id"); value.Exists() && !data.SecurityZoneId.IsNull() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() && !data.Ipv4StaticAddress.IsNull() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() && !data.Ipv4StaticNetmask.IsNull() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() && !data.Ipv4DhcpObtainRoute.IsNull() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() && !data.Ipv4DhcpRouteMetric.IsNull() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() && !data.Ipv6Enable.IsNull() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() && !data.Ipv6EnforceEui.IsNull() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() && !data.Ipv6EnableAutoConfig.IsNull() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpAddress.IsNull() {
		data.Ipv6EnableDhcpAddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpNonaddress.IsNull() {
		data.Ipv6EnableDhcpNonaddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() && !data.Ipv6EnableRa.IsNull() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	for i := 0; i < len(data.Ipv6Addresses); i++ {
		keys := [...]string{"address", "prefix"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString(), data.Ipv6Addresses[i].Prefix.ValueString()}

		parent := &data
		data := (*parent).Ipv6Addresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ipv6.addresses").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv6Addresses[%d] = %+v",
				i,
				(*parent).Ipv6Addresses[i],
			))
			(*parent).Ipv6Addresses = slices.Delete((*parent).Ipv6Addresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && !data.Prefix.IsNull() {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("enforceEUI64"); value.Exists() && !data.EnforceEui.IsNull() {
			data.EnforceEui = types.BoolValue(value.Bool())
		} else {
			data.EnforceEui = types.BoolNull()
		}
		(*parent).Ipv6Addresses[i] = data
	}
	if value := res.Get("nveOnly"); value.Exists() && !data.NveOnly.IsNull() {
		data.NveOnly = types.BoolValue(value.Bool())
	} else if data.NveOnly.ValueBool() != false {
		data.NveOnly = types.BoolNull()
	}
	for i := 0; i < len(data.SelectedInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SelectedInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).SelectedInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("selectedInterfaces").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SelectedInterfaces[%d] = %+v",
				i,
				(*parent).SelectedInterfaces[i],
			))
			(*parent).SelectedInterfaces = slices.Delete((*parent).SelectedInterfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else if data.Type.ValueString() != "PhysicalInterface" {
			data.Type = types.StringNull()
		}
		(*parent).SelectedInterfaces[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceEtherChannelInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk