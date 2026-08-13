// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"os"
	"testing"
)

// StepInfo carries the per-step name and rendered HCL from build time
// (when Config is evaluated) to step-execution time (when PreConfig fires).
// This lets tests print the step header / snapshot path / TF config inline
// with the actual API call logs instead of all at once at test setup.
type StepInfo struct {
	Index int
	Name  string
	Cfg   string
}

// StepTracker ensures acceptance-test steps only run after the preceding step
// has completed successfully.
type StepTracker struct {
	t             testing.TB
	completedStep int
}

// NewStepTracker creates a sequential acceptance-test step tracker.
func NewStepTracker(t testing.TB) *StepTracker {
	t.Helper()
	return &StepTracker{t: t}
}

// RequirePrevious fails the test unless the step immediately before step has
// completed successfully.
func (s *StepTracker) RequirePrevious(step int) {
	s.t.Helper()

	expectedPreviousStep := step - 1
	if s.completedStep != expectedPreviousStep {
		s.t.Fatalf(
			"Step %d cannot run because step %d did not complete successfully; last completed step: %d",
			step,
			expectedPreviousStep,
			s.completedStep,
		)
	}
}

// Complete marks step as complete after confirming all preceding steps have
// completed successfully.
func (s *StepTracker) Complete(step int) {
	s.t.Helper()
	s.RequirePrevious(step)
	s.completedStep = step
}

// LogStep emits the per-step log block (header, snapshot path, optional HCL
// echo) at the moment the acceptance test framework is about to apply that
// step. Call it from a TestStep's PreConfig so the Go test output shows the
// step header before Terraform/provider execution for that step.
//
// Set ND_TEST_PRINT_TF=1 to also dump the rendered HCL between BEGIN/END
// markers.
func LogStep(t *testing.T, idx int, name, cfg string) {
	t.Logf("===== STEP %d: %s =====", idx, name)
	t.Logf("Snapshot: %s", SnapshotPath(name))
	if os.Getenv("ND_TEST_PRINT_TF") != "" {
		t.Logf("\n----- BEGIN TF CONFIG: %s -----\n%s\n----- END TF CONFIG: %s -----",
			name, cfg, name)
	}
}
