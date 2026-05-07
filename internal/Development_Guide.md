# Terraform Provider ND Development Guide

## Code Structure Overview

This Terraform provider manages Cisco Nexus Dashboard resources using a modular architecture that supports multiple ND services while avoiding cyclic imports.

### Project Structure

```
terraform-provider-nd/
├── main.go                    # Provider entry point
├── go.mod                     # Go module definition
├── docs/                      # Documentation
└── internal/                  # Internal packages
    ├── registry/              # Shared interfaces and registration
    │   └── registry.go        # ClientProvider interface, resource registry
    ├── provider/              # Core provider logic
    │   ├── provider.go        # Main provider implementation
    │   └── client.go          # NDClient with shared API client
    └── manage/                # Manage module
        ├── manage.go          # Module client implementation
        ├── register.go        # Resource/datasource registry access
        ├── api/               # Manage API
        │   ├── api.go         # API client wrapper
        │   └── fabric_api.go  # Fabric-specific API
        └── resource_*/        # Individual resources
            ├── init.go        # Auto-registration
            └── *.go           # Resource implementation
```

### Key Architectural Patterns

#### 1. **Registry Pattern**
- Resources self-register via `init()` functions
- No direct imports between modules
- Clean separation of concerns

#### 2. **Shared API Client**
- Single `*nd.Client` instance shared across all modules
- Connection pooling and rate limiting work correctly
- Efficient resource usage

#### 3. **Eager Initialization**
- All modules created during provider configuration
- Early error detection
- Simple, predictable code path

#### 4. **Interface-Based Communication**
- `registry.ClientProvider` interface breaks cyclic dependencies
- Resources access modules through interfaces, not concrete types

### How It Works

1. **Provider Startup**: Creates shared API client and all modules
2. **Resource Registration**: Resources auto-register via `init()` functions
3. **Resource Access**: Resources get module clients through `ClientProvider` interface
4. **API Calls**: All modules use the same shared API client

### Benefits

- ✅ **No Cyclic Imports**: Registry pattern prevents import cycles
- ✅ **Shared Resources**: Single API client for efficiency
- ✅ **Modular Design**: Easy to add new modules
- ✅ **Clean Architecture**: Clear separation of concerns

---

# Adding a New Module to Terraform Provider ND

This guide explains how to add a new submodule (like `onemanage`, `insights`, etc.) to the Terraform Provider using the registry pattern with shared API client.

## Architecture Overview

The provider uses a **registry pattern with shared API client**:

1. **Registry Pattern**: Resources self-register via `init()` functions
2. **Shared API Client**: All modules share the same `*nd.Client` instance
3. **Eager Initialization**: Modules are created during provider configuration
4. **No Cyclic Imports**: Clean separation of concerns via the `registry` package

### Component Structure

```
internal/
├── registry/           # Shared interfaces (no domain imports)
│   └── registry.go     # ClientProvider interface, registration functions
├── provider/           # Provider configuration
│   ├── provider.go     # Main provider logic
│   └── client.go       # NDClient with shared API client
├── manage/             # Example module
│   ├── manage.go       # Module client
│   ├── register.go     # Returns registered resources
│   ├── api/            # Manage API
│   │   ├── api.go     # API client wrapper
│   │   └── fabric_api.go # Fabric-specific API
│   └── resource_*/     # Individual resources
│       ├── init.go     # Auto-registration
│       └── *.go        # Resource implementation
└── onemanage/          # Your new module (to be created)
```

## Step-by-Step Guide

### Step 1: Create Module Package Structure

Create the following directory structure:

```
internal/onemanage/
├── manage.go           # Module client and initialization
├── register.go         # Resource/datasource retrieval
└── resource_*/         # Your resources
    ├── init.go
    └── *.go
```

### Step 2: Implement Module Client (`onemanage/manage.go`)

```go
package onemanage

import (
	nd "github.com/netascode/go-nd"
)

const ModuleKey = "onemanage"

type OneManageClient struct {
	ApiClient *nd.Client
}

var onemanageInstance *OneManageClient

// NewClient creates the onemanage module with the shared API client
func NewClient(client *nd.Client) *OneManageClient {
	if onemanageInstance == nil {
		onemanageInstance = &OneManageClient{
			ApiClient: client,
		}
	}
	return onemanageInstance
}
```

**Key Points:**
- Module key should be unique across all modules
- Client receives the shared `*nd.Client` instance
- Singleton pattern prevents multiple instances
- No error handling needed (client already created by provider)

### Step 3: Create Registry Integration (`onemanage/register.go`)

```go
package onemanage

import (
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the onemanage module.
// Resources auto-register via init() in their packages.
func GetResources() []func() resource.Resource {
	return registry.GetResources(ModuleKey)
}

// GetDataSources returns all data sources for the onemanage module.
// Data sources auto-register via init() in their packages.
func GetDataSources() []func() datasource.DataSource {
	return registry.GetDataSources(ModuleKey)
}
```

### Step 4: Create a Resource with Auto-Registration

Example: `onemanage/resource_example/init.go`

```go
package resource_example

import (
	"terraform-provider-nd/internal/registry"
)

const ModuleKey = "onemanage"

func init() {
	registry.RegisterResource(ModuleKey, NewExampleResource)
}
```

Example: `onemanage/resource_example/example.go`

```go
package resource_example

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/onemanage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ resource.Resource              = &exampleResource{}
	_ resource.ResourceWithConfigure = &exampleResource{}
)

type exampleResource struct {
	client *onemanage.OneManageClient
}

func NewExampleResource() resource.Resource {
	return &exampleResource{}
}

func (r *exampleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T.", req.ProviderData),
		)
		return
	}

	module := client.GetModule(ModuleKey)
	if module == nil {
		resp.Diagnostics.AddError(
			"OneManage Module Not Found",
			"The onemanage module was not registered with the provider.",
		)
		return
	}

	r.client = module.(*onemanage.OneManageClient)
}

// Implement Metadata, Schema, Create, Read, Update, Delete methods...
```

### Step 5: Register Module in Provider (`provider/provider.go`)

Add your module to the module registration in `Configure()`:

```go
// Register module-specific clients (eager initialization)
// Each team adds one line here for their module
ndClient.NDModules[manage.ModuleKey] = manage.NewManage(&client)
ndClient.NDModules[onemanage.ModuleKey] = onemanage.NewClient(&client)
```

### Step 6: Register Resources/DataSources in Provider

In `provider.go`, add your module's resources and data sources:

```go
// DataSources
func (p *NexusDashboardProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	dataSources := []func() datasource.DataSource{}
	dataSources = append(dataSources, manage.GetDataSources()...)
	dataSources = append(dataSources, onemanage.GetDataSources()...)  // Add this line
	return dataSources
}

// Resources
func (p *NexusDashboardProvider) Resources(_ context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{}
	resources = append(resources, manage.GetResources()...)
	resources = append(resources, onemanage.GetResources()...)  // Add this line
	return resources
}
```

### Step 7: Import for Side-Effects (`provider/provider.go`)

Add blank import to trigger resource registration:

```go
import (
	// ... other imports ...
	
	_ "terraform-provider-nd/internal/manage/resource_fabric_vxlan"
	_ "terraform-provider-nd/internal/onemanage/resource_example"  // Add this
)
```

## Important Concepts

### 1. Shared API Client

All modules receive the **same** `*nd.Client` instance:

```go
// Provider creates one client
client, err := nd.NewClient(url, basePath, username, password, ...)

// All modules receive the same client
ndClient.NDModules[manage.ModuleKey] = manage.NewManage(&client)
ndClient.NDModules[onemanage.ModuleKey] = onemanage.NewClient(&client)
```

**Benefits:**
- Connection pooling works correctly
- Rate limiting is shared across modules
- Authentication state is shared
- More efficient resource usage

### 2. Registry Pattern

A **registry** stores resource factories:

```go
// Resource factory function
type ResourceFactory = func() resource.Resource

// Example factory
func NewExampleResource() resource.Resource {
	return &exampleResource{}  // Creates new instance
}

// Auto-registration
func init() {
	registry.RegisterResource("onemanage", NewExampleResource)
}
```

The factory function is **stored** in the registry and **called** when needed.

### 3. Eager Initialization

Modules are created **during provider configuration**:

```go
// Provider.Configure() creates all modules upfront
ndClient.NDModules[manage.ModuleKey] = manage.NewManage(&client)
ndClient.NDModules[onemanage.ModuleKey] = onemanage.NewClient(&client)
```

**Benefits:**
- Simple code path
- Early error detection
- All modules available immediately
- Predictable startup time

### 4. Avoiding Cyclic Imports

**❌ DON'T:**
```go
// In onemanage/manage.go
import "terraform-provider-nd/internal/provider"  // Creates cycle!

type OneManageClient struct {
	ProviderClient *provider.NDClient  // Bad!
}
```

**✅ DO:**
```go
// In resource file
import "terraform-provider-nd/internal/registry"

// Use interface, not concrete type
client, ok := req.ProviderData.(registry.ClientProvider)
```

## Checklist

- [ ] Created `internal/onemanage/manage.go` with `NewClient(*nd.Client)`
- [ ] Created `internal/onemanage/register.go`
- [ ] Created resource packages with `init.go`
- [ ] Implemented resources using `registry.ClientProvider`
- [ ] Added module to `provider.go` Configure method
- [ ] Added module to `DataSources()` and `Resources()` functions
- [ ] Added blank import in `provider.go`
- [ ] Tested that resources work correctly
- [ ] Verified shared client is used by all modules


Run Terraform and check that all modules log the **same pointer address**.

## Common Pitfalls

1. **Forgetting `init.go`**: Resources won't auto-register
2. **Wrong ModuleKey**: Must match between resource and module
3. **Missing blank import**: Resource `init()` never runs
4. **Importing provider package**: Creates cyclic import
5. **Creating separate clients**: Defeats shared client benefits

## Example: Complete Module

See `internal/manage/` for a complete working example of this pattern.

---

# Concurrency: Hierarchical Locking

The provider uses a **three-level hierarchical locking** system implemented in `internal/common/ndapi/` to ensure safe concurrent API calls and coordinated deploy operations. All locks are acquired top-down and released bottom-up (LIFO) to prevent deadlocks.

## Lock Hierarchy

```
Global RWMutex        (single, always exists)               ← sync.RWMutex
  └─ Fabric RWMutex   (one per fabric, lazily created)      ← sync.RWMutex
       └─ Resource Mutex (one per fabric:resource, lazily)   ← sync.Mutex
```

All three levels live as package-level variables in `locks.go`:

```go
var (
    globalMu   sync.RWMutex                     // single, always exists
    fabricMu   = make(map[string]*sync.RWMutex) // lazy per fabric
    resourceMu = make(map[string]*sync.Mutex)   // lazy per fabric:resource
    muInit     sync.Mutex                       // protects map creation
)
```

### Level 1 — Global RWMutex

A single provider-wide lock. CRUD and deploy operations acquire it as a **read lock**; only truly global exclusive operations (e.g. provider shutdown) would need a **write lock**. Currently no operations use `LockGlobal`.

### Level 2 — Per-Fabric RWMutex

One per fabric, keyed by fabric name. When no fabric is specified the constant `DefaultFabric = "global"` is used. CRUD operations acquire it as a **read lock** (multiple CRUD ops on the same fabric run in parallel). Deploy operations acquire it as a **write lock** (blocks all CRUD on that fabric, waits for in-flight CRUD to finish).

### Level 3 — Per-Resource Mutex

One per (fabric, resource) pair — e.g. `"global:fabric"`, `"global:inventory_switch"`. Serialises all API calls for the same resource type within a fabric. Different resource types on the same fabric run in parallel.

## Lock Modes

| Mode | Global | Fabric | Resource | Used by |
|---|---|---|---|---|
| **`LockCRUD`** | `RLock` | `RLock` | `Lock` | `Get`, `Post`, `Put`, `Delete` on a specific resource |
| **`LockDeploy`** | `RLock` | **`WLock`** | — | `ConfigSaveAndDeploy` (config-save + config-deploy) |
| **`LockGlobal`** | **`WLock`** | — | — | Reserved for provider-wide exclusive operations |

## The `Acquire` / `LockGuard` API

All locking is done through a single function — `ndapi.Acquire(fabric, resource, mode)` — which returns a `*LockGuard`. The guard is released via `defer guard.Release()`.

```go
guard := ndapi.Acquire(fabricScope, resourceName, ndapi.LockCRUD)
defer guard.Release()
```

- Locks are acquired top-down (global → fabric → resource).
- `Release()` unlocks in LIFO order (resource → fabric → global).
- Mutexes are **lazily created** on first access — no upfront registration in `provider.go` is needed.
- Detailed acquire/release logging is built in (`[NDAPI] Acquiring locks: CRUD global:inventory_switch`).

## How CRUD Methods Use Locking

The standard HTTP methods in `NexusDashboardAPICommon` (`Get`, `Post`, `Put`, `Delete`) **automatically** acquire `LockCRUD` using the API instance's `FabricScope()` and `RscName()`. Callers never need to lock manually:

```go
// From ndapi.go — Post acquires locks automatically
func (c NexusDashboardAPICommon) Post(payload []byte, disablePayloadLog ...bool) (gjson.Result, error) {
    guard := Acquire(c.FabricScope(), c.NexusDashboardAPI.RscName(), LockCRUD)
    defer guard.Release()

    url := c.NexusDashboardAPI.PostUrl()
    // Panic guard: deploy URLs must use DeployPost, not Post
    if strings.Contains(url, "deploy") {
        panic("Deploy URL detected in Post call. Use DeployPost method for deployments")
    }
    // ... HTTP call ...
}
```

`DeployPost` is the exception — it acquires **no locks** because it is always called within a deploy context where `LockDeploy` is already held by the caller.

## Adding a New Resource — Step by Step

### Step 1: Define a Resource Name Constant

Define a resource name constant **in your own API file** (not in `common/ndapi/`). The common `ndapi` package must stay resource-agnostic — resource developers should never need to edit it.

```go
// manage/api/my_new_resource_api.go
const RscNameMyNewResource = "my_new_resource"
```

Existing examples:
- `manage/api/inventory_api.go` → `const RscNameInventorySwitch = "inventory_switch"`
- `manage/api/fabric_api.go` → `const RscNameFabric = "fabric"`
- `manage/api/config_api.go` → `const RscNameConfig = "config"`
- `infra/api/multi_cluster_connectivity_api.go` → `const RscNameMultiClusterConnectivity = "multi_cluster_connectivity"`

No registration call or edit to `locks.go` is needed — the mutex is lazily created on first `Acquire`.

### Step 2: Create the API Struct

Embed `ndapi.NexusDashboardAPICommon` and implement the `NexusDashboardAPI` interface.

**Interface contract** (from `ndapi.go`):
```go
type NexusDashboardAPI interface {
    GetUrl() string
    PostUrl() string
    PutUrl() string
    DeleteUrl() string
    GetDeleteQP() []string
    RscName() string
}
```

**Real example — `InventoryAPI`** (from `manage/api/inventory_api.go`):
```go
package api

import (
    "fmt"
    "terraform-provider-nd/internal/common/ndapi"
    nd "github.com/netascode/go-nd"
)

const RscNameInventorySwitch = "inventory_switch"

type InventoryAPI struct {
    ndapi.NexusDashboardAPICommon           // embeds Client, Fabric, and all HTTP methods
    FabricName   string
    SerialNumber string
    Operation    InventoryOperation          // enum for URL dispatch
}

func NewInventoryAPI(client *nd.Client, fabric string) *InventoryAPI {
    papi := new(InventoryAPI)
    papi.Client = client
    papi.Fabric = fabric
    papi.NexusDashboardAPI = papi            // self-reference for interface dispatch
    return papi
}

func (c *InventoryAPI) RscName() string      { return RscNameInventorySwitch }
func (c *InventoryAPI) GetDeleteQP() []string { return nil }

// PostUrl dispatches to the correct endpoint based on the current operation:
func (c *InventoryAPI) PostUrl() string {
    switch c.Operation {
    case OpAddSwitches:
        return fmt.Sprintf("/manage/fabrics/%s/switches", c.FabricName)
    case OpShallowDiscovery:
        return fmt.Sprintf("/manage/fabrics/%s/actions/shallowDiscovery", c.FabricName)
    case OpBootstrap:
        return fmt.Sprintf("/manage/fabrics/%s/inventory/poap", c.FabricName)
    // ... other operations
    }
}
```

**Key constructor points:**
- Takes `(client *nd.Client, fabric string)` — `Fabric` is used by `FabricScope()` for lock scoping.
- `papi.NexusDashboardAPI = papi` is **mandatory** — enables self-referencing interface dispatch that makes `Post()`, `Get()`, etc. call your `PostUrl()`, `GetUrl()`, etc.
- No `mutex` field or `GetLock()` method — locking is handled internally by `Acquire`.

### Step 3: The `SetOperation` Pattern

When a single API struct handles multiple endpoints, use an operation enum + setter:

```go
type InventoryOperation int

const (
    OpGetSwitch InventoryOperation = iota
    OpGetAllSwitches
    OpAddSwitches
    OpRemoveSwitches
    OpShallowDiscovery
    OpBootstrap
    // ...
)

func (c *InventoryAPI) SetOperation(op InventoryOperation) *InventoryAPI {
    c.Operation = op
    return c
}
```

Usage in resource code:
```go
invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
invAPI.FabricName = fabricName

// Discover
invAPI.SetOperation(api.OpShallowDiscovery)
resp, err := invAPI.Post(discoveryPayload, true)  // true = suppress payload logging

// Add switches
invAPI.SetOperation(api.OpAddSwitches)
resp, err = invAPI.Post(addPayload)

// Remove switches
invAPI.SetOperation(api.OpRemoveSwitches)
resp, err = invAPI.Post(removePayload)
```

Each `Post()` call automatically acquires `LockCRUD` for `inventory_switch` — no manual locking needed.

### Step 4: Use in Resource CRUD Methods

Locking is **automatic** — just create the API instance and call methods:

```go
func (r *myResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    myApi := api.NewMyResourceAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
    // Post() automatically acquires Global.RLock → Fabric.RLock → Resource.Lock
    result, err := myApi.Post(payload)
    // ...
}
```

You do **not** need to manually call `Acquire` in resource CRUD methods — the API `Get`/`Post`/`Put`/`Delete` methods handle this.

### Step 5: Suppress Payload Logging for Sensitive Data

`Post()` and `Put()` accept an optional `disablePayloadLog` flag:

```go
// Credentials payload — suppress logging
resp, err := invAPI.Post(credentialPayload, true)
```

Use this whenever the payload contains passwords or other sensitive data.

## Deploy Operations — Critical Section

**Deploy operations require special handling.** A deploy acquires `LockDeploy` (fabric write lock) to prevent any concurrent CRUD operations from running on that fabric during the deploy.

### Why This Matters

Without the fabric write lock:
- A Create could POST new config while a deploy is in progress, causing the deploy to operate on stale or inconsistent state.
- A Delete could remove config that the deploy expects to exist.
- Race conditions between CRUD and deploy lead to non-deterministic failures on the NDFC appliance.

### Reference Implementation

From `internal/manage/deployment/deployment.go`:

```go
func ConfigSaveAndDeploy(ctx context.Context, client *nd.Client, fabricName string,
    recalculate bool, deploy bool, dg *diag.Diagnostics) {
    // Acquire fabric write lock — blocks all CRUD on this fabric
    guard := ndapi.Acquire(ndapi.DefaultFabric, "", ndapi.LockDeploy)
    defer guard.Release()

    if recalculate {
        respMsg, err := configSave(ctx, client, fabricName, dg)
        if err != nil {
            dg.AddError("Error Saving Config", fmt.Sprintf("%v: %s", err, respMsg))
            return
        }
    }
    if deploy {
        respMsg, err := configDeploy(ctx, client, fabricName)
        if err != nil {
            dg.AddError("Error Deploying Config", fmt.Sprintf("%v: %s", err, respMsg))
            return
        }
    }
}
```

Inside `configSave` and `configDeploy`:
```go
func configSave(ctx context.Context, client *nd.Client, fabricName string, dg *diag.Diagnostics) (string, error) {
    configAPI := api.NewConfigAPI(client, ndapi.DefaultFabric)
    configAPI.FabricName = fabricName
    configAPI.SetOperation(api.OpConfigSave)
    resp, err := configAPI.DeployPost(nil)   // DeployPost — no CRUD locks
    // ...
}
```

### Rules for Deploy

1. **Acquire `LockDeploy` before calling `DeployPost()`** — `DeployPost()` intentionally acquires no locks; the caller controls the scope via `Acquire`.
2. **Always `defer guard.Release()`** — ensures locks are released even if the deploy fails.
3. **Never call `Post()` for deploy URLs** — `Post()` has a **panic guard** that detects deploy URLs and crashes immediately. Always use `DeployPost()` instead.
4. **Use `DeployPost()` for all API calls within a deploy context** — calling `Post()`, `Put()`, etc. inside a `LockDeploy` scope will deadlock because those methods attempt to acquire `LockCRUD` (which takes a Fabric RLock, but the deploy already holds the Fabric WLock).

### Deadlock Prevention

The key invariant: **never acquire a `LockCRUD` while holding `LockDeploy`** on the same fabric. The `LockCRUD` path takes `Fabric.RLock`, which will deadlock against the `Fabric.WLock` held by `LockDeploy`. This is why `configSave` and `configDeploy` use `DeployPost()` instead of `Post()`.

## Custom API Methods

If your resource needs a custom API method beyond `Get`/`Post`/`Put`/`Delete`, use `Acquire` directly.

**Real example — `PostDelete` for multi-cluster connectivity** (from `infra/api/multi_cluster_connectivity_api.go`):

```go
func (c *ClusterAPI) PostDelete(payload []byte) (nd.Res, error) {
    guard := ndapi.Acquire(c.FabricScope(), c.RscName(), ndapi.LockCRUD)
    defer guard.Release()

    res, err := c.Client.Post(
        fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName), string(payload),
    )
    return res, err
}
```

This uses a POST to a `/remove` endpoint for deletion — the standard `Delete()` method doesn't fit, so manual `Acquire` is required.

## FSM and Long-Running Operations

The `InventoryFSM` (in `inventory_fsm.go`) drives multi-step workflows like discovery → add → wait → credentials → roles → config-save → deploy. Each step calls the API methods which acquire and release locks independently per call.

**Key principle:** Locks are held only for the duration of a single API call, not across the entire FSM lifecycle. This allows other resources to interleave their operations between FSM steps.

The FSM handles deploy via `ConfigSaveAndDeploy`, which acquires `LockDeploy` only for the config-save + deploy steps — not for the entire workflow.

## Source Files

| File | Contents |
|---|---|
| `internal/common/ndapi/locks.go` | Lock hierarchy (`globalMu`, `fabricMu`, `resourceMu`), `LockMode`, `LockGuard`, `Acquire()`, `DefaultFabric` |
| `internal/common/ndapi/ndapi.go` | `NexusDashboardAPI` interface, `NexusDashboardAPICommon` with `Get`/`Post`/`Put`/`Delete`/`DeployPost`, `FabricScope()` |
| `internal/manage/deployment/deployment.go` | Reference deploy implementation using `Acquire(fabric, "", LockDeploy)` |
| `internal/manage/api/inventory_api.go` | Real example: `InventoryAPI` with `SetOperation` pattern |
| `internal/manage/api/config_api.go` | `ConfigAPI` used by deploy (`DeployPost`) |
| `internal/manage/api/fabric_api.go` | `FabricAPI` — simple single-endpoint example |
| `internal/infra/api/multi_cluster_connectivity_api.go` | `ClusterAPI` with custom `PostDelete` using manual `Acquire` |

## Checklist for New Resources

- [ ] Defined resource name constant in your own API file (never in `common/ndapi/`)
- [ ] API struct embeds `ndapi.NexusDashboardAPICommon`
- [ ] Constructor takes `(client *nd.Client, fabric string)` and sets `papi.Fabric = fabric`
- [ ] `papi.NexusDashboardAPI = papi` (self-reference for interface dispatch)
- [ ] `RscName()` returns the resource name constant
- [ ] All interface methods implemented: `GetUrl`, `PostUrl`, `PutUrl`, `DeleteUrl`, `GetDeleteQP`
- [ ] Multi-endpoint API uses `SetOperation` pattern (not multiple API structs)
- [ ] Sensitive payloads use `Post(payload, true)` to suppress logging
- [ ] Any custom mutating methods use `ndapi.Acquire(..., ndapi.LockCRUD)`
- [ ] Deploy operations use `ndapi.Acquire(..., ndapi.LockDeploy)` + `DeployPost()`
- [ ] No `Post()`/`Put()`/`Delete()` calls inside a `LockDeploy` scope (deadlock risk)
- [ ] Tested concurrent operations to verify no deadlocks or races
