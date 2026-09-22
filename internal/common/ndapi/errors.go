// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndapi

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

// ErrNotFound identifies an HTTP 404 returned by Nexus Dashboard.
var ErrNotFound = errors.New("Nexus Dashboard resource not found")

var statusCodePattern = regexp.MustCompile(`^HTTP Request failed: StatusCode ([1-5][0-9]{2})$`)

// RequestError records request context around the HTTP status code that go-nd
// currently exposes only through error text. Response is retained for callers
// that can safely include it in a diagnostic; Error does not include it.
type RequestError struct {
	StatusCode int
	Method     string
	URL        string
	Response   []byte
	Err        error
}

func (e *RequestError) Error() string {
	if e == nil || e.Err == nil {
		return "Nexus Dashboard request failed"
	}
	return e.Err.Error()
}

func (e *RequestError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func (e *RequestError) Is(target error) bool {
	return target == ErrNotFound && e != nil && e.StatusCode == http.StatusNotFound
}

// ClassifyRequestError converts the exact status-code error format emitted by
// go-nd into a typed error. Unknown error formats are returned unchanged.
func ClassifyRequestError(method, url string, response []byte, err error) error {
	if err == nil {
		return nil
	}

	var requestErr *RequestError
	if errors.As(err, &requestErr) {
		return err
	}

	matches := statusCodePattern.FindStringSubmatch(err.Error())
	if len(matches) != 2 {
		return err
	}

	statusCode, conversionErr := strconv.Atoi(matches[1])
	if conversionErr != nil {
		return err
	}

	return &RequestError{
		StatusCode: statusCode,
		Method:     method,
		URL:        url,
		Response:   append([]byte(nil), response...),
		Err:        err,
	}
}
