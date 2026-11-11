// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_config_deploy"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/manage/resource_interface_common"
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
	"terraform-provider-nd/internal/manage/resource_vpc_pair"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/netascode/go-nd"
)

const integratedFabricResourceName = "nd_fabric_vxlan.fabric"

const (
	integratedStageFabric = iota + 1
	integratedStageSwitches
	integratedStageVpcPairs
	integratedStageInitialVpcDeploy
	integratedStageRemoveSecondVpcPair
	integratedStageDeployVpcRemoval
	integratedStageRestoreSecondVpcPair
	integratedStageDeployVpcRestore
	integratedStageInitialInterfaces
	integratedStageDeployInitialInterfaces
	integratedStageMutateInterfaces
	integratedStageDeployInterfaceMutations
	integratedStageSettleDeployBehavior
	integratedStageNoOpPlan
)

var integratedStageNames = map[int]string{
	integratedStageFabric:                   "create fabric",
	integratedStageSwitches:                 "add four switches",
	integratedStageVpcPairs:                 "create two vPC pairs without deploy",
	integratedStageInitialVpcDeploy:         "deploy both vPC pairs",
	integratedStageRemoveSecondVpcPair:      "remove second vPC pair without deploy",
	integratedStageDeployVpcRemoval:         "deploy second vPC pair removal",
	integratedStageRestoreSecondVpcPair:     "restore second vPC pair without deploy",
	integratedStageDeployVpcRestore:         "deploy restored second vPC pair",
	integratedStageInitialInterfaces:        "create initial interfaces without deploy",
	integratedStageDeployInitialInterfaces:  "deploy initial interfaces",
	integratedStageMutateInterfaces:         "add update and delete interfaces without deploy",
	integratedStageDeployInterfaceMutations: "deploy interface mutations",
	integratedStageSettleDeployBehavior:     "disable always-execute after lifecycle",
	integratedStageNoOpPlan:                 "verify settled no-op plan",
}

type integratedFabricTestResources struct {
	fabric              *resource_fabric_common.NDFCFabricCommonModel
	switches            []*resource_inventory_switch.NDFCInventorySwitchModel
	switchIndexBySerial map[string]int
	pairIndexBySerial   map[string]int
	vpcPairs            []*resource_vpc_pair.NDFCVpcPairModel
	initialInterfaces   map[int]*resource_interface_common.NDFCInterfaceCommonModel
	mutatedInterfaces   map[int]*resource_interface_common.NDFCInterfaceCommonModel
	interfaceConfig     map[string]helper.IntegratedInterface
}

type integratedInterfaceSummaryEntry struct {
	InterfaceName    string
	OperStatus       string
	OperState        string
	LinkState        string
	ConfigSyncStatus string
}

type integratedInterfaceSummaryResponse struct {
	Interfaces []integratedInterfaceSummaryEntry
}

type integratedFabricSwitchEntry struct {
	SerialNumber  string
	VpcConfigured bool
	VpcData       struct {
		PeerSwitchID string
	}
}

type integratedFabricSwitchesResponse struct {
	Switches []integratedFabricSwitchEntry
}

// TestAccIntegratedFabricTwoVpcPairsDeploy exercises a full fabric lifecycle:
// create and deploy two pairs, remove/deploy/restore/deploy one pair, then
// create and mutate Ethernet interfaces with explicit deployment checkpoints.
func TestAccIntegratedFabricTwoVpcPairsDeploy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("acceptance test requires TF_ACC=1")
	}
	cfg := helper.GetConfig("global")
	integration := validateIntegratedFabricConfig(t, cfg.ND.IntegrationTest)
	resources := buildIntegratedFabricResources(integration)

	providerConfig := map[string]string{
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	render := func(stage int) string {
		t.Helper()
		models, names := integratedFabricStage(resources, integration.Fabric, stage)
		providerConfig["RscName"] = strings.Join(names, ",")
		var tfConfig *string
		helper.GetTFConfigWithSingleResource(
			fmt.Sprintf("%s_%02d", t.Name(), stage),
			providerConfig,
			models,
			&tfConfig,
		)
		return *tfConfig
	}

	var pairOneID string
	steps := make([]resource.TestStep, 0, integratedStageNoOpPlan)
	for stage := integratedStageFabric; stage <= integratedStageNoOpPlan; stage++ {
		stage := stage
		checks := integratedFabricStateChecks(resources, stage)

		if stage == integratedStageSwitches {
			checks = append(checks, checkIntegratedInterfaceCandidates(cfg, integration))
		}
		if stage >= integratedStageVpcPairs {
			checks = append(checks, checkIntegratedVpcEndpoints(cfg, resources, stage))
		}
		if stage == integratedStageInitialVpcDeploy {
			checks = append(checks,
				captureIntegratedResourceID("nd_vpc_pair.vpc_pair_1", &pairOneID),
				checkIntegratedVpcDeployment(cfg, resources, true),
			)
		}
		if stage > integratedStageInitialVpcDeploy {
			checks = append(checks, checkIntegratedResourceID("nd_vpc_pair.vpc_pair_1", &pairOneID))
		}
		if stage == integratedStageDeployVpcRemoval {
			checks = append(checks, checkIntegratedVpcDeployment(cfg, resources, false))
		}
		if stage == integratedStageDeployVpcRestore {
			checks = append(checks, checkIntegratedVpcDeployment(cfg, resources, true))
		}
		if stage == integratedStageDeployInitialInterfaces {
			checks = append(checks, checkIntegratedInterfacesDeployed(cfg, resources.initialInterfaces))
		}
		if stage == integratedStageDeployInterfaceMutations {
			checks = append(checks, checkIntegratedInterfacesDeployed(cfg, resources.mutatedInterfaces))
		}

		step := resource.TestStep{
			PreConfig: func() {
				t.Logf("Integrated fabric TC%02d: %s", stage, integratedStageNames[stage])
			},
			Config: render(stage),
		}
		if stage != integratedStageNoOpPlan {
			step.Check = resource.ComposeTestCheckFunc(checks...)
		}
		if stage >= integratedStageInitialVpcDeploy && stage < integratedStageSettleDeployBehavior {
			// always_execute intentionally leaves a refresh diff through TC12.
			// TC13 applies the transition to the final stable false value and is
			// expected to be empty after that apply.
			step.ExpectNonEmptyPlan = true
		}
		if stage == integratedStageNoOpPlan {
			// TC13 already applied always_execute=false. Reusing that identical
			// configuration here must now produce an empty plan.
			step.PlanOnly = true
			step.ExpectNonEmptyPlan = false
		}
		steps = append(steps, step)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			requireIntegratedFabricAbsent(t, cfg, integration.Fabric)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
		CheckDestroy:             checkIntegratedFabricDestroyed(cfg, integration.Fabric),
	})
}

func validateIntegratedFabricConfig(t *testing.T, cfg helper.IntegratedConfig) helper.IntegratedConfig {
	t.Helper()

	if cfg.Fabric == "" || cfg.BgpAsn == "" || cfg.FabricType == "" {
		t.Fatal("integration_test requires fabric, bgp_asn, and fabric_type")
	}
	if cfg.User == "" || cfg.Password == "" {
		t.Fatal("integration_test requires discovery switch user and pwd")
	}
	if len(cfg.Switches) != 4 {
		t.Fatalf("integration_test requires exactly 4 switches, got %d", len(cfg.Switches))
	}
	if len(cfg.VpcPairs) != 2 {
		t.Fatalf("integration_test requires exactly 2 vpc_pairs, got %d", len(cfg.VpcPairs))
	}
	if len(cfg.Interfaces) != 4 {
		t.Fatalf("integration_test requires exactly 4 interfaces, got %d", len(cfg.Interfaces))
	}

	switches := make(map[string]helper.InventorySwitch, len(cfg.Switches))
	for i, sw := range cfg.Switches {
		if sw.Serial == "" || sw.IP == "" {
			t.Fatalf("integration_test.switches[%d] requires serial and ip", i)
		}
		if !strings.EqualFold(sw.Role, "leaf") {
			t.Fatalf("integration_test.switches[%d] must have role leaf, got %q", i, sw.Role)
		}
		if _, exists := switches[sw.Serial]; exists {
			t.Fatalf("integration_test switch serial %q is duplicated", sw.Serial)
		}
		switches[sw.Serial] = sw
	}

	paired := make(map[string]bool, len(cfg.Switches))
	for i, pair := range cfg.VpcPairs {
		if pair.SwitchID1 == pair.SwitchID2 {
			t.Fatalf("integration_test.vpc_pairs[%d] uses the same switch twice", i)
		}
		for _, serial := range []string{pair.SwitchID1, pair.SwitchID2} {
			if _, exists := switches[serial]; !exists {
				t.Fatalf("integration_test.vpc_pairs[%d] references unknown switch %q", i, serial)
			}
			if paired[serial] {
				t.Fatalf("integration_test switch %q belongs to more than one vPC pair", serial)
			}
			paired[serial] = true
		}
	}

	configuredInterfaces := make(map[string]bool, len(cfg.Interfaces))
	for i, intf := range cfg.Interfaces {
		if _, exists := switches[intf.SwitchID]; !exists {
			t.Fatalf("integration_test.interfaces[%d] references unknown switch %q", i, intf.SwitchID)
		}
		if strings.TrimSpace(intf.InterfaceName) == "" {
			t.Fatalf("integration_test.interfaces[%d] requires interface_name", i)
		}
		if configuredInterfaces[intf.SwitchID] {
			t.Fatalf("integration_test switch %q has more than one reserved interface entry", intf.SwitchID)
		}
		configuredInterfaces[intf.SwitchID] = true
	}
	for serial := range switches {
		if !configuredInterfaces[serial] {
			t.Fatalf("integration_test switch %q has no reserved interface", serial)
		}
	}

	return cfg
}

func buildIntegratedFabricResources(cfg helper.IntegratedConfig) integratedFabricTestResources {
	result := integratedFabricTestResources{
		fabric:              new(resource_fabric_common.NDFCFabricCommonModel),
		switches:            make([]*resource_inventory_switch.NDFCInventorySwitchModel, len(cfg.Switches)),
		switchIndexBySerial: make(map[string]int, len(cfg.Switches)),
		pairIndexBySerial:   make(map[string]int, len(cfg.Switches)),
		vpcPairs:            make([]*resource_vpc_pair.NDFCVpcPairModel, len(cfg.VpcPairs)),
		initialInterfaces:   make(map[int]*resource_interface_common.NDFCInterfaceCommonModel),
		mutatedInterfaces:   make(map[int]*resource_interface_common.NDFCInterfaceCommonModel),
		interfaceConfig:     make(map[string]helper.IntegratedInterface, len(cfg.Interfaces)),
	}

	helper.GenerateFabricVxlanObject(&result.fabric, cfg.Fabric, cfg.BgpAsn, cfg.FabricType, nil)
	for i, sw := range cfg.Switches {
		result.switches[i] = new(resource_inventory_switch.NDFCInventorySwitchModel)
		helper.GenerateInventorySwitchFromConfig(&result.switches[i], cfg.Fabric, cfg.User, cfg.Password, sw)
		result.switchIndexBySerial[sw.Serial] = i + 1
	}
	for i, pair := range cfg.VpcPairs {
		result.vpcPairs[i] = new(resource_vpc_pair.NDFCVpcPairModel)
		helper.GenerateVpcPairObject(
			&result.vpcPairs[i], cfg.Fabric, pair.SwitchID1, pair.SwitchID2,
			pair.UseVirtualPeerlink, false,
		)
		result.pairIndexBySerial[pair.SwitchID1] = i + 1
		result.pairIndexBySerial[pair.SwitchID2] = i + 1
	}
	for _, intf := range cfg.Interfaces {
		result.interfaceConfig[intf.SwitchID] = intf
	}

	result.initialInterfaces[1] = integratedAccessInterface(
		cfg.Fabric, cfg.Switches[0].Serial, result.interfaceConfig[cfg.Switches[0].Serial].InterfaceName,
		"integration initial leaf7", "",
	)
	result.initialInterfaces[2] = integratedAccessInterface(
		cfg.Fabric, cfg.Switches[1].Serial, result.interfaceConfig[cfg.Switches[1].Serial].InterfaceName,
		"integration initial leaf8", "",
	)
	result.initialInterfaces[3] = integratedTrunkInterface(
		cfg.Fabric, cfg.Switches[2].Serial, result.interfaceConfig[cfg.Switches[2].Serial].InterfaceName,
		"integration control leaf11",
	)

	result.mutatedInterfaces[1] = integratedAccessInterface(
		cfg.Fabric, cfg.Switches[0].Serial, result.interfaceConfig[cfg.Switches[0].Serial].InterfaceName,
		"integration updated leaf7", "jumbo",
	)
	result.mutatedInterfaces[3] = result.initialInterfaces[3]
	result.mutatedInterfaces[4] = integratedAccessInterface(
		cfg.Fabric, cfg.Switches[3].Serial, result.interfaceConfig[cfg.Switches[3].Serial].InterfaceName,
		"integration added leaf12", "",
	)

	return result
}

func integratedAccessInterface(fabricName, switchID, interfaceName, description, mtu string) *resource_interface_common.NDFCInterfaceCommonModel {
	model := new(resource_interface_common.NDFCInterfaceCommonModel)
	adminState := false
	accessVlan := int64(1)
	helper.GenerateInterfaceEthernetObject(
		&model,
		fabricName,
		switchID,
		false,
		map[string]helper.InterfaceEthernetEntry{
			interfaceName: {
				Mode:          "access",
				NetworkOsType: "nx-os",
				PolicyType:    "accessHost",
				AdminState:    &adminState,
				Description:   description,
				Mtu:           mtu,
				AccessVlan:    &accessVlan,
			},
		},
	)
	return model
}

func integratedTrunkInterface(fabricName, switchID, interfaceName, description string) *resource_interface_common.NDFCInterfaceCommonModel {
	model := new(resource_interface_common.NDFCInterfaceCommonModel)
	adminState := false
	helper.GenerateInterfaceEthernetObject(
		&model,
		fabricName,
		switchID,
		false,
		map[string]helper.InterfaceEthernetEntry{
			interfaceName: {
				Mode:          "trunk",
				NetworkOsType: "nx-os",
				PolicyType:    "trunkHost",
				AdminState:    &adminState,
				Description:   description,
				AllowedVlans:  "1",
			},
		},
	)
	return model
}

func integratedFabricStage(r integratedFabricTestResources, fabricName string, stage int) ([]interface{}, []string) {
	models := []interface{}{helper.VxlanResource(r.fabric)}
	names := []string{"fabric"}
	if stage < integratedStageSwitches {
		return models, names
	}

	for i, sw := range r.switches {
		models = append(models, helper.WithDependsOn(sw, integratedFabricResourceName))
		names = append(names, fmt.Sprintf("switch_%d", i+1))
	}
	if stage < integratedStageVpcPairs {
		return models, names
	}

	pairIndexes := []int{0, 1}
	if stage == integratedStageRemoveSecondVpcPair || stage == integratedStageDeployVpcRemoval {
		pairIndexes = []int{0}
	}
	for _, i := range pairIndexes {
		pair := r.vpcPairs[i]
		switch1 := r.switchIndexBySerial[pair.SwitchId1]
		switch2 := r.switchIndexBySerial[pair.SwitchId2]
		models = append(models, helper.WithDependsOn(
			pair,
			fmt.Sprintf("nd_inventory_switch.switch_%d", switch1),
			fmt.Sprintf("nd_inventory_switch.switch_%d", switch2),
		))
		names = append(names, fmt.Sprintf("vpc_pair_%d", i+1))
	}

	activeInterfaces := map[int]*resource_interface_common.NDFCInterfaceCommonModel(nil)
	if stage == integratedStageInitialInterfaces || stage == integratedStageDeployInitialInterfaces {
		activeInterfaces = r.initialInterfaces
	}
	if stage >= integratedStageMutateInterfaces {
		activeInterfaces = r.mutatedInterfaces
	}
	for _, switchIndex := range sortedIntegratedInterfaceIndexes(activeInterfaces) {
		intf := activeInterfaces[switchIndex]
		pairIndex := r.pairIndexBySerial[intf.SwitchId]
		models = append(models, helper.WithDependsOn(
			intf,
			fmt.Sprintf("nd_vpc_pair.vpc_pair_%d", pairIndex),
		))
		names = append(names, fmt.Sprintf("interface_%d", switchIndex))
	}

	if stage >= integratedStageInitialVpcDeploy {
		executeOnUpdate := integratedDeployStage(stage)
		alwaysExecute := stage < integratedStageSettleDeployBehavior
		deployModel := new(resource_config_deploy.NDFCConfigDeployModel)
		helper.GenerateConfigDeployObject(
			&deployModel,
			fabricName,
			true,
			true,
			nil,
			&resource_config_deploy.NDFCBehaviourValue{
				ExecuteOnCreate:  helper.BoolPtr(true),
				ExecuteOnUpdate:  helper.BoolPtr(executeOnUpdate),
				ExecuteOnDestroy: helper.BoolPtr(false),
			},
		)
		helper.ModifyConfigDeployObject(&deployModel, map[string]interface{}{
			"always_execute": alwaysExecute,
		})

		dependencies := integratedDeployDependencies(pairIndexes, activeInterfaces)
		models = append(models, helper.WithDependsOn(deployModel, dependencies...))
		names = append(names, "deploy")
	}

	return models, names
}

func integratedDeployStage(stage int) bool {
	switch stage {
	case integratedStageInitialVpcDeploy,
		integratedStageDeployVpcRemoval,
		integratedStageDeployVpcRestore,
		integratedStageDeployInitialInterfaces,
		integratedStageDeployInterfaceMutations:
		return true
	default:
		return false
	}
}

func integratedDeployDependencies(pairIndexes []int, interfaces map[int]*resource_interface_common.NDFCInterfaceCommonModel) []string {
	dependencies := make([]string, 0, len(pairIndexes)+len(interfaces))
	for _, i := range pairIndexes {
		dependencies = append(dependencies, fmt.Sprintf("nd_vpc_pair.vpc_pair_%d", i+1))
	}
	for _, switchIndex := range sortedIntegratedInterfaceIndexes(interfaces) {
		dependencies = append(dependencies, fmt.Sprintf("nd_interface_ethernet.interface_%d", switchIndex))
	}
	return dependencies
}

func sortedIntegratedInterfaceIndexes(interfaces map[int]*resource_interface_common.NDFCInterfaceCommonModel) []int {
	indexes := make([]int, 0, len(interfaces))
	for index := range interfaces {
		indexes = append(indexes, index)
	}
	sort.Ints(indexes)
	return indexes
}

func integratedFabricStateChecks(r integratedFabricTestResources, stage int) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(integratedFabricResourceName, "fabric_name", r.fabric.FabricName),
		resource.TestCheckResourceAttr(integratedFabricResourceName, "bgp_asn", r.fabric.Management.BgpAsn),
		resource.TestCheckResourceAttr(integratedFabricResourceName, "fabric_type", r.fabric.Management.FabricType),
		resource.TestCheckResourceAttrSet(integratedFabricResourceName, "id"),
	}
	if stage < integratedStageSwitches {
		return checks
	}

	for i, sw := range r.switches {
		address := fmt.Sprintf("nd_inventory_switch.switch_%d", i+1)
		checks = append(checks,
			resource.TestCheckResourceAttr(address, "fabric_name", sw.FabricName),
			resource.TestCheckResourceAttr(address, "switch_detail.serial_number", sw.SwitchDetail.SerialNumber),
			resource.TestCheckResourceAttr(address, "switch_detail.ip_address", sw.SwitchDetail.IpAddress),
			resource.TestCheckResourceAttr(address, "switch_detail.switch_role", sw.SwitchDetail.SwitchRole),
			resource.TestCheckResourceAttrSet(address, "switch_detail.model"),
			resource.TestCheckResourceAttrSet(address, "switch_detail.software_version"),
			resource.TestCheckResourceAttrSet(address, "id"),
		)
	}
	if stage < integratedStageVpcPairs {
		return checks
	}

	pairIndexes := []int{0, 1}
	if stage == integratedStageRemoveSecondVpcPair || stage == integratedStageDeployVpcRemoval {
		pairIndexes = []int{0}
		checks = append(checks, checkIntegratedResourceAbsent("nd_vpc_pair.vpc_pair_2"))
	}
	for _, i := range pairIndexes {
		pair := r.vpcPairs[i]
		address := fmt.Sprintf("nd_vpc_pair.vpc_pair_%d", i+1)
		checks = append(checks,
			resource.TestCheckResourceAttr(address, "fabric_name", pair.FabricName),
			resource.TestCheckResourceAttr(address, "switch_id_1", pair.SwitchId1),
			resource.TestCheckResourceAttr(address, "switch_id_2", pair.SwitchId2),
			resource.TestCheckResourceAttr(address, "deploy", "false"),
			resource.TestCheckResourceAttrSet(address, "id"),
		)
	}

	activeInterfaces := map[int]*resource_interface_common.NDFCInterfaceCommonModel(nil)
	if stage == integratedStageInitialInterfaces || stage == integratedStageDeployInitialInterfaces {
		activeInterfaces = r.initialInterfaces
	}
	if stage >= integratedStageMutateInterfaces {
		activeInterfaces = r.mutatedInterfaces
	}
	for _, switchIndex := range sortedIntegratedInterfaceIndexes(activeInterfaces) {
		checks = append(checks, integratedInterfaceStateChecks(
			fmt.Sprintf("nd_interface_ethernet.interface_%d", switchIndex),
			activeInterfaces[switchIndex],
		)...)
	}
	if stage >= integratedStageMutateInterfaces {
		checks = append(checks, checkIntegratedResourceAbsent("nd_interface_ethernet.interface_2"))
	}

	if stage >= integratedStageInitialVpcDeploy {
		alwaysExecute := stage < integratedStageSettleDeployBehavior
		checks = append(checks,
			resource.TestCheckResourceAttr("nd_config_deploy.deploy", "fabric_name", r.fabric.FabricName),
			resource.TestCheckResourceAttr("nd_config_deploy.deploy", "config_save", "true"),
			resource.TestCheckResourceAttr("nd_config_deploy.deploy", "deploy", "true"),
			resource.TestCheckResourceAttr("nd_config_deploy.deploy", "always_execute", fmt.Sprintf("%t", alwaysExecute)),
			resource.TestCheckResourceAttrSet("nd_config_deploy.deploy", "id"),
		)
		if integratedDeployStage(stage) {
			checks = append(checks, resource.TestCheckResourceAttrSet("nd_config_deploy.deploy", "status"))
		}
	}

	return checks
}

func integratedInterfaceStateChecks(address string, model *resource_interface_common.NDFCInterfaceCommonModel) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(address, "fabric_name", model.FabricName),
		resource.TestCheckResourceAttr(address, "switch_id", model.SwitchId),
		resource.TestCheckResourceAttrSet(address, "id"),
	}
	for name, intf := range model.Interfaces {
		prefix := "interfaces." + name
		policy := intf.ConfigData.NetworkOS.Policy
		checks = append(checks,
			resource.TestCheckResourceAttr(address, prefix+".mode", intf.ConfigData.Mode),
			resource.TestCheckResourceAttr(address, prefix+".network_os_type", intf.ConfigData.NetworkOS.NetworkOsType),
			resource.TestCheckResourceAttr(address, prefix+".policy_type", policy.PolicyType),
		)
		if policy.AdminState != nil {
			checks = append(checks,
				resource.TestCheckResourceAttr(address, prefix+".admin_state", fmt.Sprintf("%t", *policy.AdminState)),
			)
		}
		if policy.Description != "" {
			checks = append(checks, resource.TestCheckResourceAttr(address, prefix+".description", policy.Description))
		}
		if policy.Mtu != "" {
			checks = append(checks, resource.TestCheckResourceAttr(address, prefix+".mtu", policy.Mtu))
		}
		if policy.AccessVlan != nil {
			checks = append(checks,
				resource.TestCheckResourceAttr(address, prefix+".access_vlan", fmt.Sprintf("%d", *policy.AccessVlan)),
			)
		}
		if policy.AllowedVlans != "" {
			checks = append(checks, resource.TestCheckResourceAttr(address, prefix+".allowed_vlans", policy.AllowedVlans))
		}
	}
	return checks
}

func checkIntegratedResourceAbsent(address string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if _, exists := state.RootModule().Resources[address]; exists {
			return fmt.Errorf("resource %q unexpectedly remains in Terraform state", address)
		}
		return nil
	}
}

func captureIntegratedResourceID(address string, target *string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, exists := state.RootModule().Resources[address]
		if !exists || resourceState.Primary == nil || resourceState.Primary.ID == "" {
			return fmt.Errorf("resource %q has no ID to capture", address)
		}
		*target = resourceState.Primary.ID
		return nil
	}
}

func checkIntegratedResourceID(address string, expected *string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceState, exists := state.RootModule().Resources[address]
		if !exists || resourceState.Primary == nil {
			return fmt.Errorf("resource %q is missing while checking its stable ID", address)
		}
		if *expected == "" {
			return fmt.Errorf("resource %q baseline ID was not captured", address)
		}
		if resourceState.Primary.ID != *expected {
			return fmt.Errorf("resource %q ID changed: expected %q, got %q", address, *expected, resourceState.Primary.ID)
		}
		return nil
	}
}

func newIntegratedNDClient(cfg *helper.Config) (*nd.Client, error) {
	client, err := nd.NewClient(
		cfg.ND.URL, "/api/v1", cfg.ND.User, cfg.ND.Password, "",
		cfg.ND.Insecure == "true", nd.MaxRetries(3),
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func requireIntegratedFabricAbsent(t *testing.T, cfg *helper.Config, fabricName string) {
	t.Helper()
	client, err := newIntegratedNDClient(cfg)
	if err != nil {
		t.Fatalf("failed to create ND client for integrated fabric precheck: %v", err)
	}
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	_, err = fabricAPI.Get()
	if err == nil {
		t.Fatalf("integration test requires fabric %q to be absent before it starts", fabricName)
	}
	if !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to verify that integration fabric %q is absent: %v", fabricName, err)
	}
}

func checkIntegratedFabricDestroyed(cfg *helper.Config, fabricName string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		client, err := newIntegratedNDClient(cfg)
		if err != nil {
			return fmt.Errorf("create ND client for destroy check: %w", err)
		}
		fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
		fabricAPI.FabricName = fabricName
		_, err = fabricAPI.Get()
		if err == nil {
			return fmt.Errorf("fabric %q still exists after Terraform destroy", fabricName)
		}
		if !strings.Contains(err.Error(), "StatusCode 404") {
			return fmt.Errorf("verify fabric %q destroy: %w", fabricName, err)
		}
		return nil
	}
}

func checkIntegratedVpcEndpoints(cfg *helper.Config, r integratedFabricTestResources, stage int) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		client, err := newIntegratedNDClient(cfg)
		if err != nil {
			return err
		}
		for i, pair := range r.vpcPairs {
			expected := true
			if i == 1 && (stage == integratedStageRemoveSecondVpcPair || stage == integratedStageDeployVpcRemoval) {
				expected = false
			}
			pairAPI := api.NewVpcPairAPI(client, pair.FabricName)
			pairAPI.FabricName = pair.FabricName
			pairAPI.SwitchID = pair.SwitchId2
			_, getErr := pairAPI.Get()
			if expected && getErr != nil {
				return fmt.Errorf("expected vPC pair %d to exist: %w", i+1, getErr)
			}
			if !expected {
				if getErr == nil {
					return fmt.Errorf("vPC pair %d still exists after unpair", i+1)
				}
				if !strings.Contains(getErr.Error(), "StatusCode 404") {
					return fmt.Errorf("verify vPC pair %d absence: %w", i+1, getErr)
				}
			}
		}
		return nil
	}
}

func checkIntegratedVpcDeployment(cfg *helper.Config, r integratedFabricTestResources, expectSecondPair bool) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		client, err := newIntegratedNDClient(cfg)
		if err != nil {
			return err
		}
		return waitForIntegratedRemoteState("vPC deployment state", func() error {
			inventoryAPI := api.NewInventoryAPI(client, r.fabric.FabricName)
			inventoryAPI.FabricName = r.fabric.FabricName
			inventoryAPI.SetOperation(api.OpGetAllSwitches)
			payload, err := inventoryAPI.Get()
			if err != nil {
				return fmt.Errorf("read fabric switches for vPC deployment check: %w", err)
			}
			var response integratedFabricSwitchesResponse
			if err := json.Unmarshal(payload, &response); err != nil {
				return fmt.Errorf("decode fabric switches for vPC deployment check: %w", err)
			}
			bySerial := make(map[string]integratedFabricSwitchEntry, len(response.Switches))
			for _, sw := range response.Switches {
				bySerial[sw.SerialNumber] = sw
			}
			for i, pair := range r.vpcPairs {
				expected := i == 0 || expectSecondPair
				for _, endpoint := range []struct {
					serial string
					peer   string
				}{
					{serial: pair.SwitchId1, peer: pair.SwitchId2},
					{serial: pair.SwitchId2, peer: pair.SwitchId1},
				} {
					actual, ok := bySerial[endpoint.serial]
					if !ok {
						return fmt.Errorf("switch %q missing from fabric inventory", endpoint.serial)
					}
					if expected && (!actual.VpcConfigured || actual.VpcData.PeerSwitchID != endpoint.peer) {
						return fmt.Errorf("switch %q expected deployed vPC peer %q, got configured=%t peer=%q",
							endpoint.serial, endpoint.peer, actual.VpcConfigured, actual.VpcData.PeerSwitchID)
					}
					if !expected && (actual.VpcConfigured || actual.VpcData.PeerSwitchID != "") {
						return fmt.Errorf("switch %q unexpectedly retains vPC state configured=%t peer=%q",
							endpoint.serial, actual.VpcConfigured, actual.VpcData.PeerSwitchID)
					}
				}
			}
			return nil
		})
	}
}

func checkIntegratedInterfaceCandidates(cfg *helper.Config, integration helper.IntegratedConfig) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		client, err := newIntegratedNDClient(cfg)
		if err != nil {
			return err
		}
		for _, candidate := range integration.Interfaces {
			summary, err := readIntegratedInterfaceSummary(client, integration.Fabric, candidate.SwitchID)
			if err != nil {
				return err
			}
			found := false
			available := make([]string, 0)
			for _, intf := range summary.Interfaces {
				status := firstNonEmpty(intf.OperStatus, intf.OperState, intf.LinkState)
				if !strings.EqualFold(status, "up") && !strings.EqualFold(status, "connected") {
					available = append(available, intf.InterfaceName)
				}
				if intf.InterfaceName != candidate.InterfaceName {
					continue
				}
				found = true
				if strings.EqualFold(status, "up") || strings.EqualFold(status, "connected") {
					return fmt.Errorf("reserved interface %s/%s is active with status %q",
						candidate.SwitchID, candidate.InterfaceName, status)
				}
			}
			if !found {
				sort.Strings(available)
				if len(available) > 12 {
					available = available[:12]
				}
				return fmt.Errorf("reserved interface %s/%s was not found; sample inactive interfaces: %v",
					candidate.SwitchID, candidate.InterfaceName, available)
			}
		}
		return nil
	}
}

func checkIntegratedInterfacesDeployed(cfg *helper.Config, interfaces map[int]*resource_interface_common.NDFCInterfaceCommonModel) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		client, err := newIntegratedNDClient(cfg)
		if err != nil {
			return err
		}
		return waitForIntegratedRemoteState("interface deployment state", func() error {
			for _, switchIndex := range sortedIntegratedInterfaceIndexes(interfaces) {
				model := interfaces[switchIndex]
				summary, err := readIntegratedInterfaceSummary(client, model.FabricName, model.SwitchId)
				if err != nil {
					return err
				}
				statusByName := make(map[string]string, len(summary.Interfaces))
				for _, intf := range summary.Interfaces {
					statusByName[intf.InterfaceName] = intf.ConfigSyncStatus
				}
				for name := range model.Interfaces {
					status, exists := statusByName[name]
					if !exists {
						return fmt.Errorf("deployed interface %s/%s missing from interface summary", model.SwitchId, name)
					}
					if !strings.EqualFold(status, "inSync") && !strings.EqualFold(status, "deployed") {
						return fmt.Errorf("interface %s/%s is not deployed: configSyncStatus=%q", model.SwitchId, name, status)
					}
				}
			}
			return nil
		})
	}
}

func waitForIntegratedRemoteState(description string, check func() error) error {
	const (
		timeout      = 5 * time.Minute
		pollInterval = 15 * time.Second
	)

	deadline := time.Now().Add(timeout)
	for {
		err := check()
		if err == nil {
			return nil
		}
		if time.Now().Add(pollInterval).After(deadline) {
			return fmt.Errorf("timed out after %s waiting for %s: %w", timeout, description, err)
		}
		time.Sleep(pollInterval)
	}
}

func readIntegratedInterfaceSummary(client *nd.Client, fabricName, switchID string) (integratedInterfaceSummaryResponse, error) {
	interfaceAPI := api.NewInterfaceAPI(client, fabricName)
	interfaceAPI.FabricName = fabricName
	interfaceAPI.SwitchId = switchID
	interfaceAPI.Operation = api.OpInterfaceSummary
	payload, err := interfaceAPI.Get()
	if err != nil {
		return integratedInterfaceSummaryResponse{}, fmt.Errorf("read interface summary for switch %q: %w", switchID, err)
	}
	var summary integratedInterfaceSummaryResponse
	if err := json.Unmarshal(payload, &summary); err != nil {
		return integratedInterfaceSummaryResponse{}, fmt.Errorf("decode interface summary for switch %q: %w", switchID, err)
	}
	return summary, nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
