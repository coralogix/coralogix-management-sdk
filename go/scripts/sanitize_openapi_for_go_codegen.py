#!/usr/bin/env python3
"""Prepare an OpenAPI spec for Go SDK code generation.

The checked-in OpenAPI spec may contain independent oneOf constraints under
allOf to model protobuf oneof groups. Those constraints are useful for schema
validation, but OpenAPI Generator's Go client flattens their branch `required`
fields into ordinary model fields. That makes every oneof arm required in Go.

For SDK codegen, strip only those validation-only allOf entries. The generated
package still receives the original spec file after code generation.
"""

from __future__ import annotations

import json
import sys
from pathlib import Path
from typing import Any


def is_required_branch(schema: Any) -> bool:
    if not isinstance(schema, dict):
        return False
    return set(schema) <= {"type", "required"} and isinstance(schema.get("required"), list)


def is_unset_branch(schema: Any) -> bool:
    if not isinstance(schema, dict):
        return False
    if not (set(schema) <= {"type", "not"} and isinstance(schema.get("not"), dict)):
        return False

    any_of = schema["not"].get("anyOf")
    return isinstance(any_of, list) and all(is_required_branch(branch) for branch in any_of)


def is_independent_oneof_constraint(schema: Any) -> bool:
    if not isinstance(schema, dict) or set(schema) != {"oneOf"}:
        return False

    branches = schema["oneOf"]
    return isinstance(branches, list) and branches and all(
        is_required_branch(branch) or is_unset_branch(branch) for branch in branches
    )


def required_oneof_fields(schema: Any) -> list[str] | None:
    """Return selected field names for required independent oneOf constraints."""
    if not is_independent_oneof_constraint(schema):
        return None

    fields: list[str] = []
    for branch in schema["oneOf"]:
        if is_unset_branch(branch):
            return None
        required = branch.get("required", [])
        if len(required) != 1:
            return None
        fields.append(required[0])

    return fields


def optional_oneof_fields(schema: Any) -> list[str] | None:
    """Return selected field names for optional independent oneOf constraints."""
    if not is_independent_oneof_constraint(schema):
        return None

    fields: list[str] = []
    has_unset_branch = False
    for branch in schema["oneOf"]:
        if is_unset_branch(branch):
            has_unset_branch = True
            for unset_branch in branch["not"]["anyOf"]:
                required = unset_branch.get("required", [])
                if len(required) != 1:
                    return None
                fields.append(required[0])
            continue

        required = branch.get("required", [])
        if len(required) != 1:
            return None
        fields.append(required[0])

    if not has_unset_branch:
        return None

    return list(dict.fromkeys(fields))


def sanitize(value: Any) -> Any:
    if isinstance(value, list):
        return [sanitize(item) for item in value]

    if not isinstance(value, dict):
        return value

    sanitized = {key: sanitize(child) for key, child in value.items()}

    all_of = sanitized.get("allOf")
    if isinstance(all_of, list):
        required_oneof_groups = []
        for entry in all_of:
            fields = required_oneof_fields(entry)
            if fields is None:
                continue
            required_oneof_groups.append(
                {
                    "index": len(required_oneof_groups),
                    "fields": fields,
                    "fieldList": ", ".join(fields),
                }
            )
        if required_oneof_groups:
            vendor_extensions = sanitized.setdefault("x-cx-codegen", {})
            vendor_extensions["requiredOneOfGroups"] = required_oneof_groups

        optional_oneof_groups = []
        for entry in all_of:
            fields = optional_oneof_fields(entry)
            if fields is None:
                continue
            optional_oneof_groups.append(
                {
                    "index": len(optional_oneof_groups),
                    "fields": fields,
                    "fieldList": ", ".join(fields),
                }
            )
        if optional_oneof_groups:
            vendor_extensions = sanitized.setdefault("x-cx-codegen", {})
            vendor_extensions["optionalOneOfGroups"] = optional_oneof_groups

        filtered = [
            entry for entry in all_of if not is_independent_oneof_constraint(entry)
        ]
        if filtered:
            sanitized["allOf"] = filtered
        else:
            sanitized.pop("allOf")

    return sanitized


def main() -> int:
    if len(sys.argv) != 3:
        print("usage: sanitize_openapi_for_go_codegen.py <input> <output>", file=sys.stderr)
        return 2

    input_path = Path(sys.argv[1])
    output_path = Path(sys.argv[2])

    with input_path.open() as f:
        spec = json.load(f)

    output_path.parent.mkdir(parents=True, exist_ok=True)
    with output_path.open("w") as f:
        json.dump(sanitize(spec), f, separators=(",", ":"), sort_keys=True)
        f.write("\n")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
