// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vpc_pair

import "testing"

func TestIsAcceptableVpcPairRecommendation(t *testing.T) {
	tests := map[string]struct {
		reason string
		want   bool
	}{
		"connected peers": {
			reason: vpcPairConnectedRecommendationReason,
			want:   true,
		},
		"transiently disconnected peers": {
			reason: vpcPairTransientDisconnectedRecommendation,
			want:   true,
		},
		"different roles": {
			reason: "Switches have different roles",
			want:   false,
		},
		"missing recommendation": {
			reason: "",
			want:   false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isAcceptableVpcPairRecommendation(test.reason); got != test.want {
				t.Fatalf("isAcceptableVpcPairRecommendation(%q) = %t, want %t", test.reason, got, test.want)
			}
		})
	}
}
