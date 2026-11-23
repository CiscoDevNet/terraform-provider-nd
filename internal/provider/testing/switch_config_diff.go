// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

// SwitchConfigDiffDataSourceTestData contains the inputs rendered by the switch
// config diff data-source acceptance-test template.
type SwitchConfigDiffDataSourceTestData struct {
	RscName    string
	FabricName string
	SwitchID   string
}

// SwitchConfigDiffDataSource creates test data for the switch config diff data source.
func SwitchConfigDiffDataSource(rscName, fabricName, switchID string) *SwitchConfigDiffDataSourceTestData {
	return &SwitchConfigDiffDataSourceTestData{
		RscName:    rscName,
		FabricName: fabricName,
		SwitchID:   switchID,
	}
}
