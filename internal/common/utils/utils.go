// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"context"
	"errors"
	"strings"
	"time"
)

// This package contains shared utility functions used across all modules

// ErrPollTimeout is returned when PollUntil reaches its total timeout.
var ErrPollTimeout = errors.New("poll timed out")

// PollUntil runs poll immediately, then waits interval after each incomplete
// attempt before retrying. The callback receives a context bounded by the total
// timeout so an in-flight polling request can be canceled.
func PollUntil(ctx context.Context, interval time.Duration, timeout time.Duration, poll func(context.Context) (bool, error)) error {
	pollCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		done, err := poll(pollCtx)
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if errors.Is(pollCtx.Err(), context.DeadlineExceeded) {
			return ErrPollTimeout
		}
		if err != nil {
			return err
		}
		if done {
			return nil
		}

		wait := time.NewTimer(interval)
		select {
		case <-ctx.Done():
			wait.Stop()
			return ctx.Err()
		case <-pollCtx.Done():
			wait.Stop()
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return ErrPollTimeout
		case <-wait.C:
		}
	}
}

func EnabledDisabledToBool(status string) bool {
	return strings.EqualFold(status, "enabled")
}

func BoolToEnabledDisabled(enabled bool) string {
	if enabled {
		return "enabled"
	}
	return "disabled"
}
