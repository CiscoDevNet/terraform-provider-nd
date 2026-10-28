# ND Policy Terraform Provider Design

Source: [ND Policy terraform provider design](https://ciscoteams.atlassian.net/wiki/spaces/NDTF/pages/1936294012/ND+Policy+terraform+provider+design) (Confluence NDTF, page 1936294012). API contracts aligned to Nexus Dashboard Manage v1, ND 4.2+ ([Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/get-policies/)).

Base URL: `https://{cluster}/api/v1/manage`

## Purpose and Scope

This design defines how Terraform creates, reads, updates, deploys, replaces, and removes **switch-level policies** and **policy groups** through the Nexus Dashboard Manage API. Terraform exposes **one** resource, `nd_policy`, with `is_policy_group` selecting the policy vs policy-group API path. It covers:

- Unified Terraform schema (NDFC-style `nd_policy`) mapped to Manage v1 payloads
- Request and response payloads from the public Policies APIs
- Template validation
- Lifecycle boundaries and identifiers
- Deletion semantics and device cleanup
- Operational safeguards for mixed IOS-XE and NX-OS fabrics

**Out of initial implementation scope**

- [List Policies History Data](https://developer.cisco.com/docs/nexus-dashboard/latest/list-policies-history-data/) (`GET /fabrics/{fabricName}/policyHistory`)
- [Get Policy Summary](https://developer.cisco.com/docs/nexus-dashboard/latest/get-policy-summary/)
- Legacy LAN Fabric Controller routes such as [Update PG Without Status](https://developer.cisco.com/docs/nexus-dashboard/latest/update-pg-without-status/) (`PUT /appcenter/cisco/ndfc/api/v1/lan-fabric/rest/control/policies/policygroup/{policyId}`) and [Mark Delete Switch Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/mark-delete-switch-policies/)

Those APIs can become later data sources or remain unused. Do not mix legacy `/lan-fabric/rest/...` paths with Manage `/api/v1/manage` resources.

## Unified Terraform Resource

Nexus Dashboard exposes two controller objects. Terraform keeps a **single** `nd_policy` resource, matching the NDFC provider, and uses `is_policy_group` as a payload-hidden discriminator.

| `is_policy_group` | Controller object | Stable ID | Membership | API family |
| --- | --- | --- | --- | --- |
| `false` (default) | Single policy | `policyId` (`POLICY-<n>`) | Exactly one serial in `serial_numbers` | `/fabrics/{fabricName}/policies` |
| `true` | Policy group | `policyId` (`POLICY-GROUP-<n>`), used as path `{policyGroupId}` | One or more serials in `serial_numbers` | `/fabrics/{fabricName}/policyGroups` |

**Lifecycle boundary:** persist the create-response `policyId`. Do not match later updates or deletes only by template name and switch serials. Group path operations use that same `policyId` value as `{policyGroupId}`; GET/list JSON still returns the ID in `policyId`.

**Scope rule:** identify switches with the controller-recognized serial. The codec maps `serial_numbers` to `switchId` (policy) or `switchIds` (group).

Changing `is_policy_group` requires replace. `serial_numbers` requires replace only when `is_policy_group` is `false` (`tf_requires_replace_if_key/value`). Group membership is an in-place PUT.

## Shared Attributes

Write schemas (`createPolicy`, `policyPut`, `createPolicyGroup`, `putPolicyGroup`) and read schemas (`policyGet`, `policyGroupsGet`) are not identical. Terraform must send the write shape and store extra GET-only fields as computed.

| Attribute | Type | Create/Update | Read | Notes |
| --- | --- | --- | --- | --- |
| **templateName** | string, max 255 | Required | Required | Treat as immutable. PUT requires the current name; changing it is a replace. |
| **entityType** | enum | Required | Present | Allowed values: `switch`, `configProfile`, `interface`. |
| **entityName** | string, max 255 | Required | Present | Use `SWITCH` when `entityType` is `switch`. Use an interface name such as `Ethernet1/1` when `entityType` is `interface`. |
| **switchId** | string | Required on policy | Present on policy | Serial number. Policy only. Terraform `serial_numbers[0]`. |
| **switchIds** | string array | Required on group | Present on group | Group only. Terraform `serial_numbers`. Updatable in place. |
| **templateInputs** | object, additionalProperties | Optional | Required on GET | Name/value map passed to the template. Terraform `policy_parameters`. |
| **description** | string, max 255 | Optional | Nullable | |
| **priority** | integer | Optional, 1–2000 | Default 500, −2000–2000 | Negative values are reserved for mark-deleted rows. Writes must stay in 1–2000. |
| **source** | string, max 255 | Optional | Default `""` | Empty source means any source can update the policy. Examples: `UNDERLAY`, `OVERLAY`, `LINK`. |
| **secondaryEntityName** | string | Optional | Optional | Overlay name when `entityType` is `configProfile`. |
| **secondaryEntityType** | same enum as `entityType` | Optional | Optional | |
| **policyId** | string | Assigned on create | Required | Controller ID. Hidden from create payload. |
| **generatedConfig** | string | — | Computed | CLI produced from the template and inputs. |
| **markDeleted** | boolean | — | Computed, default false | True after mark-delete. |
| **createTimestamp** | date-time | — | Computed | |
| **updateTimestamp** | date-time | — | Computed | |

Query parameters on mutating policy/group calls:

- `ticketId` (optional): change-control ticket. Pattern `^[a-zA-Z][a-zA-Z0-9_-]+$`, length 1–64.
- `clusterName` (optional): target cluster in a multi-cluster deployment. The provider already selects the cluster; do not expose this on the resource.

List APIs also accept `offset`, `max`, `sort`, and Lucene `filter`.

## Manage Policies API Catalog

All paths are relative to `/api/v1/manage`. Bulk action APIs return **207** with a per-item `status` (`success` / `failed`). Single update and single delete return **204**.

### Single policy

| Operation | Method and path | Request body | Success |
| --- | --- | --- | --- |
| [Get Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/get-policies/) | `GET /fabrics/{fabricName}/policies` | — | `{ "meta": ..., "policies": [ policyGet ] }` |
| [Create Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/create-policies/) | `POST /fabrics/{fabricName}/policies` | `{ "policies": [ createPolicy ] }` | 207 `{ "policies": [ { policyId, status, message, switchId, ... } ] }` |
| [Get Policy](https://developer.cisco.com/docs/nexus-dashboard/latest/get-policy/) | `GET /fabrics/{fabricName}/policies/{policyId}` | — | `policyGet` |
| [Update Policy](https://developer.cisco.com/docs/nexus-dashboard/latest/update-policy/) | `PUT /fabrics/{fabricName}/policies/{policyId}` | `policyPut` (not wrapped) | 204 |
| [Delete Policy](https://developer.cisco.com/docs/nexus-dashboard/latest/delete-policy/) | `DELETE /fabrics/{fabricName}/policies/{policyId}` | — | 204 |
| [Execute Mark Delete Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/execute-mark-delete-policies/) | `POST /fabrics/{fabricName}/policyActions/markDelete` | `{ "policyIds": [ "POLICY-..." ] }` | 207 |
| [Create Push Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/create-push-policies/) | `POST /fabrics/{fabricName}/policyActions/pushConfig` | `{ "policyIds": [ "POLICY-..." ] }` | 207 |
| [Delete Policies For Ids](https://developer.cisco.com/docs/nexus-dashboard/latest/delete-policies-for-ids/) | `POST /fabrics/{fabricName}/policyActions/remove` | `{ "policyIds": [ "POLICY-..." ] }` | 207 |

`policyActions/remove` is the bulk-delete counterpart of group `policyGroups/actions/remove`. The Terraform destroy path uses single `DELETE /policies/{policyId}` first, then mark-delete plus switch deploy for device cleanup. Do not call bulk remove from a single-resource lifecycle unless falling back after a failed single delete.

### Policy group

| Operation | Method and path | Request body | Success |
| --- | --- | --- | --- |
| [List Group Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/list-group-policies/) | `GET /fabrics/{fabricName}/policyGroups` | — | `{ "meta": ..., "policyGroups": [ policyGroupsGet ] }` |
| [Create Policy Groups For Switches](https://developer.cisco.com/docs/nexus-dashboard/latest/create-policy-groups-for-switches/) | `POST /fabrics/{fabricName}/policyGroups` | `{ "policyGroups": [ createPolicyGroup ] }` | 207 `{ "policyGroups": [ { policyId, status, switchIds, ... } ] }` |
| [List Switches For Policy Group](https://developer.cisco.com/docs/nexus-dashboard/latest/list-switches-for-policy-group/) | `GET /fabrics/{fabricName}/policyGroups/{policyGroupId}` | — | `policyGroupsGet` |
| [Update Policy Group](https://developer.cisco.com/docs/nexus-dashboard/latest/update-policy-group/) | `PUT /fabrics/{fabricName}/policyGroups/{policyGroupId}` | `putPolicyGroup` (not wrapped) | 204 |
| [Delete Group Policy](https://developer.cisco.com/docs/nexus-dashboard/latest/delete-group-policy/) | `DELETE /fabrics/{fabricName}/policyGroups/{policyGroupId}` | — | 204 |
| [Execute Mark Delete Group Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/execute-mark-delete-group-policies/) | `POST /fabrics/{fabricName}/policyGroups/actions/markDelete` | `{ "policyIds": [ "POLICY-GROUP-..." ] }` | 207 |
| [Delete Policy Groups](https://developer.cisco.com/docs/nexus-dashboard/latest/delete-policy-groups/) | `POST /fabrics/{fabricName}/policyGroups/actions/remove` | `{ "policyIds": [ "POLICY-GROUP-..." ] }` | 207 |

There is no group-specific `pushConfig`. After creating or updating a group, deploy member switches (see Helper APIs).

### Helper APIs used by the lifecycle

| Purpose | Method and path |
| --- | --- |
| Resolve management IP to serial | `GET /fabrics/{fabricName}/switches` ([List Fabric Switches](https://developer.cisco.com/docs/nexus-dashboard/latest/list-fabric-switches/)). Match `fabricManagementIp` to `switchId`. |
| Read template schema and platforms | `GET /configTemplates/{templateName}/parameters` ([Get Configuration Template Parameters](https://developer.cisco.com/docs/nexus-dashboard/latest/get-configuration-template-parameters/)) |
| Preview pending switch config | `POST /fabrics/{fabricName}/switchActions/preview` (existing provider `UrlSwitchPreview`; optional `recompute=true`) |
| Deploy pending switch config / cleanup | `POST /fabrics/{fabricName}/switchActions/deploy` (existing provider `UrlSwitchDeploy`) |

DevNet also documents [Execute Switch Action Config Deploy](https://developer.cisco.com/docs/nexus-dashboard/latest/execute-switch-action-config-deploy/) as `POST /inventory/switchActions/deploy` with `{ "switchIds": [...] }`. Policy resources must use the same fabric-scoped switch deploy/preview paths already used by `nd_config_deploy` so cleanup stays consistent with this provider.

## Request Payloads

### Create policy (`createPolicy`)

Required: `templateName`, `entityType`, `entityName`, `switchId`.

```json
{
  "policies": [
    {
      "switchId": "FDO25031SY4",
      "entityType": "switch",
      "entityName": "SWITCH",
      "templateName": "feature_enable",
      "templateInputs": {
        "featureName": "lacp"
      },
      "description": "Enable LACP",
      "priority": 500,
      "source": ""
    }
  ]
}
```

Create response item (207):

```json
{
  "policies": [
    {
      "entityName": "SWITCH",
      "entityType": "switch",
      "message": "Action is successful",
      "policyId": "POLICY-121110",
      "status": "success",
      "switchId": "FDO23240AK6",
      "templateName": "feature_enable"
    }
  ]
}
```

Treat any non-`success` status as apply failure. Then `GET` the returned `policyId` to populate computed fields.

### Update policy (`policyPut`)

PUT body is a single object, not `{ "policies": [...] }`. Required: `templateName`, `entityType`, `entityName`, `switchId`. Mutable fields called out by the API: `templateInputs`, `priority`, and similar non-identity fields.

```json
{
  "description": "Update policy to change the priority value",
  "entityName": "SWITCH",
  "entityType": "switch",
  "priority": 50,
  "source": "",
  "switchId": "FDO23240AK6",
  "templateInputs": {
    "featureName": "lacp"
  },
  "templateName": "feature_enable"
}
```

Update returns 204. Refresh state with `GET /policies/{policyId}`.

### Create policy group (`createPolicyGroup`)

Required: `templateName`, `entityType`, `entityName`, `switchIds`.

```json
{
  "policyGroups": [
    {
      "description": "Enable LACP on leaf switches",
      "entityName": "SWITCH",
      "entityType": "switch",
      "priority": 500,
      "source": "",
      "switchIds": ["FDO25031SY4", "FDO245206N5"],
      "templateInputs": {
        "featureName": "lacp"
      },
      "templateName": "feature_enable"
    }
  ]
}
```

Create response item (207):

```json
{
  "policyGroups": [
    {
      "entityName": "SWITCH",
      "entityType": "switch",
      "message": "Policy group is successfully created",
      "policyId": "POLICY-GROUP-143310",
      "status": "success",
      "switchIds": ["FDO25031SY4", "FDO245206N5"],
      "templateName": "feature_enable"
    }
  ]
}
```

### Update policy group (`putPolicyGroup`)

PUT body is a single object. Required: `templateName`, `entityType`, `entityName`, `switchIds`. `switchIds` is the membership list to apply, so membership changes are in-place updates.

```json
{
  "description": "Update priority value for group policy",
  "entityName": "SWITCH",
  "entityType": "switch",
  "priority": 50,
  "switchIds": ["FDO245206N5", "FDO245206N9"],
  "templateInputs": {
    "featureName": "lacp"
  },
  "templateName": "feature_enable"
}
```

### Mark-delete / push / bulk-remove

Shared ID list wrapper. Use `POLICY-*` IDs for policy actions and `POLICY-GROUP-*` IDs for group actions.

```json
{
  "policyIds": ["POLICY-121110"]
}
```

```json
{
  "policyIds": ["POLICY-GROUP-143310"]
}
```

## Terraform Model

One `nd_policy` resource manages one controller object. The provider wraps that object in the bulk create array and fails if that item’s 207 `status` is not `success`. `is_policy_group` is Terraform-only (`payload_hide`) and selects the API family.

```hcl
resource "nd_policy" "enable_lacp" {
  serial_numbers   = ["FDO25031SY4"]
  entity_type      = "switch"
  entity_name      = "SWITCH"
  template_name    = "feature_enable"
  is_policy_group  = false

  policy_parameters = {
    featureName = "lacp"
  }

  description = "Enable LACP"
  priority    = 500
  source      = ""
  deploy      = true
}

resource "nd_policy" "enable_lacp_leaves" {
  serial_numbers  = ["FDO25031SY4", "FDO245206N5"]
  entity_type     = "switch"
  entity_name     = "SWITCH"
  template_name   = "feature_enable"
  is_policy_group = true

  policy_parameters = {
    featureName = "lacp"
  }

  description = "Enable LACP on leaf switches"
  priority    = 500
  deploy      = true
}
```

`deploy` is Terraform-only. It is not an ND policy field.

- `is_policy_group = false` and `deploy = true`: after create/update, `POST .../policyActions/pushConfig` with the resource `policyId`.
- `is_policy_group = true` and `deploy = true`: after create/update, deploy member serials through the existing fabric switch-deploy path.
- Destroy always needs device cleanup if mark-delete left pending config; see Deletion.

### Terraform attribute to Manage API mapping

| Terraform | Generator `model_name` | Manage API field | Notes |
| --- | --- | --- | --- |
| `policy_id` | `policyId` | `policyId` | Controller-assigned `POLICY-*` or `POLICY-GROUP-*`. Terraform identity for this resource. |
| `fabric_name` | `fabricName` | path `{fabricName}` | Computed. Resolved from `serial_numbers` via `GET /manage/inventory/switches` (same pattern as `nd_vpc_pair`). |
| `is_policy_group` | `policyGroup` | — | `payload_hide`. Chooses policy vs policy-group endpoints. Requires replace. |
| `deploy` | `deploy` | — | `payload_hide`. |
| `serial_numbers` | `serialNumber` | `switchId` or `switchIds` | List in Terraform. Policy writes `switchId` from the single list element. Group writes `switchIds`. Do **not** use NDFC `ndfc_type: csv`; Manage v1 is JSON. Length must be 1 when `is_policy_group` is false. |
| `policy_parameters` | `templateInputs` | `templateInputs` | NDFC used `nvPairs`. |
| `entity_type` | `entityType` | `entityType` | Enum `switch`, `configProfile`, `interface`. Not `SWITCH`. |
| `entity_name` | `entityName` | `entityName` | Use `SWITCH` for a switch-level object. |
| `template_name` | `templateName` | `templateName` | Requires replace. |
| `description` | `description` | `description` | Optional, max 255. |
| `source` | `source` | `source` | `include_empty`. |
| `priority` | `priority` | `priority` | Default 500. Writes 1–2000. |
| `secondary_entity_name` | `secondaryEntityName` | `secondaryEntityName` | Optional; not in the NDFC def. |
| `secondary_entity_type` | `secondaryEntityType` | `secondaryEntityType` | Optional. |
| `generated_config` | `generatedConfig` | `generatedConfig` | Computed. |
| `deleted` | `markDeleted` | `markDeleted` | Computed, `tf_hide`. NDFC name `deleted`. |
| `auto_generated` | — | — | NDFC-only. Not in Manage `policyGet` / `policyGroupsGet`. Omit unless a later API revision returns it. |

Codec rules:

- Create policy: wrap one object in `{ "policies": [ ... ] }` with `switchId`.
- Create group: wrap one object in `{ "policyGroups": [ ... ] }` with `switchIds`.
- PUT is an unwrapped object on the matching path.
- On GET, copy user-set `templateInputs` keys into `policy_parameters`. Drop controller-injected keys (`FABRIC_NAME`, `SERIAL_NUMBER`, `POLICY_ID`) and other leftovers; do not expose them as a Terraform attribute.
- `rest_endpoint` in the generator is the policy collection. Custom resource/codec logic switches to `/policyGroups` when `is_policy_group` is true.

### Generator definition sketch

```yaml
resource:
  name: policy
  rest_endpoint: /manage/fabrics/{fabricName}/policies
  doc_category: Fabric
  generate_tf_resource: true
  skip_equality: false
  category: manage
  description: "Resource to configure policies on a switch"
  attributes:
    - model_name: policyId
      tf_name: policy_id
      type: String
      description: "Policy ID"
      computed: true
      use_state: true
      example: POLICY-26910

    - model_name: fabricName
      tf_name: fabric_name
      type: String
      description: Name of the fabric resolved from serial_numbers in the switch inventory
      computed: true
      payload_hide: true
      example: my_fabric

    - model_name: policyGroup
      tf_name: is_policy_group
      type: Bool
      description: "Policy Group"
      default_value: false
      payload_hide: true
      example: false
      ndfc_type: bool
      tf_requires_replace: true

    - model_name: deploy
      tf_name: deploy
      type: Bool
      description: "Deploy policy"
      default_value: false
      ndfc_type: bool
      payload_hide: true
      example: true

    - model_name: entityName
      tf_name: entity_name
      type: String
      description: "Name of the entity. Use SWITCH when entity_type is switch"
      mandatory: true
      example: "SWITCH"

    - model_name: entityType
      tf_name: entity_type
      type: String
      description: "Type of the entity"
      mandatory: true
      validator: 'OneOf("switch", "configProfile", "interface")'
      example: "switch"

    - model_name: description
      tf_name: description
      type: String
      description: "Description of the policy"
      optional: true
      example: "Policy for switch"

    - model_name: templateName
      tf_name: template_name
      type: String
      description: "Name of the template"
      mandatory: true
      tf_requires_replace: true
      example: "TelemetryDst_EF"

    - model_name: source
      tf_name: source
      type: String
      description: "Source of the policy"
      optional: true
      include_empty: true
      example: "CLI"

    - model_name: priority
      tf_name: priority
      type: Int64
      description: "Priority of the policy"
      optional: true
      default_value: 500
      validator: 'int64validator.Between(1, 2000)'
      example: 500

    - model_name: serialNumber
      tf_name: serial_numbers
      type: List:String
      mandatory: true
      description: "Serial numbers of the device. Use only one device for policy. Multiple devices can be in the list if this is a policy group"
      example: '["FDO245206N5", "FDO245206N6"]'
      tf_requires_replace_if_key: is_policy_group
      tf_requires_replace_if_value: false
      tf_requires_replace_if_type: bool

    - model_name: templateInputs
      tf_name: policy_parameters
      type: "Map:String"
      description: "List of name value pairs passed to the template"
      optional: true
      example: '{"DSTGRP":"501","IPADDR":"5.5.5.6","PORT":"57900","VRF":"management"}'

    - model_name: secondaryEntityName
      tf_name: secondary_entity_name
      type: String
      description: Secondary entity name, for example an overlay name when entity_type is configProfile
      optional: true
      example: MyNetwork_30000

    - model_name: secondaryEntityType
      tf_name: secondary_entity_type
      type: String
      description: Secondary entity type
      optional: true
      validator: 'OneOf("switch", "configProfile", "interface")'
      example: configProfile

    - model_name: markDeleted
      tf_name: deleted
      type: Bool
      ndfc_type: bool
      description: "Deleted policy"
      computed: true
      tf_hide: true

    - model_name: generatedConfig
      tf_name: generated_config
      type: String
      description: "Generated configuration"
      computed: true
```

## Validation and Preparation

During `terraform apply`, validate before calling create or update.

1. **Resolve the switch.** If the configuration starts with a management IP, call `GET /fabrics/{fabricName}/switches` and map `fabricManagementIp` to `switchId`. Use `additionalData.platformType` (`nx-os`, `ios-xe`, `ios-xr`, and so on) for later eligibility checks.
2. **Read the template schema.** Call `GET /configTemplates/{templateName}/parameters`. Inspect `parameters[]` (`name`, `optional`, `parameterType`, `defaultValue`, `annotations.IsMandatory`, `metaProperties.defaultValue`), `contentType` (`python`, `python3`, `python3Cli`, `pythonCli`, `templateCli`, `text`), `templateType`, `templateSubType`, and `supportedPlatforms`.
3. **Validate template inputs.** Compare `policy_parameters` to `parameters[]` from the template schema. Fail before create or update when a required parameter is missing and has no default, or when a key is not defined by the template. Treat `annotations.IsMandatory` of `false` or a conditional expression as optional. Honor `metaProperties.defaultValue` as well as `defaultValue`.
4. **Validate policy context.** Confirm template existence, switch membership in the fabric, payload shape, and entity pairing (`entityType`/`entityName`, optional secondary entity). If `is_policy_group` is false, `serial_numbers` must contain exactly one serial. If true, it must contain at least one.
5. **Check platform eligibility.** A valid parameter schema does not guarantee support for the target platform. Use `supportedPlatforms` plus the switch `platformType`.
6. **Apply backend content rules.** Freeform policies may be rejected for unsupported commands. NX-OS policies may also be subject to indentation validation.

**Do not infer platform support from parameter validity alone.** In a mixed-platform fabric, a template can have a valid schema while remaining incompatible with the selected switch.

Create and update are not a substitute for GET. After a 207 create or 204 update, read the object so `generatedConfig`, `markDeleted`, and timestamps are stored in state.

## Lifecycle Operations

`is_policy_group` selects the column. The Terraform resource name is always `nd_policy`.

| Lifecycle step | `is_policy_group = false` | `is_policy_group = true` |
| --- | --- | --- |
| Create | `POST /fabrics/{fabricName}/policies` | `POST /fabrics/{fabricName}/policyGroups` |
| Read | `GET /fabrics/{fabricName}/policies/{policyId}` | `GET /fabrics/{fabricName}/policyGroups/{policyGroupId}` |
| List / import matching | `GET /fabrics/{fabricName}/policies` | `GET /fabrics/{fabricName}/policyGroups` |
| Update | `PUT /fabrics/{fabricName}/policies/{policyId}` | `PUT /fabrics/{fabricName}/policyGroups/{policyGroupId}` |
| Deploy (`deploy = true`) | `POST .../policyActions/pushConfig` | Fabric switch deploy for `serial_numbers` |
| Mark delete | `POST .../policyActions/markDelete` | `POST .../policyGroups/actions/markDelete` |
| Direct delete | `DELETE .../policies/{policyId}` | `DELETE .../policyGroups/{policyGroupId}` |
| Bulk remove (fallback only) | `POST .../policyActions/remove` | `POST .../policyGroups/actions/remove` |
| Preview | `POST /fabrics/{fabricName}/switchActions/preview` | Same, for member switches |
| Cleanup deploy | `POST /fabrics/{fabricName}/switchActions/deploy` | Same, for member switches |

Requires-replace: `is_policy_group`, `template_name`. `serial_numbers` requires replace only when `is_policy_group` is `false`. `entity_type` and `entity_name` should also require replace; they are identity for the policy object. `fabric_name` is computed from the switch inventory and does not require replace.

Import ID is `policyId` only: `POLICY-26910` for a policy, `POLICY-GROUP-143310` for a group. Fabric is resolved by searching fabrics. The `POLICY-GROUP-` vs `POLICY-` prefix selects the read path.

## Deletion, Cleanup, and Replacement

### Mark delete

`POST .../policyActions/markDelete` or `POST .../policyGroups/actions/markDelete` records pending removal. The controller may assign a negative priority. The row can remain visible with `markDeleted: true`. Configuration is not removed from the device until a switch deploy runs.

### Direct delete

`DELETE .../policies/{policyId}` or `DELETE .../policyGroups/{policyGroupId}` removes the controller record (204). It does not guarantee that configuration has been removed from the device.

### Bulk remove

`POST .../policyActions/remove` and `POST .../policyGroups/actions/remove` are public bulk-removal operations. The single-resource destroy path does not call them first.

**Preflight rejection:** mark-delete can fail when stored content uses a PYTHON content type or when `generatedConfig` is empty. Evaluate the controller error, then fall back to direct delete plus switch deploy.

Generated child configurations, such as `switch_freeform_config`, can outlive the controller record. After direct deletion, deploy the affected switches and verify the device config. Removing a record alone is not evidence that the device is clean.

### Recommended destroy sequence

1. Mark-delete the policy or group by ID.
2. Preview with recompute when troubleshooting cleanup.
3. Deploy the target switch (policy) or member switches (group).
4. If the marked row remains, `DELETE` the controller record.
5. If mark-delete was rejected, `DELETE` then deploy cleanup.

### Template replacement

`templateName` is not an in-place change. Replacement is non-atomic:

1. Remove or mark the existing object.
2. Run switch cleanup deployment.
3. Create the replacement with the new template.
4. Push (`is_policy_group = false`) or deploy member switches (`is_policy_group = true`).

A failure between removal and recreation can leave the switch without the intended policy.

## API Contract Notes

- Create is always bulk (`policies` / `policyGroups` arrays) even when Terraform manages one object.
- Update and single-object GET/DELETE are not wrapped.
- Group GET-by-ID examples sometimes show `switchId` as an array. The list schema and write schema use `switchIds`. Normalize reads to `switchIds`.
- Group path parameter is `{policyGroupId}`; JSON field is still `policyId` with a `POLICY-GROUP-` prefix.
- `entityType` is camelCase (`switch`, `configProfile`, `interface`). `entityName` for a switch-level object is `SWITCH`.
- Write priority is 1–2000. Read priority can be negative after mark-delete. Do not send a negative priority on create or update.

## Operational Checklist

- [ ] Resolve and verify each switch serial before create or update.
- [ ] Read the template parameter schema and `supportedPlatforms`; validate required inputs and platform fit.
- [ ] Persist the returned `policyId` (`POLICY-*` or `POLICY-GROUP-*`) for later GET, PUT, push, mark-delete, and DELETE.
- [ ] After 207 create or 204 update, GET the object before writing Terraform state.
- [ ] Fail apply when a 207 item `status` is not `success`.
- [ ] Use `pushConfig` when `is_policy_group` is false; use switch deploy for group membership and for post-delete cleanup.
- [ ] Reject `serial_numbers` length other than 1 when `is_policy_group` is false.
- [ ] Distinguish pending mark-delete (`markDeleted: true`) from direct controller-record deletion.
- [ ] After deletion, deploy switch cleanup and verify the device configuration.
- [ ] For template replacement, plan for the non-atomic removal and recreation gap.
## References

### Key terms

- **Policy ID:** Controller-generated identifier. `POLICY-<n>` for a policy, `POLICY-GROUP-<n>` for a group.
- **Policy group ID path:** `{policyGroupId}` in group URLs; same value as the group’s `policyId` field.
- **Mark delete:** Pending deletion intent that still needs switch deployment.
- **Direct delete:** Targeted removal of a controller record.
- **Generated configuration:** Device configuration produced from the template and `templateInputs`.

### Useful resources

- [Get Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/get-policies/)
- [Create Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/create-policies/)
- [Update Policy](https://developer.cisco.com/docs/nexus-dashboard/latest/update-policy/)
- [List Group Policies](https://developer.cisco.com/docs/nexus-dashboard/latest/list-group-policies/)
- [Create Policy Groups For Switches](https://developer.cisco.com/docs/nexus-dashboard/latest/create-policy-groups-for-switches/)
- [Update Policy Group](https://developer.cisco.com/docs/nexus-dashboard/latest/update-policy-group/)
- [Policy entity type](https://developer.cisco.com/docs/nexus-dashboard/latest/policyentitytype)
- [Switch ID schema](https://developer.cisco.com/docs/nexus-dashboard/latest/schemas-switchid)
- Template parameters: `GET /api/v1/manage/configTemplates/{templateName}/parameters`
- Switch preview: `POST /api/v1/manage/fabrics/{fabricName}/switchActions/preview`
- Switch deploy: `POST /api/v1/manage/fabrics/{fabricName}/switchActions/deploy`
