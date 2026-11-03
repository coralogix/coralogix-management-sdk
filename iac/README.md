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

## A request shape that differs from the read shape

A released API can take a value in a Create or Update request under another name than the resource component (Get) has, or deeper in an object. A `requestAs` key on the field line and an `api` key of the overrides file describe it. Without them, the model matches top-level fields by identical property names, so such a resource gets `FIELD_LIFECYCLE_UNSUPPORTED`.

| Key | Meaning | Model |
|---|---|---|
| `requestAs` (field line) | The path of the property that holds the field in the Create and Update bodies, after the `requestWrapper` is removed, dot-separated through object properties, for example `integrationParameters.parameters`. Only a field of the resource component can have it. | `Policy.RequestNames`, `ResourceField.RequestPath` |
| `api.updateIDOutsideWrapper` | With `requestWrapper` and `updateIDInBody`, the Update body carries the id property beside the wrapper property, not inside it. | `Policy.UpdateIDOutsideWrapper`, `Resource.UpdateIDOutsideWrapper` |

The model matches the field by its request path, so the request property is not a separate request-only field, and the type checks run on the leaf. A path that only one body has makes the field absent from the other request. A container object that loses all its properties to the keys is not a field either. `api.idProperty` names the id field when the path parameter has another name (`integration_id` and `id`).

| Code | Problem |
|---|---|
| `REQUEST_NAME_UNKNOWN_FIELD` | The field line is not on a property of the resource component. |
| `REQUEST_NAME_NOT_FOUND` | Neither the Create nor the Update body has the path. |
| `REQUEST_NAME_NOT_OBJECT` | A segment before the last is not an object. |
| `REQUEST_NAME_DUPLICATE` | Two fields map to one path, or to a path that holds the other. |
| `REQUEST_NAME_CONFLICT` | A request body has a property with the name of the field, which `requestAs` also maps. |
| `OVERRIDE_UNUSED` | The path is the name of the field. |
| `UPDATE_ID_OUTSIDE_WRAPPER_UNSUPPORTED` | `updateIDOutsideWrapper` is set without `requestWrapper` and `updateIDInBody`. |
| `UPDATE_ID_NOT_IN_BODY` | The Update body has no id property beside the wrapper. |
| `UPDATE_ID_INSIDE_WRAPPER` | The wrapper also has the id property, so the key is wrong or ambiguous. |

```yaml
api:
  requestWrapper: metadata
  updateIDInBody: true
  updateIDOutsideWrapper: true   # Update is {id, metadata}
  idProperty: id                 # the path parameter is integration_id
types:
  DeployedIntegrationInformation:
    fields:
      definitionKey: {requestAs: integrationKey}
      definitionVersion: {requestAs: version}
      parameters: {requestAs: integrationParameters.parameters}
```

A tool renders the expand code from `ResourceField.RequestPath` and `Resource.UpdateIDOutsideWrapper`. The model also records the component that the request wrapper holds (`Operation.RequestComponent`, for example `IntegrationMetadata`) and the components of the objects that a dotted path goes through (`ResourceField.RequestObjects`). `sdkcheck.Names` reads all of them: the wrapper field has the request component as its type, a renamed field is a field of that component or of an object inside it, and the Update id beside the wrapper is a field of the Update body.

A `skip` line on a top-level field also ends the shape errors that the model finds inside the field, such as an enum without its zero value: the resource does not manage the field.

## A field that only the requests have

A Create or Update body can have a top-level field that the resource component does not have, such as a `teamId` that the API reads only on write. The model reports it as `FIELD_LIFECYCLE_UNSUPPORTED`. A `skip` line of the resource component can name it: the model then skips it before the lifecycle check, the enum checks, and the other field checks, and Build does not add it, so the generated requests never send it. A request that requires the field still makes the line an error (`SKIP_REQUIRED_FIELD`).

The resource component has no property of that name, so a tool that reports a field line without its field needs another list for a `skip` line. `model.RequestOnlyFields(doc, name, ids, policy)` returns the top-level fields that the Create or Update body has and the resource component does not, read as the field checks read them: without the request wrapper, and with the `requestAs` and `createAs` keys applied. With `Policy.UpdateMaskInBody`, the mask property of the Update body is not in the list: `Build` reads it as the mask. A `skip` line on one of them matches the contract. A line on a field that neither the resource component nor this list has matches nothing, and the tool reports it as unused.

```yaml
types:
  RuleGroup:
    fields:
      teamId: {skip: true}   # Create and Update have it, Get does not
```

## A collection resource

Some released APIs have no item lifecycle. One write replaces every item of a kind, and one Get lists the items of every kind. The resource is then the list: the response component of Get, with one list field. `api.overwrite` describes it.

| Key | Meaning | Model |
|---|---|---|
| `api.overwrite` | The resource is the whole list that one write replaces. `api.operations` names that write as `create`, `update`, and `delete`. Delete sends every list empty. The write response is not the resource, so a tool reads the resource with Get after Create and Update. The write path can differ from the Get path. | `Policy.Overwrite`, `Resource.Overwrite` |
| `api.overwrite.getQuery` | Query parameters that Get always sends, with their values, when Get lists the items of more than one resource. The value must be one that the parameter accepts. | `Policy.GetQuery`, `Resource.GetQuery` (sorted, with the type of each parameter) |
| `requestWrapper` (tool line of a nested component) | The request object of the component holds its fields in this property, for example `{"policy": {"name": ...}, "logRules": {...}}` for a Get item `{"name": ..., "logRules": ...}`. A field beside the wrapper keeps its place. A new field of the wrapper needs no line. | `Policy.RequestWrappers`, `Field.RequestPath`, `Field.RequestObjects` |
| `requestAs` (field line of a nested component) | As on the resource component, relative to the `requestWrapper` of the component. | `Policy.RequestNames`, `Field.RequestPath` |
| `negate` (field line) | The request sends the opposite of the boolean field, for example `enabled` sent as `disabled`. | `Policy.RequestNegated`, `Field.Negated` |

An overwrite resource is a singleton: no call takes an id. Each field that the write sends is a list or a map, because Delete sends it empty, and the list takes the item limits of the write: Get lists the items of every resource that its query can select. A required list that accepts `[]` is the contract here, not an issue.

When the write of one kind sends one arm of a `oneOf` of the Get item as a plain field, it always selects that arm. The other arms are not in the request, so they are computed. A tool can leave them out of the resource.

```yaml
api:
  overwrite:
    getQuery: {source_type: SOURCE_TYPE_LOGS}
  operations:
    create: PoliciesService_AtomicOverwriteLogPolicies
    get: PoliciesService_GetCompanyPolicies
    update: PoliciesService_AtomicOverwriteLogPolicies
    delete: PoliciesService_AtomicOverwriteLogPolicies
types:
  Policy:
    requestWrapper: policy         # the tool reads this line into Exceptions.RequestWrappers
    fields:
      enabled: {requestAs: disabled, negate: true}
```

`sdkcheck.Names` adds the builder method of each Get query parameter and the constant of an enum value, the inline body of the write for Delete, and the fields of a request item through its wrapper (`KindAThingWrite.Thing`, then `ThingCore.Name`). The response keeps its own names.

| Code | Problem |
|---|---|
| `OVERWRITE_SHAPE_UNSUPPORTED` | An operation has a path parameter, or Create, Update, and Delete are not one write. |
| `OVERWRITE_FIELD_NOT_COLLECTION` | The write sends a field that is not a list or a map. |
| `GET_QUERY_UNKNOWN` | Get has no query parameter of that name. |
| `GET_QUERY_VALUE_INVALID` | The parameter does not accept the value. |
| `REQUEST_WRAPPER_NOT_FOUND` | The request object of the component has no object property of that name. |
| `REQUEST_NAME_NOT_FOUND` | A nested `requestAs` path is not in the request object. |
| `REQUEST_NAME_CONFLICT` | A field is both beside and inside the wrapper, or two lines map one field. |
| `REQUEST_NAME_INCONSISTENT` | Create and Update hold a nested field at different paths. |
| `REQUEST_NEGATED_UNSUPPORTED` | The negated field is not a boolean. |
| `REQUEST_NAME_UNKNOWN_FIELD` | A nested line names a component that the resource does not hold, or matches no request of it. |

The golden contract is `model/testdata/overwritething.yaml`.

A tool pins `iac` by the pseudo-version of a commit on `master`.

## Create and Update with one inline body

openapi-generator makes one Go type of two inline bodies that are the same schema, and names it after the first: the Create body. When the Update body is an inline schema without a title that renders the same as the inline Create body, the model sets `Update.BodyAsCreate`, and `sdkcheck` expects the Update builder to take the Create body type. The parsing rule groups API has such a pair: `RuleGroupsServiceUpdateRuleGroup(ctx, id).RuleGroupsServiceCreateRuleGroupRequest(body)`. A title names each body, so two titled bodies keep their own types.
