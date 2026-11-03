# Auto Generator Input Help

`generator/defs/<resource_name>.yaml` files act as input to generate provider code.

See [README.md](README.md) for the full generation workflow, prerequisites, and directory structure.

---

## YAML File Structure

A YAML file can define a `resource`, a `datasource`, or both.

```yaml
---
resource:
  name: <resource_name>
  generate_tf_resource: true
  category: manage
  doc_category: Interface Management
  description: Manages loopback interfaces for Nexus Dashboard
  skip_equality: false
  inherited_model: false
  parent_package: resource_interface_common
  parent_model: interface_common
  rest_endpoint: /manage/fabrics/{fabricName}/interfaces
  import_id:
    - "fabric1/FDO123/Loopback0"
  import_desc:
    - "fabric_name/switch_id/interface_name"
  attributes:
    - model_name: id
      tf_name: id
      description: Terraform unique ID for the resource
      type: String
      id: true
      computed: true
      payload_hide: true
      use_state: true
      example: "1"
    ...

datasource:
  name: <ds_name>
  generate_tf_resource: true
  category: infra
  doc_category: Authentication
  description: Reads local user data from Nexus Dashboard
  skip_equality: true
  attributes:
    ...
```

---

## Resource-Level Fields

### name
Name of the resource or datasource. Used for package names, file names, and the TF resource type (`nd_<name>`).

### generate_tf_resource
When `true`, generates a full TF resource with schema. When `false`, generates only the Go model and codec (used for shared model definitions like `interface_common`).

### category
Output subdirectory: `manage` or `infra`. Determines where generated files are placed (`internal/{category}/resource_<name>/`).

### doc_category
Documentation subcategory for the Terraform registry sidebar (e.g. `Interface Management`, `Authentication`). Must be listed in `generator/defs/doc_categories.yaml`.

### subcategory
TBD

### description
Resource-level description. Goes into the TF Schema `Description` field and generated documentation.

### inherited_model
Set to `true` when this YAML defines a shared model that child resources inherit (e.g. `interface_common.yaml`). Used with `generate_tf_resource: false`.

### parent_package
For child resources that inherit a shared model: the Go package name containing the parent model (e.g. `resource_interface_common`). The child's codec will reference the parent's Go struct.

### parent_model
For child resources: the `name` of the parent YAML definition (e.g. `interface_common`). Used together with `parent_package`.

### skip_equality
When `true`, skips generating `resource_compare_gen.go` (the DeepEqual-style comparison functions). Set to `true` for simple resources that don't need bulk comparison logic.

### rest_endpoint
Reference documentation for the NDFC REST API path. Not used by the generator for code generation.

### import_id
Array of sample import ID strings for documentation. Example:
```yaml
import_id:
  - "fabric1/FDO123/Loopback0"
  - "fabric1/FDO123/Loopback0,Loopback1"
```

### import_desc
Array of descriptions matching each `import_id` entry. Example:
```yaml
import_desc:
  - "fabric_name/switch_id/interface_name"
  - "fabric_name/switch_id/comma_separated_interface_names"
```

---

## Attribute-Level Fields

### model_name
Parameter name used in the NDFC REST API JSON payload.

### tf_name
Attribute name used in the Terraform `.tf` configuration file.

### description
Description of the attribute. Goes into the TF schema and into the generated documentation. Provide relevant, user-facing information.

### type, ndfc_type
`type` is the Terraform attribute type. `ndfc_type` controls how the value is represented in the JSON payload sent to NDFC.

**Default type mapping:**

| type | ndfc_type | JSON payload type |
|---|---|---|
| `String` | _(empty)_ | `string` |
| `Int64` | _(empty)_ | `int64` |
| `Int64` | `string` | `string` |
| `Bool` | _(empty)_ | `string` (`"true"` / `"false"`) |
| `Bool` | `bool` | `bool` (`true` / `false`) |

**Special and nested types:**

| type | ndfc_type | JSON payload type |
|---|---|---|
| `Map` | _(empty)_ | `map[string]<value type>` |
| `List` | _(empty)_ | `[]<struct>` (array of nested objects) |
| `Set` | `string` | `[]<struct>` (array, no ordering) |
| `List:String` | _(empty)_ | `[]string` |
| `List:String` | `csvstring` | `[]string` -- converted to/from CSV on marshal |
| `List:Int64` | _(empty)_ | `[]int64` |
| `Map:String` | _(empty)_ | `map[string]string` |
| `Float64` | _(empty)_ | `float64` |

### map_value_type
For `Map` type attributes, specifies the value type. Use `Struct` when the map values are nested objects with their own `attributes:` list. Example:
```yaml
- model_name: securityDomains
  type: Map
  map_value_type: Struct
  attributes:
    - model_name: roles
      type: List:String
      ...
```

### id
Set to `true` on the attribute that serves as the Terraform resource ID (typically the `id` attribute). Exactly one attribute per resource should have this.

---

### Mandatory, Computed, Optional, Default

These flags control Terraform's lifecycle behavior for the attribute.

**Required:** User must set this in the `.tf` config file.
Use `mandatory: true`.

**Optional:** User may set this but is not required to.
Use `optional: true`.

**Computed:** Filled in by the provider. User cannot set this in the config.
Use `computed: true` alone.

**Computed+Optional:** User can set the value, but if omitted, the provider fills it in (e.g. from the API response).
Use `optional: true` + `computed: true`. Also set `use_state: true` so the plan modifier preserves the value.

**Computed+Optional with default:** User can set the value, but if omitted a default is used.
Use `default_value: <value>`. No plan modifier needed since the default fills in the gap.

**Truth table:**

```
|---------|----------|----------|---------|---------------------|
|Mandatory| Computed | Optional | Default |  Result             |
|---------|----------|----------|---------|---------------------|
|  true   |   x      |   x      |  x      |  required           |  x = don't care
|  x      |   x      |   x      |  set    |  computed_optional  |
|  false  |   true   |   false  |  -      |  computed            |
|  false  |   false  |   true   |  -      |  optional            |
|  false  |   true   |   true   |  -      |  computed_optional   |
|---------|----------|----------|---------|---------------------|
|         Special case: all false          |  optional           |
|---------|----------|----------|---------|---------------------|
```

### default_value
Default value for the attribute. When set, the attribute becomes `computed_optional` -- if the user doesn't set it, this value is used. No plan modifier is needed.

### use_state
Adds the `UseStateForUnknown` plan modifier to the schema. Required for `computed_optional` attributes that do NOT have a `default_value`. Without this, Terraform treats the value as unknown on every plan, causing unnecessary diffs.

---

### Validation

### validator
Inline validator expression. The value is inserted directly into the generated schema code as a `stringvalidator`, `int64validator`, etc. Examples:

```yaml
# String validators
validator: 'OneOf("scp", "sftp")'
validator: 'LengthBetween(1, 128)'
validator: 'LengthAtLeast(1)'
validator: 'RegexMatches(regexp.MustCompile(`^[a-z0-9]+$`), "must be alphanumeric")'

# Int64 validators
validator: 'Between(1, 65535)'

# Cross-field validators
validator: 'ExactlyOneOf(path.MatchRoot("scp_sftp"))'
validator: 'LengthAtLeast(1), stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("ssh_key"))'

# List element validators
validator: 'ValueStringsAre(stringvalidator.OneOf("approver", "designer", "super-admin"))'
```

Multiple validators can be comma-separated within a single string.

### custom_validator
A custom validator function call. Unlike `validator` which uses standard terraform-plugin-framework validators, this references a hand-written Go function in the provider codebase:

```yaml
custom_validator: 'remoteStorageHostKeyValidator()'
```

### enum_values
List of allowed values for the attribute. Used for documentation and can be used by the generator to produce `OneOf` validators. Example:

```yaml
enum_values:
  - access
  - trunk
  - routed
  - monitor
```

### min_int
Minimum value for provider-level `Int64` attributes (used in the provider schema, e.g. timeout, max_retries). Example:

```yaml
- model_name: timeout
  type: Int64
  min_int: 1
```

### sensitive
Marks the attribute as sensitive in the TF schema. Terraform will redact the value in plan output and state display. Use for passwords, keys, and secrets:

```yaml
- model_name: password
  type: String
  sensitive: true
```

---

### Payload Control

### ndfc_nested
Flattens nested NDFC JSON structures into the TF config. The value is an array of keys representing the nesting path. The attribute appears flat in the TF config but is marshalled/unmarshalled into the nested JSON structure.

Example: `routingTag` lives at `configData.networkOS.policy.routingTag` in the API payload but appears as a top-level attribute in the TF config:
```yaml
- model_name: routingTag
  tf_name: routing_tag
  type: String
  ndfc_nested:
    - configData
    - networkOS
    - policy
```

### payload_hide
Attribute is NOT sent to NDFC in the API payload. Used for control flags that trigger provider-side logic only. Example: `deploy: true` triggers a deploy operation but is not part of the NDFC REST payload.

### tf_hide
Attribute is NOT exposed in the Terraform schema (user cannot see or set it). Used for internal control values the provider needs to inject into the API payload. Example: `mode: "managed"` for loopback interfaces -- always hardcoded, never user-visible.

### include_empty
By default, generated Go struct tags include `omitempty` so empty fields are omitted from the JSON payload. Setting `include_empty: true` removes `omitempty`, so that the field is always present in the payload even when empty. Needed for NDFC APIs that require the field to be present.

### handle_empty
Only for `Int64` types. NDFC sometimes returns `""` (empty string) for integer fields that are not set. Standard `json.Unmarshal` fails on this. When `handle_empty: true` is set, a custom type (`Int64Custom`) is used with overridden marshallers that handle empty strings gracefully.

### unmarshal_tag
For NDFC APIs that use different field names in POST vs GET responses. `model_name` is used for POST (marshal). `unmarshal_tag` is used for GET (unmarshal). A duplicate struct member is generated to handle this:

```yaml
- model_name: switchId
  unmarshal_tag: serialNumber
```

---

### Replace Behavior

### requires_replace
For attributes **inside** a bulk resource's Map/List entry. When the attribute value changes, the entry must be deleted and recreated (not updated in-place).

### tf_requires_replace
For **top-level** resource attributes. When the attribute value changes, the entire resource must be destroyed and recreated.

---

### Bulk Resource Fields

These fields are used for resources that contain arrays of entries (bulk/map resources).

### filtered
Set on an array entry field. When this field's value is `true` on an entry, that entry is filtered out (skipped) when populating the TF model from the API response. Used to exclude non-managed entries.

### bulk
Set on an array attribute when it has a large number of entries. Generates a `map[list_key]*entry` for O(1) lookup instead of linear search.

### list_key
The field within each array entry used as the map key when `bulk: true`. Example: if `list_key` is set on `switchId`, a map `map[switchId]*entry` is created.

### sort
When `true`, sort helper functions are generated so `sort.Sort()` can be used on the array.

---

### Documentation Fields

### example
Value used to construct the generated HCL example in `examples/resources/nd_<name>/resource.tf`. Important for proper documentation generation.

For `Map` type attributes, the `example` value is used as the map key in the generated example.

To suppress an attribute from the generated example, omit the `example` field entirely.

---

## Example: Minimal Resource YAML

```yaml
---
resource:
  name: interface_loopback
  generate_tf_resource: true
  category: manage
  doc_category: Interface Management
  description: Manages loopback interfaces on Nexus Dashboard Fabric Controller
  parent_package: resource_interface_common
  parent_model: interface_common
  import_id:
    - "Trial_VXLAN/FDO22412H7U/Loopback0"
  import_desc:
    - "fabric_name/switch_id/interface_name"
  attributes:
    - model_name: id
      tf_name: id
      type: String
      id: true
      computed: true
      payload_hide: true
      use_state: true
      description: The unique identifier of the resource

    - model_name: fabricName
      tf_name: fabric_name
      type: String
      mandatory: true
      tf_requires_replace: true
      payload_hide: true
      example: Trial_VXLAN
      description: Name of the fabric

    - model_name: deploy
      tf_name: deploy
      type: Bool
      ndfc_type: bool
      default_value: true
      payload_hide: true
      description: Deploy interfaces after creation
      example: true
```
