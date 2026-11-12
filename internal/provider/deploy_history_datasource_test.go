package provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDeployHistoryDataSource(test *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		test.Skip("Set TF_ACC to run acceptance tests")
	}
	config := helper.GetConfig("global")
	if config.ND.Fabric == "" {
		test.Skip("Need nd.fabric in the testbed config")
	}
	const sourceName = "data.nd_deploy_history.test"
	for _, recordLimit := range []int{1, 5} {
		test.Run(fmt.Sprintf("limit_%d", recordLimit), func(test *testing.T) {
			resource.Test(test, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(test, "global") },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{{
					Config: fmt.Sprintf(`
provider "nd" {
  username = %q
  password = %q
  url      = %q
  insecure = %t
}

data "nd_deploy_history" "test" {
  fabric_name   = %q
  no_of_records = %d
}
`, config.ND.User, config.ND.Password, config.ND.URL, config.ND.Insecure == "true", config.ND.Fabric, recordLimit),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(sourceName, "fabric_name", config.ND.Fabric),
						resource.TestCheckResourceAttr(sourceName, "no_of_records", strconv.Itoa(recordLimit)),
						func(state *terraform.State) error {
							saved, exists := state.RootModule().Resources[sourceName]
							if !exists || saved.Primary == nil {
								return fmt.Errorf("missing state for %s", sourceName)
							}
							countValue := saved.Primary.Attributes["deployment_records.#"]
							if countValue == "" {
								return nil
							}
							count, err := strconv.Atoi(countValue)
							if err != nil || count < 0 || count > recordLimit {
								return fmt.Errorf("expected at most %d records, got %q", recordLimit, countValue)
							}
							return nil
						},
					),
				}},
			})
		})
	}
}
