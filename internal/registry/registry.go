// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package registry

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type ClientProvider interface {
	GetModule(name string) interface{}
}

type ResourceFactory = func() resource.Resource

type DataSourceFactory = func() datasource.DataSource

type ModuleRegistry struct {
	resources   map[string][]ResourceFactory
	dataSources map[string][]DataSourceFactory
}

var globalRegistry = &ModuleRegistry{
	resources:   make(map[string][]ResourceFactory),
	dataSources: make(map[string][]DataSourceFactory),
}

func RegisterResource(moduleName string, factory ResourceFactory) {
	globalRegistry.resources[moduleName] = append(
		globalRegistry.resources[moduleName],
		factory,
	)
}

func RegisterDataSource(moduleName string, factory DataSourceFactory) {
	globalRegistry.dataSources[moduleName] = append(
		globalRegistry.dataSources[moduleName],
		factory,
	)
}

func GetResources(moduleName string) []ResourceFactory {
	return globalRegistry.resources[moduleName]
}

func GetDataSources(moduleName string) []DataSourceFactory {
	return globalRegistry.dataSources[moduleName]
}

func GetAllResources() []ResourceFactory {
	var all []ResourceFactory
	for _, factories := range globalRegistry.resources {
		all = append(all, factories...)
	}
	return all
}

func GetAllDataSources() []DataSourceFactory {
	var all []DataSourceFactory
	for _, factories := range globalRegistry.dataSources {
		all = append(all, factories...)
	}
	return all
}
