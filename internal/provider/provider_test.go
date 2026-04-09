// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"log"
	"os"
	"testing"
	"time"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"nd": providerserver.NewProtocol6WithError(New("test")()),
}

func TestMain(m *testing.M) {
	accTest := os.Getenv("TF_ACC")
	if accTest == "" {
		// Unit test run — skip acceptance test initialization
		os.Exit(m.Run())
	}

	// Acceptance tests — initialize ND details from YAML
	testConfigPath := os.Getenv("NDFC_TEST_CONFIG_FILE")
	if testConfigPath == "" {
		panic("NDFC_TEST_CONFIG_FILE env variable not set")
	}

	mockedServer := os.Getenv("ND_MOCKED_SERVER")
	helper.InitConfig(testConfigPath, mockedServer)

	res := m.Run()
	helper.StopMock()
	os.Exit(res)
}

// testAccPreCheck is called by each acceptance test to perform module-specific pre-checks.
// When mocked, it starts the module-specific mock server.
func testAccPreCheck(t *testing.T, module string) {
	t.Helper()
	t.Logf("Starting testAccPreCheck for %s", module)
	if !helper.IsMocked() {
		return
	}
	go helper.StartMockServer(module)
	time.Sleep(10 * time.Second)
	log.Printf("Mock server started for module %s", module)
}
