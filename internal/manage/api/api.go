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

type NDManageAPI interface {
	GetLock() *sync.Mutex
	//ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error)
	GetUrl() string
	PostUrl() string
	PutUrl() string
	DeleteUrl() string
	GetDeleteQP() []string
	RscName() string
}

type NDManageAPICommon struct {
	NDManageAPI
	LockedForDeploy bool
	Client          *nd.Client
}

func (c NDManageAPICommon) Get() ([]byte, error) {
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	url := c.NDManageAPI.GetUrl()
	log.Printf("Get URL: %s\n", url)
	if c.Client == nil {
		log.Printf("************Client is nil********************")
	}
	res, err := c.Client.GetRawJson(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Finished GET: %s %v\n", c.NDManageAPI.GetUrl(), lock)
	return res, nil
}

func (c NDManageAPICommon) Post(payload []byte) (gjson.Result, error) {

	url := c.NDManageAPI.PostUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
	}
	log.Printf("Post URL: %s\n", c.NDManageAPI.PostUrl())
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	log.Printf("Post URL acquired lock: %s\n", c.NDManageAPI.PostUrl())
	var res nd.Res
	var err error
	if !json.Valid(payload) {
		res, err = c.Client.Post(c.NDManageAPI.PostUrl(), string(payload), nd.RemoveContentType)
	} else {
		res, err = c.Client.Post(c.NDManageAPI.PostUrl(), string(payload))
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDManageAPICommon) Put(payload []byte) (gjson.Result, error) {
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	res, err := c.Client.Put(c.NDManageAPI.PutUrl(), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDManageAPICommon) Delete() (gjson.Result, error) {
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	qp := c.NDManageAPI.GetDeleteQP()
	var res nd.Res
	var err error
	if qp != nil {
		res, err = c.Client.Delete(c.NDManageAPI.DeleteUrl(), "", func(req *nd.Req) {
			q := req.HttpReq.URL.Query()
			for _, s := range qp {
				keys := strings.Split(s, "=")
				q.Add(keys[0], keys[1])

			}
			req.HttpReq.URL.RawQuery = q.Encode()
		})
	} else {
		res, err = c.Client.Delete(c.NDManageAPI.DeleteUrl(), "")
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDManageAPICommon) DeleteWithPayload(payload []byte) (gjson.Result, error) {
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	res, err := c.Client.Delete(c.NDManageAPI.DeleteUrl(), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *NDManageAPICommon) SetDeployLocked() {
	c.LockedForDeploy = true
}

func (c NDManageAPICommon) DeployPost(payload []byte) (gjson.Result, error) {
	log.Printf("Deploy Post URL: %s\n", c.NDManageAPI.PostUrl())
	lock := c.NDManageAPI.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	log.Printf("Deploy Post URL acquired lock: %s\n", c.NDManageAPI.PostUrl())
	res, err := c.Client.Post(c.NDManageAPI.PostUrl(), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}
