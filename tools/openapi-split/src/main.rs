use openapiv3::{OpenAPI, Paths, ReferenceOr};
use serde_yaml;
use std::{collections::HashSet, fs};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let input_path = "openapi.yaml";
    let output_dir = "specs";
    fs::create_dir_all(output_dir)?;

    // 1. Load and Parse the input file using the openapiv3 crate
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

fn split_specs_by_tag(
    content: String,
) -> Result<Vec<(String, String)>, Box<dyn std::error::Error>> {
    let full_spec: OpenAPI =
        serde_yaml::from_str(&content).map_err(|e| format!("Failed to parse YAML spec: {}", e))?;
    println!("Identifying unique tags...");
    let all_tags = full_spec
        .tags
        .into_iter()
        .map(|tag| tag.name)
        .collect::<HashSet<_>>();
    if all_tags.is_empty() {
        println!("No tags found in any operation. Exiting.");
        return Ok(vec![]);
    }
    println!("Found tags: {:?}", all_tags);

    let mut specs_by_tag: Vec<(String, String)> = Vec::new();
    for target_tag in &all_tags {
        let mut new_spec = OpenAPI {
            openapi: full_spec.openapi.clone(),
            info: full_spec.info.clone(),
            paths: Paths::default(),
            components: full_spec.components.clone(),
            ..Default::default()
        };
        // Copy top-level extensions (servers, security, etc.)
        new_spec.extensions = full_spec.extensions.clone();

        // 3a. Filter Paths and Collect References
        for (path_str, path_item_ref_or) in full_spec.paths.iter() {
            match path_item_ref_or {
                ReferenceOr::Reference { .. } => todo!(),
                ReferenceOr::Item(path_item) => {
                    let operations = [
                        &path_item.get,
                        &path_item.put,
                        &path_item.post,
                        &path_item.delete,
                        &path_item.options,
                        &path_item.head,
                        &path_item.patch,
                        &path_item.trace,
                    ]
                    .into_iter()
                    .filter_map(|op| op.as_ref())
                    .cloned()
                    .collect::<Vec<_>>();
                    for operation in operations {
                        if operation.tags.iter().any(|tag| tag == target_tag) {
                            // Add path to new spec
                            new_spec
                                .paths
                                .paths
                                .insert(path_str.clone(), ReferenceOr::Item(path_item.clone()));
                            break;
                        }
                    }
                }
            }
        }
        let serialized_spec = serde_json::to_string_pretty(&new_spec)?;
        specs_by_tag.push((target_tag.to_string(), serialized_spec));
    }

    Ok(specs_by_tag)
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs;

    #[test]
    fn test_split_specs_by_tag() {
        let sample_spec = fs::read_to_string("test_spec.yaml").unwrap();
        let result = split_specs_by_tag(sample_spec).unwrap();
        assert!(result.len() == 2);
        let users_spec: OpenAPI = serde_json::from_str(
            result
                .iter()
                .filter(|(tag, _)| tag == "users")
                .next()
                .unwrap()
                .1
                .as_str(),
        )
        .unwrap();
        assert!(users_spec.paths.paths.len() == 1);
        assert!(
            users_spec.paths.paths["/users"]
                .as_item()
                .unwrap()
                .get
                .is_some()
        );

        assert!(
            users_spec.paths.paths["/users"]
                .as_item()
                .unwrap()
                .post
                .is_some()
        );

        let product_spec: OpenAPI = serde_json::from_str(
            result
                .iter()
                .filter(|(tag, _)| tag == "products")
                .next()
                .unwrap()
                .1
                .as_str(),
        )
        .unwrap();

        assert!(product_spec.paths.paths.len() == 1);
        assert!(
            product_spec.paths.paths["/products/{productId}"]
                .as_item()
                .unwrap()
                .get
                .is_some()
        );
    }
}
