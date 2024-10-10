package provider

import (
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"nd": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ND_USERNAME"); v == "" {
		t.Fatal("ND_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("ND_PASSWORD"); v == "" {
		t.Fatal("ND_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("ND_URL"); v == "" {
		t.Fatal("ND_URL must be set for acceptance tests")
	}
	if v := os.Getenv("ND_VAL_REL_DN"); v == "" {
		t.Fatal("ND_VAL_REL_DN must be set for acceptance tests")
		boolValue, err := strconv.ParseBool(v)
		if err != nil || boolValue == true {
			t.Fatal("ND_VAL_REL_DN must be a 'false' boolean value")
		}
	}
}
