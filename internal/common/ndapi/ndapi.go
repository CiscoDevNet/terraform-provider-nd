// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndapi

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/netascode/go-nd"

	"github.com/tidwall/gjson"
)

type NexusDashboardAPI interface {
	GetUrl() string
	PostUrl() string
	PutUrl() string
	DeleteUrl() string
	GetDeleteQP() []string
	RscName() string
}

type NexusDashboardAPICommon struct {
	NexusDashboardAPI
	Client *nd.Client
	Fabric string // Fabric scope for locking; defaults to DefaultFabric if empty
}

// APIOptions encapsulates optional settings for API calls.
type APIOptions struct {
	DisablePayloadLog bool // suppress body logging for sensitive payloads
}

// Mods builds the request modifier list from the options and payload.
// Returns nil if the receiver is nil and the payload is valid JSON.
// Payload is given just in case there are payload based options, eg. Json here
func (o *APIOptions) Mods(payload []byte) []func(*nd.Req) {
	var mods []func(*nd.Req)
	if o != nil && o.DisablePayloadLog {
		mods = append(mods, nd.NoLogPayload)
	}
	if !json.Valid(payload) {
		mods = append(mods, nd.RemoveContentType)
	}
	return mods
}

// FabricScope returns the fabric name used for lock scoping.
// Returns DefaultFabric ("global") when Fabric is not set.
func (c NexusDashboardAPICommon) FabricScope() string {
	return fabricOrDefault(c.Fabric)
}

func (c NexusDashboardAPICommon) Get() ([]byte, error) {
	guard := Acquire(c.FabricScope(), c.NexusDashboardAPI.RscName(), LockCRUD)
	defer guard.Release()
	url := c.NexusDashboardAPI.GetUrl()
	log.Printf("Get URL: %s\n", url)
	if c.Client == nil {
		log.Printf("************Client is nil********************")
	}
	res, err := c.Client.GetRawJson(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Finished GET: %s\n", c.NexusDashboardAPI.GetUrl())
	return res, nil
}

// Post executes a CRUD POST with hierarchical locking.
func (c NexusDashboardAPICommon) Post(payload []byte, opts *APIOptions) (gjson.Result, error) {
	guard := Acquire(c.FabricScope(), c.NexusDashboardAPI.RscName(), LockCRUD)
	defer guard.Release()

	url := c.NexusDashboardAPI.PostUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
	}

	return c.Client.Post(url, string(payload), opts.Mods(payload)...)
}

// Put executes a CRUD PUT with hierarchical locking.
func (c NexusDashboardAPICommon) Put(payload []byte, opts *APIOptions) (gjson.Result, error) {
	guard := Acquire(c.FabricScope(), c.NexusDashboardAPI.RscName(), LockCRUD)
	defer guard.Release()

	url := c.NexusDashboardAPI.PutUrl()
	if strings.Contains(url, "deploy") {
		panic("Deploy URL detected in Put call. Use DeployPost method for deployments")
	}

	return c.Client.Put(url, string(payload), opts.Mods(payload)...)
}

func (c NexusDashboardAPICommon) Delete(payload ...[]byte) (gjson.Result, error) {
	guard := Acquire(c.FabricScope(), c.NexusDashboardAPI.RscName(), LockCRUD)
	defer guard.Release()

	body := ""
	if len(payload) > 0 && payload[0] != nil {
		body = string(payload[0])
	}

	qp := c.NexusDashboardAPI.GetDeleteQP()
	var res nd.Res
	var err error
	if qp != nil && body == "" {
		res, err = c.Client.Delete(c.NexusDashboardAPI.DeleteUrl(), "", func(req *nd.Req) {
			q := req.HttpReq.URL.Query()
			for _, s := range qp {
				keys := strings.Split(s, "=")
				q.Add(keys[0], keys[1])
			}
			req.HttpReq.URL.RawQuery = q.Encode()
		})
	} else {
		res, err = c.Client.Delete(c.NexusDashboardAPI.DeleteUrl(), body)
	}
	return res, err
}

// DeployPost performs a POST for deploy/save operations (no CRUD locks).
func (c NexusDashboardAPICommon) DeployPost(payload []byte, opts *APIOptions) (gjson.Result, error) {
	url := c.NexusDashboardAPI.PostUrl()
	log.Printf("Deploy Post URL: %s\n", url)

	return c.Client.Post(url, string(payload), opts.Mods(payload)...)
}

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
