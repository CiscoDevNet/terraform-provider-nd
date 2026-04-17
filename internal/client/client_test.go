// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"net/url"
	"testing"
)

var TestBaseUrls = [...]string{
	"https://nd.host.cisco",
	"https://nd.host.cisco/",
	"https://nd.host.cisco//",
	"https://nd.host.cisco/test",
	"https://nd.host.cisco/test/",
	"https://nd.host.cisco/test//",
}

func AssertFullUrl(t *testing.T, baseUrl string, method string, path string, expected string) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	ndclient := &Client{
		baseURL: url,
	}

	actual, err := ndclient.makeFullUrl(method, path)
	if actual != expected || err != nil {
		t.Errorf(`makeFullUrl("%s", "%s") = %q, %v, expected %#q`, method, path, actual, err, expected)
	}
}

func TestMakeFullUrl_Login(t *testing.T) {
	expected := "https://nd.host.cisco/login"
	paths := [...]string{
		"login",
		"/login",
		"///login",
	}
	for _, baseUrl := range TestBaseUrls {
		for _, path := range paths {
			AssertFullUrl(t, baseUrl, "GET", path, expected)
		}
	}
}

func TestMakeFullUrl_Get(t *testing.T) {
	expected_nd := "https://nd.host.cisco/nexus/api/sitemanagement/v4/sites"
	paths := [...]string{
		"nexus/api/sitemanagement/v4/sites",
		"/nexus/api/sitemanagement/v4/sites",
		"///nexus/api/sitemanagement/v4/sites",
	}
	for _, baseUrl := range TestBaseUrls {
		for _, path := range paths {
			AssertFullUrl(t, baseUrl, "GET", path, expected_nd)
		}
	}
}

func TestMakeFullUrl_Patch(t *testing.T) {
	expected := "https://nd.host.cisco/nexus/api/sitemanagement/v4/sites?validate=false"
	path := "/nexus/api/sitemanagement/v4/sites"
	for _, baseUrl := range TestBaseUrls {
		AssertFullUrl(t, baseUrl, "PATCH", path, expected)
	}
}

func TestMakeFullUrl_PatchExtraQuery(t *testing.T) {
	expected := "https://nd.host.cisco/nexus/api/sitemanagement/v4/sites?extra=query&validate=false"
	path := "nexus/api/sitemanagement/v4/sites?extra=query"
	for _, baseUrl := range TestBaseUrls {
		AssertFullUrl(t, baseUrl, "PATCH", path, expected)
	}
}
