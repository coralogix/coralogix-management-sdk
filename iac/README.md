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

A tool pins `iac` by the pseudo-version of a commit on `master`.
