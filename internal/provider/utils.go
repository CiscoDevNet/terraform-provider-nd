package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func getStringAttributeValue(ctx context.Context, attributeValue basetypes.StringValue, envKey string) string {
	if attributeValue.ValueString() == "" {
		return os.Getenv(envKey)
	}
	return attributeValue.ValueString()
}
