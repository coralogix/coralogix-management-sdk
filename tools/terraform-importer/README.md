# **Guide to Using the Terraform Migration Script**

This guide provides step-by-step instructions on how to use the Terraform migration script effectively.

---

## **Prerequisites**

1. **Terraform Installed**  
   Ensure you have Terraform installed. You can download it [here](https://www.terraform.io/downloads).

2. **Go Installed**  
   Install Go from [golang.org](https://golang.org/dl/).

3. **Python Installed**  
   The script uses Python 3 for JSON processing, so make sure you have Python 3 installed.

4. **`hcl2json` Installed**  
   Install the `hcl2json` utility. You can find it [here](https://github.com/tmccombs/hcl2json).

5. **Homebrew (for macOS users)**  
   Used to manage dependencies via the `Makefile`.

---

## **Installing Prerequisites (macOS)**

To automatically install all required tools using Homebrew, run:

```bash
make install
```

This will:

- Install **Terraform**, **Go**, **Python3**, and **hcl2json** using Homebrew.
- Install Homebrew itself if it's not already available.

To verify installation:

```bash
make check
```

---

## **Usage**

### 1. Script Purpose

The script allows you to:

- Migrate Terraform configurations based on:
  - A folder containing a `terraform.tfstate` file.
  - A specific resource type (e.g., `alert`, `dashboard`).
- Generate a migration folder with cleaned and updated configurations.
- Specify the provider version interactively.
- Optionally apply the Terraform plan based on a prompt.

---

### 2. Running the Script with Make

Before running the script, you can set the required environment variables:

- `CORALOGIX_API_KEY`
- `CORALOGIX_ENV`

Then use the `Makefile`:

```bash
make run
```

> If you are not setting up the env vars, then make command will take care of that during execution

You will be prompted during execution:

```
Do you want to run terraform apply? (yes/no):
```

This prompt sets an environment variable that determines whether `terraform apply` is executed.

> ⚠️ You do **not** need to pass `APPLY_TERRAFORM=yes` manually. The prompt will handle it.

---

### 3. Interactive Steps

#### Step 1: Select Migration Type

You will be prompted to choose the migration type:

- **Option 1**: Migrate based on a folder containing a `terraform.tfstate` file.
  - Provide the path to the folder.
  - The script checks for the presence of the state file.
- **Option 2**: Migrate based on a specific resource type.
  - A list of resource types will be displayed.
  - Choose from options like: `alert`, `dashboard`, `archive_logs`, `events2metrics`, etc.

#### Step 2: Specify Provider Version

- Enter the Terraform provider version (e.g., `~>2.0.0`).
- If left blank, the default `>=2.0.0` is used.

---

### 4. What Happens Next

#### Step 3: Generate Migration Folder

- Folder structure is created based on your input.

#### Step 4: Run `generate_imports.go`

- A Go program generates `imports.tf` based on the resource or folder.

#### Step 5: Generate `provider.tf`

- The required provider version is written into `provider.tf`.

#### Step 6: Initialize Terraform

- Terraform is initialized using `terraform init`.

#### Step 7: Generate Configuration Plan

- `terraform plan -generate-config-out` creates `generated.tf`.

#### Step 8: JSON Cleanup

- `generated.tf` is converted to JSON.
- Null values are removed using a Python script.
- JSON is converted back to cleaned HCL.

#### Step 9: (Optional) Terraform Apply

- If you answered `yes` when prompted, the script runs `terraform apply`.
- If you answered `no`, this step is skipped.

#### Step 10: Cleanup

- Temporary files like `imports.tf`, `config.json`, and cleaned JSON are deleted.

---

## **Example Outputs**

### Migration Type Selection

```text
[INFO] Select the migration type:
[INFO] 1) Migrate based on a folder containing terraform.tfstate
[INFO] 2) Migrate based on a specific resource name
Enter your choice (1 or 2): 2
```

### Provider Version Prompt

```text
Enter the Terraform provider version to migrate to (e.g., ~>1.19.0): >=2.0.0
```

### Apply Prompt

```text
Do you want to run terraform apply? (yes/no): yes
```

### Logs During Execution

```text
2025-05-28 12:25:10 [INFO] Creating migration folder: ./alert_migration
2025-05-28 12:25:10 [INFO] Running generate_imports.go with -type...
2025-05-28 12:25:10 [INFO] Provider configuration generated in ./alert_migration/provider.tf.
2025-05-28 12:25:10 [INFO] Initializing Terraform...
...
2025-05-28 12:25:10 [INFO] Terraform apply completed.
2025-05-28 12:25:10 [INFO] Cleanup completed.
```

---

## **Notes**

- **Error Handling**  
  Uses `set -euo pipefail` for strict error checking. Logs errors with `[ERROR]` prefix and exits on failure.

- **Makefile**  
  Includes targets to check/install dependencies and run the script with a friendly interactive experience.

- **Safe Execution**  
  Terraform apply is skipped unless the user explicitly consents via prompt.

---
