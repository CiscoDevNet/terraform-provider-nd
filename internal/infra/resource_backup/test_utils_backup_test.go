// Code generated;  DO NOT EDIT.

package resource_backup

import (
	"strconv"
	"terraform-provider-nd/internal/infra/resource_backup"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func BackupModelHelperStateCheck(RscName string, c resource_backup.NDFCBackupModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.Type != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("type").String(), c.Type))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("type").String(), "configOnly"))
	}
	if c.Destination != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination").String(), c.Destination))
	}
	if c.EncryptionKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("encryption_key").String(), c.EncryptionKey))
	}
	if c.TelemetryData != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_data").String(), strconv.FormatBool(*c.TelemetryData)))
	}
	return ret
}
