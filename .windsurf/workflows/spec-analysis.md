---
description: How to research OpenAPI specs to find URLs, schemas, and attributes for a given feature
---

# OpenAPI Spec Analysis Workflow

Use this workflow when given a feature name (e.g., "bootstrap", "credentials", "vpc-pairs") and you need to find the correct API endpoints, request/response schemas, and attribute details.

## 1. Locate the Spec Repository

The authoritative OpenAPI specs live at:
```
/Users/muralidm/code/go/src/openapi/specs/v1/
```

Directory structure:
- `manage/` — NDFC manage-plane APIs (fabrics, switches, networks, etc.)
- `infra/` — ND infrastructure APIs (users, auth, clusters, etc.)
- `shared/` — Shared components referenced across specs

Each domain has a top-level YAML (e.g., `manage-switches.yaml`, `manage-fabrics.yaml`) and a `shared/` or `components/` directory for reusable schemas.

## 2. Identify the Right Spec File

// turbo
Run a broad grep across the spec directory to find which file(s) mention the feature:
```bash
grep -rl "<feature_keyword>" /Users/muralidm/code/go/src/openapi/specs/v1/manage/
```

Example for bootstrap:
```bash
grep -rl "bootstrap" /Users/muralidm/code/go/src/openapi/specs/v1/manage/
```

This will return files like `manage-switches.yaml` and `shared/switchSchemas.yaml`.

## 3. Find the API Endpoints (Paths)

// turbo
Search for path definitions containing the feature keyword:
```bash
grep -n "/<feature>" /Users/muralidm/code/go/src/openapi/specs/v1/manage/manage-switches.yaml
```

Look for lines matching the pattern `  /fabrics/{fabricName}/<path>:` — these are the endpoint definitions.

For each endpoint, extract:
- **HTTP method** (get, post, put, delete)
- **operationId** — unique operation name
- **summary/description** — what it does
- **parameters** — path params, query params
- **requestBody** — `$ref` to the request schema
- **responses.200** — `$ref` to the response schema

Use `sed -n '<start>,<end>p'` to read the full endpoint block once you know the line numbers.

## 4. Resolve Schema References

Endpoints reference schemas via `$ref`. Common patterns:
- `"./shared/switchSchemas.yaml#/components/schemas/<SchemaName>"` — schema in shared file
- `"#/components/schemas/<SchemaName>"` — schema in same file

// turbo
Find the schema definition:
```bash
grep -n "<SchemaName>" /Users/muralidm/code/go/src/openapi/specs/v1/manage/shared/switchSchemas.yaml
```

Then read the full schema block with `sed -n`.

## 5. Resolve Composed Schemas (allOf)

Many schemas use `allOf` to compose from base schemas. Example:
```yaml
bootstrapImportSwitch:
  allOf:
    - $ref: '#/components/schemas/bootstrapBase'
    - $ref: '#/components/schemas/bootstrapCredential'
    - $ref: '#/components/schemas/bootstrapImportSpecific'
```

You must resolve **each** `$ref` to get the full attribute list. Read each referenced schema and merge their properties.

## 6. Extract Attribute Details

For each schema, extract a table of:
| Field | Type | Required | Description | Example |

Pay attention to:
- **required** arrays — which fields are mandatory
- **format** — `ip`, `cidr`, etc.
- **enum** references (e.g., `$ref: './common-enums.yaml#/...'`) — resolve these too
- **writeOnly** — fields only sent in requests, never returned
- **default** values
- **nested objects** — `$ref` to other schemas or inline `type: object`

## 7. Cross-Reference with Live API Response

If you have a sample API response (e.g., from logs or user-provided), compare it against the spec:
- Verify field names match
- Check for fields present in the response but missing from the spec (undocumented)
- Check for spec fields absent from the response (optional/conditional)

## 8. Map to Provider Code

Once you have the full API surface, map it to the provider:

### API Layer (`internal/manage/api/` or `internal/infra/api/`)
- Add URL constants matching the spec paths (prefixed with `/manage` or `/infra`)
- Add operation enum values for each endpoint
- Wire `GetUrl()` / `PostUrl()` dispatch

### YAML Definition (`generator/defs/*.yaml`)
- Add/update attributes matching the schema fields
- Set `type`, `optional`, `computed`, `mandatory`, `sensitive` based on spec
- Use `tf_hide: true` for fields not exposed to Terraform users

### Go Structs (`switch_ops.go` or similar)
- Create request/response structs with JSON tags matching the spec field names
- Use `omitempty` for optional fields

## 9. Checklist

- [ ] All endpoints identified (GET, POST, PUT, DELETE)
- [ ] All schemas fully resolved (including allOf composition)
- [ ] Required vs optional fields documented
- [ ] Sensitive fields identified (passwords, keys)
- [ ] URL constants added to API layer
- [ ] Request/response Go structs match spec field names
- [ ] Enum values resolved from shared enum files

## Tips

- The spec repo uses `$ref` extensively. Always follow refs to get actual types.
- Shared enums are typically in `shared/common-enums.yaml`.
- Shared parameter definitions (path params, query params) are in `shared/parameters.yaml`.
- When the spec and live API disagree, the live API is authoritative — but flag the discrepancy.
- Use `grep -c` to quickly count matches and gauge which file is most relevant.
