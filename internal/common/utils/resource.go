// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SetResourceID copies the resource's unique identifier into the computed
func SetResourceID(id *types.String, value types.String) string {
	// Always mirror the resource identifier into Terraform id, including null or unknown.
	if id != nil {
		*id = value
	}
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return value.ValueString()
}

// NotFoundOptions extends IsNotFoundError with resource-specific signals.
type NotFoundOptions struct {
	StatusCodes []int
	Messages    []string
}

// IsNotFoundError detects not-found responses. Resource-specific options take
// priority when supplied; otherwise it falls back to the default 404 checks.
//
// Default usage:
//
//	if utils.IsNotFoundError(err, responseBody) {
//		return true
//	}
//
// This matches standard not-found bodies such as:
//
//	{"code": 404, "description": "", "message": "local user not found"}
//
// Custom usage:
//
//	if utils.IsNotFoundError(err, responseBody, utils.NotFoundOptions{
//		StatusCodes: []int{404},
//		Messages:    []string{"local user not found"},
//	}) {
//		return true
//	}
//
// When custom values are supplied, only those custom values are used for matching.
func IsNotFoundError(err error, response []byte, options ...NotFoundOptions) bool {
	if err == nil && len(strings.TrimSpace(string(response))) == 0 {
		return false
	}

	errorText := ""
	if err != nil {
		errorText = err.Error()
	}

	text := strings.ToLower(strings.TrimSpace(errorText + " " + string(response)))
	compactText := strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "").Replace(text)

	hasCustomValue := false
	for _, option := range options {
		for _, message := range option.Messages {
			message = strings.ToLower(strings.TrimSpace(message))
			if message == "" {
				continue
			}
			hasCustomValue = true
			if strings.Contains(text, message) {
				return true
			}
		}

		for _, statusCode := range option.StatusCodes {
			if statusCode <= 0 {
				continue
			}
			hasCustomValue = true
			code := strconv.Itoa(statusCode)
			if strings.Contains(compactText, `"code":`+code) ||
				strings.Contains(compactText, `"status":`+code) ||
				strings.Contains(text, "status code "+code) ||
				strings.Contains(text, code+" not found") {
				return true
			}
		}
	}

	if hasCustomValue {
		return false
	}

	has404Code := strings.Contains(compactText, `"code":404`) ||
		strings.Contains(compactText, `"status":404`) ||
		strings.Contains(text, "status code 404") ||
		strings.Contains(text, "404 not found")

	return has404Code && strings.Contains(text, "not found")
}
