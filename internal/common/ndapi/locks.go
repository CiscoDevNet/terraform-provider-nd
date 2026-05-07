// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndapi

import (
	"fmt"
	"log"
	"sync"
)

// DefaultFabric is used when no fabric scope is specified.
// All current resources use this value. When per-fabric locking is needed,
// pass the actual fabric name instead.
const DefaultFabric = "global"

// LockMode defines the type of hierarchical locking to acquire.
type LockMode int

const (
	// LockCRUD acquires:  Global.RLock → Fabric.RLock → Resource.Lock
	// Used by GET, POST, PUT, DELETE on a specific resource.
	LockCRUD LockMode = iota

	// LockDeploy acquires:  Global.RLock → Fabric.WLock
	// Used by config-save / config-deploy. Blocks all CRUD on that fabric.
	LockDeploy

	// LockGlobal acquires:  Global.WLock
	// Used for provider-wide exclusive operations. Blocks everything.
	LockGlobal
)

// String returns a human-readable name for the lock mode.
func (m LockMode) String() string {
	switch m {
	case LockCRUD:
		return "CRUD"
	case LockDeploy:
		return "Deploy"
	case LockGlobal:
		return "Global"
	default:
		return "Unknown"
	}
}

// ---------------------------------------------------------------------------
// Lock hierarchy:  Global (RWMutex)  →  Fabric (RWMutex)  →  Resource (Mutex)
// Locks are always acquired top-down to prevent deadlocks.
// ---------------------------------------------------------------------------

var (
	globalMu   sync.RWMutex                     // single, always exists
	fabricMu   = make(map[string]*sync.RWMutex) // lazy per fabric
	resourceMu = make(map[string]*sync.Mutex)   // lazy per fabric:resource
	muInit     sync.Mutex                       // protects map creation
)

// fabricOrDefault returns DefaultFabric when fabric is empty.
func fabricOrDefault(fabric string) string {
	if fabric == "" {
		return DefaultFabric
	}
	return fabric
}

// rscKey builds the composite map key for a (fabric, resource) pair.
func rscKey(fabric, resource string) string {
	return fabricOrDefault(fabric) + ":" + resource
}

// getFabricMu returns (or lazily creates) the per-fabric RWMutex.
func getFabricMu(fabric string) *sync.RWMutex {
	fab := fabricOrDefault(fabric)
	muInit.Lock()
	defer muInit.Unlock()
	if _, ok := fabricMu[fab]; !ok {
		fabricMu[fab] = new(sync.RWMutex)
		log.Printf("[NDAPI] Created fabric mutex: %s\n", fab)
	}
	return fabricMu[fab]
}

// getResourceMu returns (or lazily creates) the per-resource Mutex.
func getResourceMu(fabric, resource string) *sync.Mutex {
	key := rscKey(fabric, resource)
	muInit.Lock()
	defer muInit.Unlock()
	if _, ok := resourceMu[key]; !ok {
		resourceMu[key] = new(sync.Mutex)
		log.Printf("[NDAPI] Created resource mutex: %s\n", key)
	}
	return resourceMu[key]
}

// ---------------------------------------------------------------------------
// LockGuard — returned by Acquire, released via defer guard.Release()
// ---------------------------------------------------------------------------

// LockGuard holds the state needed to release acquired locks in reverse order.
type LockGuard struct {
	releases []func()
	desc     string
}

// Release unlocks all held locks in reverse acquisition order (LIFO).
func (g *LockGuard) Release() {
	for i := len(g.releases) - 1; i >= 0; i-- {
		g.releases[i]()
	}
	log.Printf("[NDAPI] Released locks: %s\n", g.desc)
	g.releases = nil
}

func (g *LockGuard) add(fn func()) {
	g.releases = append(g.releases, fn)
}

// ---------------------------------------------------------------------------
// Acquire — single entry point for hierarchical locking
// ---------------------------------------------------------------------------

// Acquire acquires the appropriate locks for the given mode.
// Always call defer guard.Release() immediately after.
//
// Lock acquisition order (always global → fabric → resource):
//
//	CRUD:    globalMu.RLock  → fabricMu[fabric].RLock  → resourceMu[fabric:resource].Lock
//	Deploy:  globalMu.RLock  → fabricMu[fabric].Lock
//	Global:  globalMu.Lock
func Acquire(fabric, resource string, mode LockMode) *LockGuard {
	fabric = fabricOrDefault(fabric)
	guard := &LockGuard{}

	switch mode {
	case LockCRUD:
		guard.desc = fmt.Sprintf("CRUD %s:%s", fabric, resource)
		log.Printf("[NDAPI] Acquiring locks: %s\n", guard.desc)

		globalMu.RLock()
		guard.add(globalMu.RUnlock)

		fm := getFabricMu(fabric)
		fm.RLock()
		guard.add(fm.RUnlock)

		rm := getResourceMu(fabric, resource)
		rm.Lock()
		guard.add(rm.Unlock)

	case LockDeploy:
		guard.desc = fmt.Sprintf("Deploy %s", fabric)
		log.Printf("[NDAPI] Acquiring locks: %s\n", guard.desc)

		globalMu.RLock()
		guard.add(globalMu.RUnlock)

		fm := getFabricMu(fabric)
		fm.Lock()
		guard.add(fm.Unlock)

	case LockGlobal:
		guard.desc = "Global"
		log.Printf("[NDAPI] Acquiring locks: %s\n", guard.desc)

		globalMu.Lock()
		guard.add(globalMu.Unlock)
	}

	log.Printf("[NDAPI] Acquired locks: %s\n", guard.desc)
	return guard
}
