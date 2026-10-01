package resource_remote_storage_location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Bool = remoteStorageHostKeyBoolValidator{}

type remoteStorageHostKeyBoolValidator struct{}

func remoteStorageHostKeyValidator() validator.Bool {
	return remoteStorageHostKeyBoolValidator{}
}

func (remoteStorageHostKeyBoolValidator) Description(_ context.Context) string {
	return "accept_host_key and ignore_host_key_validation cannot both be true"
}

func (v remoteStorageHostKeyBoolValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (remoteStorageHostKeyBoolValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || !req.ConfigValue.ValueBool() {
		return
	}

	var ignoreHostKeyValidation types.Bool
	diags := req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName("ignore_host_key_validation"), &ignoreHostKeyValidation)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || ignoreHostKeyValidation.IsNull() || ignoreHostKeyValidation.IsUnknown() || !ignoreHostKeyValidation.ValueBool() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid host-key configuration",
		"Attributes `scp_sftp.accept_host_key` and `scp_sftp.ignore_host_key_validation` cannot both be true.",
	)
}
