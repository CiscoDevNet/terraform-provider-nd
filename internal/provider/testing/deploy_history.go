// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

// DeployHistoryDataSourceTestData contains inputs rendered by the deployment
// history data-source acceptance-test template.
type DeployHistoryDataSourceTestData struct {
	RscName     string
	FabricName  string
	NoOfRecords *int
}

// DeployHistoryDataSource creates test data for the deployment history data source.
func DeployHistoryDataSource(rscName, fabricName string, noOfRecords *int) *DeployHistoryDataSourceTestData {
	return &DeployHistoryDataSourceTestData{
		RscName:     rscName,
		FabricName:  fabricName,
		NoOfRecords: noOfRecords,
	}
}
