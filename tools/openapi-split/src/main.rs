use serde_json::{Map, Value};
use serde_yaml;
use std::{collections::HashSet, fs};

const OPERATION_METHODS: [&str; 8] = [
    "get", "put", "post", "delete", "options", "head", "patch", "trace",
];

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let input_path = "openapi.yaml";
    let output_dir = "specs";
    fs::create_dir_all(output_dir)?;

    // 1. Load the input file
    println!("Reading and parsing input spec: {}", input_path);
    let content = fs::read_to_string(input_path).map_err(|e| {
        format!(
            "Could not read file {}: {}. Did you create openapi.yaml?",
            input_path, e
        )
    })?;

    let specs_by_tag = split_specs_by_tag(content)?;

    for (tag, spec) in specs_by_tag {
        fs::write(
            format!(
                "{}/{}.json",
                output_dir,
                tag.replace(" ", "_").to_lowercase()
            ),
            spec,
        )?;
    }

    Ok(())
}

// The spec is kept as raw JSON, not as `openapiv3` crate types. The crate
// models a `$ref` as the link only, so it dropped keys next to a `$ref` (for
// example `x-coralogix-presence` or `default`), which OpenAPI 3.1 allows.
fn split_specs_by_tag(
    content: String,
) -> Result<Vec<(String, String)>, Box<dyn std::error::Error>> {
    let full_spec: Value =
        serde_yaml::from_str(&content).map_err(|e| format!("Failed to parse YAML spec: {}", e))?;
    let full_spec = full_spec
        .as_object()
        .ok_or("Failed to parse YAML spec: the root is not an object")?;
    println!("Identifying unique tags...");
    let all_tags = full_spec
        .get("tags")
        .and_then(Value::as_array)
        .into_iter()
        .flatten()
        .filter_map(|tag| tag.get("name").and_then(Value::as_str))
        .map(str::to_string)
        .collect::<HashSet<_>>();
    if all_tags.is_empty() {
        println!("No tags found in any operation. Exiting.");
        return Ok(vec![]);
    }
    println!("Found tags: {:?}", all_tags);

    let empty_paths = Map::new();
    let full_paths = full_spec
        .get("paths")
        .and_then(Value::as_object)
        .unwrap_or(&empty_paths);

    let mut specs_by_tag: Vec<(String, String)> = Vec::new();
    for target_tag in &all_tags {
        // Copy only openapi, info, the tag's paths, components, and the
        // top-level x-* extensions, as the crate-based version did.
        let mut new_spec = Map::new();
        for key in ["openapi", "info"] {
            if let Some(value) = full_spec.get(key) {
                new_spec.insert(key.to_string(), value.clone());
            }
        }

        // 3a. Filter Paths
        let mut paths = Map::new();
        for (path_str, path_item) in full_paths {
            if path_str.starts_with("x-") {
                continue;
            }
            if path_item.get("$ref").is_some() {
                return Err(format!("Path item references are not supported: {}", path_str).into());
            }
            if path_has_tag(path_item, target_tag) {
                paths.insert(path_str.clone(), path_item.clone());
            }
        }
        new_spec.insert("paths".to_string(), Value::Object(paths));

        if let Some(components) = full_spec.get("components") {
            new_spec.insert("components".to_string(), components.clone());
        }
        // Copy top-level extensions
        for (key, value) in full_spec {
            if key.starts_with("x-") {
                new_spec.insert(key.clone(), value.clone());
            }
        }

        let serialized_spec = serde_json::to_string_pretty(&Value::Object(new_spec))?;
        specs_by_tag.push((target_tag.to_string(), serialized_spec));
    }

    Ok(specs_by_tag)
}

fn path_has_tag(path_item: &Value, target_tag: &str) -> bool {
    OPERATION_METHODS
        .iter()
        .filter_map(|method| path_item.get(*method))
        .filter_map(|operation| operation.get("tags").and_then(Value::as_array))
        .any(|tags| tags.iter().any(|tag| tag.as_str() == Some(target_tag)))
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs;

    fn spec_for_tag(result: &[(String, String)], tag: &str) -> Value {
        let spec = &result.iter().find(|(name, _)| name == tag).unwrap().1;
        serde_json::from_str(spec).unwrap()
    }

    #[test]
    fn test_split_specs_by_tag() {
        let sample_spec = fs::read_to_string("test_spec.yaml").unwrap();
        let result = split_specs_by_tag(sample_spec).unwrap();
        assert!(result.len() == 2);

        let users_spec = spec_for_tag(&result, "users");
        let users_paths = users_spec["paths"].as_object().unwrap();
        assert!(users_paths.len() == 1);
        assert!(users_paths["/users"].get("get").is_some());
        assert!(users_paths["/users"].get("post").is_some());

        let product_spec = spec_for_tag(&result, "products");
        let product_paths = product_spec["paths"].as_object().unwrap();
        assert!(product_paths.len() == 1);
        assert!(product_paths["/products/{productId}"].get("get").is_some());
    }

    #[test]
    fn test_split_keeps_keys_next_to_ref() {
        let sample_spec = fs::read_to_string("test_spec.yaml").unwrap();
        let result = split_specs_by_tag(sample_spec).unwrap();

        let users_spec = spec_for_tag(&result, "users");
        let kind = &users_spec["components"]["schemas"]["Item"]["properties"]["kind"];
        assert_eq!(kind["$ref"], "#/components/schemas/Kind");
        assert_eq!(kind["x-coralogix-presence"], true);
        assert_eq!(kind["default"], "KIND_FAST");
    }

    #[test]
    fn test_split_keeps_only_the_same_top_level_keys() {
        let sample_spec = fs::read_to_string("test_spec.yaml").unwrap();
        let result = split_specs_by_tag(sample_spec).unwrap();

        let users_spec = spec_for_tag(&result, "users");
        let keys = users_spec.as_object().unwrap().keys().cloned().collect::<Vec<_>>();
        assert_eq!(keys, ["openapi", "info", "paths", "components"]);
    }
}
