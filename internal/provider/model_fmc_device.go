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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type Device struct {
	Id                           types.String `tfsdk:"id"`
	Domain                       types.String `tfsdk:"domain"`
	Name                         types.String `tfsdk:"name"`
	Type                         types.String `tfsdk:"type"`
	HostName                     types.String `tfsdk:"host_name"`
	NatId                        types.String `tfsdk:"nat_id"`
	LicenseCapabilities          types.Set    `tfsdk:"license_capabilities"`
	RegistrationKey              types.String `tfsdk:"registration_key"`
	DeviceGroupId                types.String `tfsdk:"device_group_id"`
	ProhibitPacketTransfer       types.Bool   `tfsdk:"prohibit_packet_transfer"`
	PerformanceTier              types.String `tfsdk:"performance_tier"`
	SnortEngine                  types.String `tfsdk:"snort_engine"`
	ObjectGroupSearch            types.Bool   `tfsdk:"object_group_search"`
	AccessPolicyId               types.String `tfsdk:"access_policy_id"`
	NatPolicyId                  types.String `tfsdk:"nat_policy_id"`
	HealthPolicyId               types.String `tfsdk:"health_policy_id"`
	InfoVersion                  types.String `tfsdk:"info_version"`
	InfoHealthStatus             types.String `tfsdk:"info_health_status"`
	InfoHealthMessage            types.String `tfsdk:"info_health_message"`
	InfoIsConnected              types.Bool   `tfsdk:"info_is_connected"`
	InfoDeploymentStatus         types.String `tfsdk:"info_deployment_status"`
	InfoFtdMode                  types.String `tfsdk:"info_ftd_mode"`
	InfoDeviceSerialNumber       types.String `tfsdk:"info_device_serial_number"`
	InfoSnortVersion             types.String `tfsdk:"info_snort_version"`
	InfoVdbVersion               types.String `tfsdk:"info_vdb_version"`
	InfoLspVersion               types.String `tfsdk:"info_lsp_version"`
	InfoDeployedAccessPolicyName types.String `tfsdk:"info_deployed_access_policy_name"`
	InfoDeployedHealthPolicyName types.String `tfsdk:"info_deployed_health_policy_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Device) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Device) toBody(ctx context.Context, state Device) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, "hostName", data.HostName.ValueString())
	}
	if !data.NatId.IsNull() {
		body, _ = sjson.Set(body, "natID", data.NatId.ValueString())
	}
	if !data.LicenseCapabilities.IsNull() {
		var values []string
		data.LicenseCapabilities.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "license_caps", values)
	}
	if !data.RegistrationKey.IsNull() {
		body, _ = sjson.Set(body, "regKey", data.RegistrationKey.ValueString())
	}
	if !data.DeviceGroupId.IsNull() {
		body, _ = sjson.Set(body, "deviceGroup.id", data.DeviceGroupId.ValueString())
	}
	if !data.ProhibitPacketTransfer.IsNull() {
		body, _ = sjson.Set(body, "prohibitPacketTransfer", data.ProhibitPacketTransfer.ValueBool())
	}
	if !data.PerformanceTier.IsNull() {
		body, _ = sjson.Set(body, "performanceTier", data.PerformanceTier.ValueString())
	}
	if !data.SnortEngine.IsNull() {
		body, _ = sjson.Set(body, "snortEngine", data.SnortEngine.ValueString())
	}
	if !data.ObjectGroupSearch.IsNull() {
		body, _ = sjson.Set(body, "advanced.enableOGS", data.ObjectGroupSearch.ValueBool())
	}
	if !data.AccessPolicyId.IsNull() {
		body, _ = sjson.Set(body, "accessPolicy.id", data.AccessPolicyId.ValueString())
	}
	if !data.NatPolicyId.IsNull() {
		body, _ = sjson.Set(body, "dummy_nat_policy_id", data.NatPolicyId.ValueString())
	}
	if !data.HealthPolicyId.IsNull() {
		body, _ = sjson.Set(body, "dummy_health_policy_id", data.HealthPolicyId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringValue("Device")
	}
	if value := res.Get("hostName"); value.Exists() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() {
		data.LicenseCapabilities = helpers.GetStringSet(value.Array())
	} else {
		data.LicenseCapabilities = types.SetNull(types.StringType)
	}
	if value := res.Get("deviceGroup.id"); value.Exists() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
	if value := res.Get("prohibitPacketTransfer"); value.Exists() {
		data.ProhibitPacketTransfer = types.BoolValue(value.Bool())
	} else {
		data.ProhibitPacketTransfer = types.BoolNull()
	}
	if value := res.Get("performanceTier"); value.Exists() {
		data.PerformanceTier = types.StringValue(value.String())
	} else {
		data.PerformanceTier = types.StringNull()
	}
	if value := res.Get("snortEngine"); value.Exists() {
		data.SnortEngine = types.StringValue(value.String())
	} else {
		data.SnortEngine = types.StringNull()
	}
	if value := res.Get("advanced.enableOGS"); value.Exists() {
		data.ObjectGroupSearch = types.BoolValue(value.Bool())
	} else {
		data.ObjectGroupSearch = types.BoolValue(true)
	}
	if value := res.Get("accessPolicy.id"); value.Exists() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}
	if value := res.Get("dummy_nat_policy_id"); value.Exists() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}
	if value := res.Get("dummy_health_policy_id"); value.Exists() {
		data.HealthPolicyId = types.StringValue(value.String())
	} else {
		data.HealthPolicyId = types.StringNull()
	}
	if value := res.Get("sw_version"); value.Exists() {
		data.InfoVersion = types.StringValue(value.String())
	} else {
		data.InfoVersion = types.StringNull()
	}
	if value := res.Get("healthStatus"); value.Exists() {
		data.InfoHealthStatus = types.StringValue(value.String())
	} else {
		data.InfoHealthStatus = types.StringNull()
	}
	if value := res.Get("healthMessage"); value.Exists() {
		data.InfoHealthMessage = types.StringValue(value.String())
	} else {
		data.InfoHealthMessage = types.StringNull()
	}
	if value := res.Get("isConnected"); value.Exists() {
		data.InfoIsConnected = types.BoolValue(value.Bool())
	} else {
		data.InfoIsConnected = types.BoolNull()
	}
	if value := res.Get("deploymentStatus"); value.Exists() {
		data.InfoDeploymentStatus = types.StringValue(value.String())
	} else {
		data.InfoDeploymentStatus = types.StringNull()
	}
	if value := res.Get("ftdMode"); value.Exists() {
		data.InfoFtdMode = types.StringValue(value.String())
	} else {
		data.InfoFtdMode = types.StringNull()
	}
	if value := res.Get("deviceSerialNumber"); value.Exists() {
		data.InfoDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.InfoDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("snortVersion"); value.Exists() {
		data.InfoSnortVersion = types.StringValue(value.String())
	} else {
		data.InfoSnortVersion = types.StringNull()
	}
	if value := res.Get("vdbVersion"); value.Exists() {
		data.InfoVdbVersion = types.StringValue(value.String())
	} else {
		data.InfoVdbVersion = types.StringNull()
	}
	if value := res.Get("lspVersion"); value.Exists() {
		data.InfoLspVersion = types.StringValue(value.String())
	} else {
		data.InfoLspVersion = types.StringNull()
	}
	if value := res.Get("accessPolicy.name"); value.Exists() {
		data.InfoDeployedAccessPolicyName = types.StringValue(value.String())
	} else {
		data.InfoDeployedAccessPolicyName = types.StringNull()
	}
	if value := res.Get("healthPolicy.name"); value.Exists() {
		data.InfoDeployedHealthPolicyName = types.StringValue(value.String())
	} else {
		data.InfoDeployedHealthPolicyName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *Device) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else if data.Type.ValueString() != "Device" {
		data.Type = types.StringNull()
	}
	if value := res.Get("hostName"); value.Exists() && !data.HostName.IsNull() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("license_caps"); value.Exists() && !data.LicenseCapabilities.IsNull() {
		data.LicenseCapabilities = helpers.GetStringSet(value.Array())
	} else {
		data.LicenseCapabilities = types.SetNull(types.StringType)
	}
	if value := res.Get("deviceGroup.id"); value.Exists() && !data.DeviceGroupId.IsNull() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
	if value := res.Get("prohibitPacketTransfer"); value.Exists() && !data.ProhibitPacketTransfer.IsNull() {
		data.ProhibitPacketTransfer = types.BoolValue(value.Bool())
	} else {
		data.ProhibitPacketTransfer = types.BoolNull()
	}
	if value := res.Get("performanceTier"); value.Exists() && !data.PerformanceTier.IsNull() {
		data.PerformanceTier = types.StringValue(value.String())
	} else {
		data.PerformanceTier = types.StringNull()
	}
	if value := res.Get("snortEngine"); value.Exists() && !data.SnortEngine.IsNull() {
		data.SnortEngine = types.StringValue(value.String())
	} else {
		data.SnortEngine = types.StringNull()
	}
	if value := res.Get("advanced.enableOGS"); value.Exists() && !data.ObjectGroupSearch.IsNull() {
		data.ObjectGroupSearch = types.BoolValue(value.Bool())
	} else if data.ObjectGroupSearch.ValueBool() != true {
		data.ObjectGroupSearch = types.BoolNull()
	}
	if value := res.Get("accessPolicy.id"); value.Exists() && !data.AccessPolicyId.IsNull() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}
	if value := res.Get("dummy_nat_policy_id"); value.Exists() && !data.NatPolicyId.IsNull() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}
	if value := res.Get("dummy_health_policy_id"); value.Exists() && !data.HealthPolicyId.IsNull() {
		data.HealthPolicyId = types.StringValue(value.String())
	} else {
		data.HealthPolicyId = types.StringNull()
	}
	if value := res.Get("sw_version"); value.Exists() && !data.InfoVersion.IsNull() {
		data.InfoVersion = types.StringValue(value.String())
	} else {
		data.InfoVersion = types.StringNull()
	}
	if value := res.Get("healthStatus"); value.Exists() && !data.InfoHealthStatus.IsNull() {
		data.InfoHealthStatus = types.StringValue(value.String())
	} else {
		data.InfoHealthStatus = types.StringNull()
	}
	if value := res.Get("healthMessage"); value.Exists() && !data.InfoHealthMessage.IsNull() {
		data.InfoHealthMessage = types.StringValue(value.String())
	} else {
		data.InfoHealthMessage = types.StringNull()
	}
	if value := res.Get("isConnected"); value.Exists() && !data.InfoIsConnected.IsNull() {
		data.InfoIsConnected = types.BoolValue(value.Bool())
	} else {
		data.InfoIsConnected = types.BoolNull()
	}
	if value := res.Get("deploymentStatus"); value.Exists() && !data.InfoDeploymentStatus.IsNull() {
		data.InfoDeploymentStatus = types.StringValue(value.String())
	} else {
		data.InfoDeploymentStatus = types.StringNull()
	}
	if value := res.Get("ftdMode"); value.Exists() && !data.InfoFtdMode.IsNull() {
		data.InfoFtdMode = types.StringValue(value.String())
	} else {
		data.InfoFtdMode = types.StringNull()
	}
	if value := res.Get("deviceSerialNumber"); value.Exists() && !data.InfoDeviceSerialNumber.IsNull() {
		data.InfoDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.InfoDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("snortVersion"); value.Exists() && !data.InfoSnortVersion.IsNull() {
		data.InfoSnortVersion = types.StringValue(value.String())
	} else {
		data.InfoSnortVersion = types.StringNull()
	}
	if value := res.Get("vdbVersion"); value.Exists() && !data.InfoVdbVersion.IsNull() {
		data.InfoVdbVersion = types.StringValue(value.String())
	} else {
		data.InfoVdbVersion = types.StringNull()
	}
	if value := res.Get("lspVersion"); value.Exists() && !data.InfoLspVersion.IsNull() {
		data.InfoLspVersion = types.StringValue(value.String())
	} else {
		data.InfoLspVersion = types.StringNull()
	}
	if value := res.Get("accessPolicy.name"); value.Exists() && !data.InfoDeployedAccessPolicyName.IsNull() {
		data.InfoDeployedAccessPolicyName = types.StringValue(value.String())
	} else {
		data.InfoDeployedAccessPolicyName = types.StringNull()
	}
	if value := res.Get("healthPolicy.name"); value.Exists() && !data.InfoDeployedHealthPolicyName.IsNull() {
		data.InfoDeployedHealthPolicyName = types.StringValue(value.String())
	} else {
		data.InfoDeployedHealthPolicyName = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

func (data *Device) fromPolicyBody(ctx context.Context, res gjson.Result) {
	query := fmt.Sprintf(`items.#(targets.#(id=="%s"))#.policy`, data.Id.ValueString())
	list := res.Get(query)
	tflog.Debug(ctx, fmt.Sprintf("gjson path %s resulted in %d policies: %s", query, len(list.Array()), list))

	if !list.Exists() {
		data.AccessPolicyId = types.StringNull()
		data.NatPolicyId = types.StringNull()
		data.HealthPolicyId = types.StringNull()
		return
	}

	value := list.Get(`#(type=="AccessPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search AccessPolicy resulted in: %s", value))
	if value.Exists() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="FTDNatPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search FTDNatPolicy resulted in: %s", value))
	if value.Exists() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="HealthPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search HealthPolicy resulted in: %s", value))
	if value.Exists() {
		data.HealthPolicyId = types.StringValue(value.String())
	} else {
		data.HealthPolicyId = types.StringNull()
	}

}

func (data *Device) updateFromPolicyBody(ctx context.Context, res gjson.Result) {
	query := fmt.Sprintf(`items.#(targets.#(id=="%s"))#.policy`, data.Id.ValueString())
	list := res.Get(query)
	tflog.Debug(ctx, fmt.Sprintf("gjson path %s resulted in %d policies for update: %s", query, len(list.Array()), list))

	if !list.Exists() {
		data.AccessPolicyId = types.StringNull()
		data.NatPolicyId = types.StringNull()
		data.HealthPolicyId = types.StringNull()
		return
	}

	value := list.Get(`#(type=="AccessPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search AccessPolicy resulted in: %s", value))
	if value.Exists() && !data.AccessPolicyId.IsNull() {
		data.AccessPolicyId = types.StringValue(value.String())
	} else {
		data.AccessPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="FTDNatPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search FTDNatPolicy resulted in: %s", value))
	if value.Exists() && !data.NatPolicyId.IsNull() {
		data.NatPolicyId = types.StringValue(value.String())
	} else {
		data.NatPolicyId = types.StringNull()
	}

	value = list.Get(`#(type=="HealthPolicy").id`)
	tflog.Debug(ctx, fmt.Sprintf("gjson search HealthPolicy resulted in: %s", value))
	if value.Exists() && !data.HealthPolicyId.IsNull() {
		data.HealthPolicyId = types.StringValue(value.String())
	} else {
		data.HealthPolicyId = types.StringNull()
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Device) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.InfoVersion.IsUnknown() {
		if value := res.Get("sw_version"); value.Exists() {
			data.InfoVersion = types.StringValue(value.String())
		} else {
			data.InfoVersion = types.StringNull()
		}
	}
	if data.InfoHealthStatus.IsUnknown() {
		if value := res.Get("healthStatus"); value.Exists() {
			data.InfoHealthStatus = types.StringValue(value.String())
		} else {
			data.InfoHealthStatus = types.StringNull()
		}
	}
	if data.InfoHealthMessage.IsUnknown() {
		if value := res.Get("healthMessage"); value.Exists() {
			data.InfoHealthMessage = types.StringValue(value.String())
		} else {
			data.InfoHealthMessage = types.StringNull()
		}
	}
	if data.InfoIsConnected.IsUnknown() {
		if value := res.Get("isConnected"); value.Exists() {
			data.InfoIsConnected = types.BoolValue(value.Bool())
		} else {
			data.InfoIsConnected = types.BoolNull()
		}
	}
	if data.InfoDeploymentStatus.IsUnknown() {
		if value := res.Get("deploymentStatus"); value.Exists() {
			data.InfoDeploymentStatus = types.StringValue(value.String())
		} else {
			data.InfoDeploymentStatus = types.StringNull()
		}
	}
	if data.InfoFtdMode.IsUnknown() {
		if value := res.Get("ftdMode"); value.Exists() {
			data.InfoFtdMode = types.StringValue(value.String())
		} else {
			data.InfoFtdMode = types.StringNull()
		}
	}
	if data.InfoDeviceSerialNumber.IsUnknown() {
		if value := res.Get("deviceSerialNumber"); value.Exists() {
			data.InfoDeviceSerialNumber = types.StringValue(value.String())
		} else {
			data.InfoDeviceSerialNumber = types.StringNull()
		}
	}
	if data.InfoSnortVersion.IsUnknown() {
		if value := res.Get("snortVersion"); value.Exists() {
			data.InfoSnortVersion = types.StringValue(value.String())
		} else {
			data.InfoSnortVersion = types.StringNull()
		}
	}
	if data.InfoVdbVersion.IsUnknown() {
		if value := res.Get("vdbVersion"); value.Exists() {
			data.InfoVdbVersion = types.StringValue(value.String())
		} else {
			data.InfoVdbVersion = types.StringNull()
		}
	}
	if data.InfoLspVersion.IsUnknown() {
		if value := res.Get("lspVersion"); value.Exists() {
			data.InfoLspVersion = types.StringValue(value.String())
		} else {
			data.InfoLspVersion = types.StringNull()
		}
	}
	if data.InfoDeployedAccessPolicyName.IsUnknown() {
		if value := res.Get("accessPolicy.name"); value.Exists() {
			data.InfoDeployedAccessPolicyName = types.StringValue(value.String())
		} else {
			data.InfoDeployedAccessPolicyName = types.StringNull()
		}
	}
	if data.InfoDeployedHealthPolicyName.IsUnknown() {
		if value := res.Get("healthPolicy.name"); value.Exists() {
			data.InfoDeployedHealthPolicyName = types.StringValue(value.String())
		} else {
			data.InfoDeployedHealthPolicyName = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
