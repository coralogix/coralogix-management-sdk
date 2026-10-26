# iac

`iac` is the shared part of the code generators for infrastructure-as-code tools, such as a
provider or a Kubernetes operator. It reads the OpenAPI of the SDK, builds a model of one
resource, and rejects a resource whose API breaks the contract.

It is a separate Go module, so SDK users do not get its dependencies.

## Rule

`iac` holds only facts and checks about the API. They are true for every tool.
A fact or check that only one tool needs belongs in that tool's generator.

| Belongs in `iac` | Belongs in a tool's generator |
|---|---|
| "An enum zero value is `<PREFIX>_UNSPECIFIED`." | "A field name must not collide with another one after the tool renames it." |
| "The API sends the id in the Update body." | "A uint64 must fit the integer type of the tool." |
| "The SDK has the method `GetThing`." | "The tool's own client set has an accessor for the API client." |

The CI step "Keep tool code out of the iac module" checks this for the first tool, by its name.

## Packages

| Package | Job |
|---|---|
| `source` | Finds the `openapi.yaml` of the SDK version that a tool's `go.mod` pins. |
| `model` | Loads the OpenAPI, builds the model of one resource, and runs the contract checks. |
| `overrides` | The keys of a `behavior-overrides.yaml` file that describe the API of a released resource, and the `Policy` they make. |
| `sdkcheck` | Lists the SDK names that the generated code uses, and checks that the pinned SDK has them. |
| `issue` | One problem with a code, and a sorted report of problems. |

## How a tool uses it

The tool's generator calls `iac`. `iac` never calls the tool.

```go
input, _ := source.Resolve(".")                              // the pinned openapi.yaml
doc, _ := model.Load(input.OpenAPI)
policy := file.Exceptions().Policy()                         // the API keys of the tool's overrides file
// The tool adds the exceptions that its own keys cover to policy.
report := model.ValidateWithPolicy(doc, "Thing", model.OperationIDs{}, policy)
// The tool appends its own checks to report, then stops if report is not empty.
resource, _ := model.BuildWithPolicy(doc, "Thing", model.OperationIDs{}, policy)
tag, _ := sdkcheck.ResourceTag(doc, resource)
refs, _ := sdkcheck.Names(resource, tag, input.SDKModule, nil)
// The tool appends the names of its own code that the generated code calls.
loaded, _ := sdkcheck.Load(refs, input.ModuleRoot)
report = sdkcheck.Issues(refs, loaded, func(sdkcheck.Ref) sdkcheck.Source {
	return sdkcheck.Source{Module: input.SDKModule, Version: input.SDKVersion, Dir: input.SDKDir}
})
```

## List

`Resource.List` describes the operation that lists the resources. The model records facts. The tool decides what to generate from them.

| Field | Fact | Example |
|---|---|---|
| `ItemsField` | The response property that holds the items. | `connectors` |
| `ItemComponent` | The component of an item. | `Connector`, `PresetSummary` |
| `Full` | The item component is the resource component. | `true` for `Connector` |
| `IDField` | The item property that holds the resource id. It is the `IDField` of the resource. | `id` |
| `NameField` | `name` when the item has a string `name`, `""` otherwise. | `name` |
| `Paged` | A query parameter pages the List, or the response has a next page token. | `pageSize`, `pagination.page_token` |

Build finds the List in one of two ways:

- With no name, the List is the GET on the Create path. When the model cannot describe it, `List` is nil and `ListProblem` says why. This is not an issue, so the resource stays eligible.
- `OperationIDs.List`, or `api.operations.list` in the overrides file, names the List. Use it when the List is not on the Create path:

  ```yaml
  api:
    operations:
      list: PresetsService_ListPresetSummaries   # GET /v1/presets:summariesList
  ```

  A named List that the model cannot describe is an issue, because the user asked for it. `list` can be the only key of `api.operations`. The four lifecycle keys are set together, or not at all.

A List needs:

- GET, no path parameters, and no required parameter that does not page.
- The one tag of Get, so the SDK has it in the same package.
- A 200 response component with exactly one array of component objects. Other properties, such as `nextPageToken` or `totalCount`, are allowed.
- An item with the id property of the resource, of the same type.

| Code | Problem |
|---|---|
| `LIST_OPERATION_NOT_FOUND` | No operation has the named operationId. |
| `LIST_OPERATION_AMBIGUOUS` | More than one operation matches. |
| `LIST_OPERATION_METHOD_INCOMPATIBLE` | The operation is not a GET. |
| `LIST_PATH_PARAMETER_UNSUPPORTED` | The operation has a path parameter. |
| `LIST_REQUIRED_PARAMETER_UNSUPPORTED` | The operation requires a parameter that does not page. |
| `LIST_TAG_INCOMPATIBLE` | The operation does not have the one tag of Get. |
| `LIST_SINGLETON_UNSUPPORTED` | The resource is a singleton. |
| `LIST_RESPONSE_UNSUPPORTED` | The 200 response is missing or is not a component. |
| `LIST_ITEMS_NOT_FOUND` | The response has no array of component objects. |
| `LIST_ITEMS_AMBIGUOUS` | The response has more than one array of component objects. |
| `LIST_ITEM_ID_MISSING` | The item has no id property. |
| `LIST_ITEM_ID_TYPE_INCOMPATIBLE` | The item id does not have the type of the resource id. |

`sdkcheck.Names` adds the List call, the items field, and the id and name fields of a summary item only when `List` is set. A resource without a List needs no new SDK name.

A tool pins `iac` by the pseudo-version of a commit on `master`.
