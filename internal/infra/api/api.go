// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/netascode/go-nd"

	"github.com/tidwall/gjson"
)

type NDInfraAPI interface {
	GetLock() *sync.Mutex
	//ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error)
	GetUrl() string
	PostUrl() string
	PutUrl() string
	DeleteUrl() string
	GetDeleteQP() []string
	RscName() string
}

type NDInfraAPICommon struct {
	NDInfraAPI
	Client *nd.Client
}

func (c NDInfraAPICommon) Get() ([]byte, error) {
	lock := c.NDInfraAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	url := c.NDInfraAPI.GetUrl()
	log.Printf("Get URL: %s\n", url)
	if c.Client == nil {
		log.Printf("************Client is nil********************")
	}
	res, err := c.Client.GetRawJson(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Finished GET: %s %v\n", c.NDInfraAPI.GetUrl(), lock)
	return res, nil
}

func (c NDInfraAPICommon) Post(payload []byte) (gjson.Result, error) {

	url := c.NDInfraAPI.PostUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
	}
	log.Printf("Post URL: %s\n", c.NDInfraAPI.PostUrl())
	lock := c.NDInfraAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	log.Printf("Post URL acquired lock: %s\n", c.NDInfraAPI.PostUrl())
	var res nd.Res
	var err error
	if !json.Valid(payload) {
		res, err = c.Client.Post(c.NDInfraAPI.PostUrl(), string(payload), nd.RemoveContentType)
	} else {
		res, err = c.Client.Post(c.NDInfraAPI.PostUrl(), string(payload))
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDInfraAPICommon) Put(payload []byte) (gjson.Result, error) {

	url := c.NDInfraAPI.PutUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
	}
	log.Printf("Put URL: %s\n", c.NDInfraAPI.PutUrl())

	lock := c.NDInfraAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	log.Printf("Put URL acquired lock: %s\n", c.NDInfraAPI.PutUrl())

	var res nd.Res
	var err error
	if !json.Valid(payload) {
		res, err = c.Client.Put(c.NDInfraAPI.PutUrl(), string(payload), nd.RemoveContentType)
	} else {
		res, err = c.Client.Put(c.NDInfraAPI.PutUrl(), string(payload))
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

func (c NDInfraAPICommon) Delete(payload ...[]byte) (gjson.Result, error) {
	lock := c.NDInfraAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}

	body := ""
	if len(payload) > 0 && payload[0] != nil {
		body = string(payload[0])
	}

	// Multi Cluster Delete API does not support query params so GetDeleteQP is not implemented for now, but keeping the logic in place in case it's needed in the future
	qp := c.NDInfraAPI.GetDeleteQP()
	var res nd.Res
	var err error
	if qp != nil && body == "" {
		res, err = c.Client.Delete(c.NDInfraAPI.DeleteUrl(), "", func(req *nd.Req) {
			q := req.HttpReq.URL.Query()
			for _, s := range qp {
				keys := strings.Split(s, "=")
				q.Add(keys[0], keys[1])
			}
			req.HttpReq.URL.RawQuery = q.Encode()
		})
	} else {
		res, err = c.Client.Delete(c.NDInfraAPI.DeleteUrl(), body)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
