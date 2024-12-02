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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBGP struct {
	Id                              types.String            `tfsdk:"id"`
	Domain                          types.String            `tfsdk:"domain"`
	DeviceId                        types.String            `tfsdk:"device_id"`
	Name                            types.String            `tfsdk:"name"`
	Type                            types.String            `tfsdk:"type"`
	AsNumber                        types.String            `tfsdk:"as_number"`
	AddressFamilyType               types.String            `tfsdk:"address_family_type"`
	LearnedRouteMapId               types.String            `tfsdk:"learned_route_map_id"`
	DefaultInformationOrginate      types.Bool              `tfsdk:"default_information_orginate"`
	AutoSummary                     types.Bool              `tfsdk:"auto_summary"`
	BgpSupressInactive              types.Bool              `tfsdk:"bgp_supress_inactive"`
	Synchronization                 types.Bool              `tfsdk:"synchronization"`
	BgpRedistributeInternal         types.Bool              `tfsdk:"bgp_redistribute_internal"`
	ExternalDistance                types.Int64             `tfsdk:"external_distance"`
	InternalDistance                types.Int64             `tfsdk:"internal_distance"`
	LocalDistance                   types.Int64             `tfsdk:"local_distance"`
	ForwardPacketsOverMultipathIbgp types.Int64             `tfsdk:"forward_packets_over_multipath_ibgp"`
	ForwardPacketsOverMultipathEbgp types.Int64             `tfsdk:"forward_packets_over_multipath_ebgp"`
	Neighbors                       []DeviceBGPNeighbors    `tfsdk:"neighbors"`
	MaximumPaths                    []DeviceBGPMaximumPaths `tfsdk:"maximum_paths"`
}

type DeviceBGPNeighbors struct {
	Ipv4Address             types.String `tfsdk:"ipv4_address"`
	RomoteAs                types.String `tfsdk:"romote_as"`
	Bfd                     types.String `tfsdk:"bfd"`
	UpdateSourceInterfaceId types.String `tfsdk:"update_source_interface_id"`
	AddressFamilyIpv4       types.Bool   `tfsdk:"address_family_ipv4"`
	Shutdown                types.Bool   `tfsdk:"shutdown"`
	Description             types.String `tfsdk:"description"`
}

type DeviceBGPMaximumPaths struct {
	Value types.Int64 `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBGP) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/bgp", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBGP) toBody(ctx context.Context, state DeviceBGP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LearnedRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.aftableMap.id", data.LearnedRouteMapId.ValueString())
	}
	if !data.DefaultInformationOrginate.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.defaultInformationOrginate", data.DefaultInformationOrginate.ValueBool())
	}
	if !data.AutoSummary.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.autoSummary", data.AutoSummary.ValueBool())
	}
	if !data.BgpSupressInactive.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpSupressInactive", data.BgpSupressInactive.ValueBool())
	}
	if !data.Synchronization.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.synchronization", data.Synchronization.ValueBool())
	}
	if !data.BgpRedistributeInternal.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpRedistributeInternal", data.BgpRedistributeInternal.ValueBool())
	}
	if !data.ExternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.externalDistance", data.ExternalDistance.ValueInt64())
	}
	if !data.InternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.internalDistance", data.InternalDistance.ValueInt64())
	}
	if !data.LocalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.localDistance", data.LocalDistance.ValueInt64())
	}
	if !data.ForwardPacketsOverMultipathIbgp.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ibgp", data.ForwardPacketsOverMultipathIbgp.ValueInt64())
	}
	if !data.ForwardPacketsOverMultipathEbgp.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ebgp", data.ForwardPacketsOverMultipathEbgp.ValueInt64())
	}
	if len(data.Neighbors) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.neighbors", []interface{}{})
		for _, item := range data.Neighbors {
			itemBody := ""
			if !item.Ipv4Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address", item.Ipv4Address.ValueString())
			}
			if !item.RomoteAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteAs", item.RomoteAs.ValueString())
			}
			if !item.Bfd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.fallOverBFD", item.Bfd.ValueString())
			}
			if !item.UpdateSourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.updateSource.id", item.UpdateSourceInterfaceId.ValueString())
			}
			if !item.AddressFamilyIpv4.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.enableAddress", item.AddressFamilyIpv4.ValueBool())
			}
			if !item.Shutdown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.shutdown", item.Shutdown.ValueBool())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.description", item.Description.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.neighbors.-1", itemBody)
		}
	}
	if len(data.MaximumPaths) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.maximumPaths", []interface{}{})
		for _, item := range data.MaximumPaths {
			itemBody := ""
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.maximumPaths.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBGP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
		data.AddressFamilyType = types.StringValue(value.String())
	} else {
		data.AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() {
		data.LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() {
		data.DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else {
		data.DefaultInformationOrginate = types.BoolValue(false)
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() {
		data.AutoSummary = types.BoolValue(value.Bool())
	} else {
		data.AutoSummary = types.BoolValue(false)
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() {
		data.BgpSupressInactive = types.BoolValue(value.Bool())
	} else {
		data.BgpSupressInactive = types.BoolValue(true)
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() {
		data.Synchronization = types.BoolValue(value.Bool())
	} else {
		data.Synchronization = types.BoolValue(false)
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() {
		data.BgpRedistributeInternal = types.BoolValue(value.Bool())
	} else {
		data.BgpRedistributeInternal = types.BoolValue(false)
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() {
		data.ExternalDistance = types.Int64Value(value.Int())
	} else {
		data.ExternalDistance = types.Int64Value(20)
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() {
		data.InternalDistance = types.Int64Value(value.Int())
	} else {
		data.InternalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() {
		data.LocalDistance = types.Int64Value(value.Int())
	} else {
		data.LocalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() {
		data.ForwardPacketsOverMultipathIbgp = types.Int64Value(value.Int())
	} else {
		data.ForwardPacketsOverMultipathIbgp = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() {
		data.ForwardPacketsOverMultipathEbgp = types.Int64Value(value.Int())
	} else {
		data.ForwardPacketsOverMultipathEbgp = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.neighbors"); value.Exists() {
		data.Neighbors = make([]DeviceBGPNeighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPNeighbors{}
			if value := res.Get("ipv4Address"); value.Exists() {
				data.Ipv4Address = types.StringValue(value.String())
			} else {
				data.Ipv4Address = types.StringNull()
			}
			if value := res.Get("remoteAs"); value.Exists() {
				data.RomoteAs = types.StringValue(value.String())
			} else {
				data.RomoteAs = types.StringNull()
			}
			if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() {
				data.Bfd = types.StringValue(value.String())
			} else {
				data.Bfd = types.StringValue("NONE")
			}
			if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() {
				data.UpdateSourceInterfaceId = types.StringValue(value.String())
			} else {
				data.UpdateSourceInterfaceId = types.StringNull()
			}
			if value := res.Get("neighborGeneral.enableAddress"); value.Exists() {
				data.AddressFamilyIpv4 = types.BoolValue(value.Bool())
			} else {
				data.AddressFamilyIpv4 = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.shutdown"); value.Exists() {
				data.Shutdown = types.BoolValue(value.Bool())
			} else {
				data.Shutdown = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			(*parent).Neighbors = append((*parent).Neighbors, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.maximumPaths"); value.Exists() {
		data.MaximumPaths = make([]DeviceBGPMaximumPaths, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPMaximumPaths{}
			if value := res.Get("value"); value.Exists() {
				data.Value = types.Int64Value(value.Int())
			} else {
				data.Value = types.Int64Value(1)
			}
			(*parent).MaximumPaths = append((*parent).MaximumPaths, data)
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
func (data *DeviceBGP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() && !data.AddressFamilyType.IsNull() {
		data.AddressFamilyType = types.StringValue(value.String())
	} else {
		data.AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() && !data.LearnedRouteMapId.IsNull() {
		data.LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() && !data.DefaultInformationOrginate.IsNull() {
		data.DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else if data.DefaultInformationOrginate.ValueBool() != false {
		data.DefaultInformationOrginate = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() && !data.AutoSummary.IsNull() {
		data.AutoSummary = types.BoolValue(value.Bool())
	} else if data.AutoSummary.ValueBool() != false {
		data.AutoSummary = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() && !data.BgpSupressInactive.IsNull() {
		data.BgpSupressInactive = types.BoolValue(value.Bool())
	} else if data.BgpSupressInactive.ValueBool() != true {
		data.BgpSupressInactive = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() && !data.Synchronization.IsNull() {
		data.Synchronization = types.BoolValue(value.Bool())
	} else if data.Synchronization.ValueBool() != false {
		data.Synchronization = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() && !data.BgpRedistributeInternal.IsNull() {
		data.BgpRedistributeInternal = types.BoolValue(value.Bool())
	} else if data.BgpRedistributeInternal.ValueBool() != false {
		data.BgpRedistributeInternal = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() && !data.ExternalDistance.IsNull() {
		data.ExternalDistance = types.Int64Value(value.Int())
	} else if data.ExternalDistance.ValueInt64() != 20 {
		data.ExternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() && !data.InternalDistance.IsNull() {
		data.InternalDistance = types.Int64Value(value.Int())
	} else if data.InternalDistance.ValueInt64() != 200 {
		data.InternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() && !data.LocalDistance.IsNull() {
		data.LocalDistance = types.Int64Value(value.Int())
	} else if data.LocalDistance.ValueInt64() != 200 {
		data.LocalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() && !data.ForwardPacketsOverMultipathIbgp.IsNull() {
		data.ForwardPacketsOverMultipathIbgp = types.Int64Value(value.Int())
	} else if data.ForwardPacketsOverMultipathIbgp.ValueInt64() != 1 {
		data.ForwardPacketsOverMultipathIbgp = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() && !data.ForwardPacketsOverMultipathEbgp.IsNull() {
		data.ForwardPacketsOverMultipathEbgp = types.Int64Value(value.Int())
	} else if data.ForwardPacketsOverMultipathEbgp.ValueInt64() != 1 {
		data.ForwardPacketsOverMultipathEbgp = types.Int64Null()
	}
	for i := 0; i < len(data.Neighbors); i++ {
		keys := [...]string{"neighborGeneral.updateSource.id"}
		keyValues := [...]string{data.Neighbors[i].UpdateSourceInterfaceId.ValueString()}

		parent := &data
		data := (*parent).Neighbors[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.neighbors").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Neighbors[%d] = %+v",
				i,
				(*parent).Neighbors[i],
			))
			(*parent).Neighbors = slices.Delete((*parent).Neighbors, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4Address"); value.Exists() && !data.Ipv4Address.IsNull() {
			data.Ipv4Address = types.StringValue(value.String())
		} else {
			data.Ipv4Address = types.StringNull()
		}
		if value := res.Get("remoteAs"); value.Exists() && !data.RomoteAs.IsNull() {
			data.RomoteAs = types.StringValue(value.String())
		} else {
			data.RomoteAs = types.StringNull()
		}
		if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() && !data.Bfd.IsNull() {
			data.Bfd = types.StringValue(value.String())
		} else if data.Bfd.ValueString() != "NONE" {
			data.Bfd = types.StringNull()
		}
		if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() && !data.UpdateSourceInterfaceId.IsNull() {
			data.UpdateSourceInterfaceId = types.StringValue(value.String())
		} else {
			data.UpdateSourceInterfaceId = types.StringNull()
		}
		if value := res.Get("neighborGeneral.enableAddress"); value.Exists() && !data.AddressFamilyIpv4.IsNull() {
			data.AddressFamilyIpv4 = types.BoolValue(value.Bool())
		} else if data.AddressFamilyIpv4.ValueBool() != false {
			data.AddressFamilyIpv4 = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.shutdown"); value.Exists() && !data.Shutdown.IsNull() {
			data.Shutdown = types.BoolValue(value.Bool())
		} else if data.Shutdown.ValueBool() != false {
			data.Shutdown = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		(*parent).Neighbors[i] = data
	}
	for i := 0; i < len(data.MaximumPaths); i++ {
		keys := [...]string{"value"}
		keyValues := [...]string{strconv.FormatInt(data.MaximumPaths[i].Value.ValueInt64(), 10)}

		parent := &data
		data := (*parent).MaximumPaths[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.maximumPaths").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing MaximumPaths[%d] = %+v",
				i,
				(*parent).MaximumPaths[i],
			))
			(*parent).MaximumPaths = slices.Delete((*parent).MaximumPaths, i, i+1)
			i--

			continue
		}
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.Int64Value(value.Int())
		} else if data.Value.ValueInt64() != 1 {
			data.Value = types.Int64Null()
		}
		(*parent).MaximumPaths[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBGP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.AsNumber.IsUnknown() {
		if value := res.Get("asNumber"); value.Exists() {
			data.AsNumber = types.StringValue(value.String())
		} else {
			data.AsNumber = types.StringNull()
		}
	}
	if data.AddressFamilyType.IsUnknown() {
		if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
			data.AddressFamilyType = types.StringValue(value.String())
		} else {
			data.AddressFamilyType = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
