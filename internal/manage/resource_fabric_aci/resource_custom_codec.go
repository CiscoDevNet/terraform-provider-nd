package resource_fabric_aci

import (
	"terraform-provider-nd/internal/common/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NDFCFabricAciManageModel is the payload shape for
// PUT /manage/fabrics/{fabricName}. The generated NDFCFabricAciModel
// is kept for the infra /clusters create/read/delete payload.
type NDFCFabricAciManageModel struct {
	Category                   string                        `json:"category,omitempty"`
	FabricName                 string                        `json:"name,omitempty"`
	LicenseTier                string                        `json:"licenseTier,omitempty"`
	SecurityDomain             string                        `json:"securityDomain,omitempty"`
	TelemetryCollection        *bool                         `json:"telemetryCollection,omitempty"`
	TelemetryCollectionType    string                        `json:"telemetryCollectionType,omitempty"`
	TelemetryStreamingProtocol string                        `json:"telemetryStreamingProtocol,omitempty"`
	Location                   *NDFCSpecLocationValue        `json:"location,omitempty"`
	Management                 *NDFCAciManageManagementValue `json:"management,omitempty"`
}

type NDFCAciManageManagementValue struct {
	Type          string `json:"type,omitempty"`
	Orchestration *bool  `json:"orchestration,omitempty"`
	Epg           string `json:"epg,omitempty"`
}

func (v FabricAciModel) GetManageModelData() *NDFCFabricAciManageModel {
	data := &NDFCFabricAciManageModel{
		Category: "fabric",
		Management: &NDFCAciManageManagementValue{
			Type: "aci",
		},
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	}
	if !v.LicenseTier.IsNull() && !v.LicenseTier.IsUnknown() {
		data.LicenseTier = v.LicenseTier.ValueString()
	}

	if !v.SecurityDomain.IsNull() && !v.SecurityDomain.IsUnknown() {
		data.SecurityDomain = v.SecurityDomain.ValueString()
	}

	if !v.TelemetryStatus.IsNull() && !v.TelemetryStatus.IsUnknown() {
		telemetryCollection := utils.EnabledDisabledToBool(v.TelemetryStatus.ValueString())
		data.TelemetryCollection = &telemetryCollection
	}
	if !v.TelemetryNetwork.IsNull() && !v.TelemetryNetwork.IsUnknown() {
		data.TelemetryCollectionType = telemetryNetworkToCollectionType(v.TelemetryNetwork.ValueString())
	}
	if !v.TelemetryStreamingProtocol.IsNull() && !v.TelemetryStreamingProtocol.IsUnknown() {
		data.TelemetryStreamingProtocol = v.TelemetryStreamingProtocol.ValueString()
	}
	if !v.Latitude.IsNull() && !v.Latitude.IsUnknown() {
		if data.Location == nil {
			data.Location = &NDFCSpecLocationValue{}
		}
		data.Location.Latitude = new(float64)
		*data.Location.Latitude = v.Latitude.ValueFloat64()
	}
	if !v.Longitude.IsNull() && !v.Longitude.IsUnknown() {
		if data.Location == nil {
			data.Location = &NDFCSpecLocationValue{}
		}
		data.Location.Longitude = new(float64)
		*data.Location.Longitude = v.Longitude.ValueFloat64()
	}
	if !v.OrchestrationStatus.IsNull() && !v.OrchestrationStatus.IsUnknown() {
		orchestration := utils.EnabledDisabledToBool(v.OrchestrationStatus.ValueString())
		data.Management.Orchestration = &orchestration
	}
	if !v.TelemetryEpg.IsNull() && !v.TelemetryEpg.IsUnknown() {
		data.Management.Epg = v.TelemetryEpg.ValueString()
	}

	return data
}

func (v *FabricAciModel) NormalizeTelemetryNetworkState() {
	if v == nil || v.TelemetryNetwork.IsNull() || v.TelemetryNetwork.IsUnknown() {
		return
	}

	v.TelemetryNetwork = types.StringValue(normalizeTelemetryNetworkState(v.TelemetryNetwork.ValueString()))
}

func telemetryNetworkToCollectionType(network string) string {
	switch normalizeTelemetryNetworkState(network) {
	case "inband":
		return "inBand"
	case "outband":
		return "outOfBand"
	}
	return network
}

func normalizeTelemetryNetworkState(network string) string {
	switch network {
	case "inband", "inBand":
		return "inband"
	case "outband", "outOfBand":
		return "outband"
	}
	return network
}
