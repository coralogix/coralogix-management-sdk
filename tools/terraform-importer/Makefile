# Makefile for setup, validation, and execution of generate_and_migrate.sh on macOS

.DEFAULT_GOAL := help

TERRAFORM := terraform
GO := go
PYTHON := python3
HCL2JSON := hcl2json

## help: Show this help message
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

## check: Check if all required tools are installed
.PHONY: check
check: check-terraform check-go check-python check-hcl2json ## Check that required tools are installed

## install: Install all missing tools using Homebrew (macOS only)
.PHONY: install
install: ensure-brew install-terraform install-go install-python install-hcl2json ## Install any missing tools via Homebrew

## ensure-brew: Ensure Homebrew is installed
.PHONY: ensure-brew
ensure-brew:
	@command -v brew >/dev/null 2>&1 || ( \
		echo "Homebrew not found. Installing..."; \
		/bin/bash -c "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"; \
	)

## run: Install missing tools, prompt for env vars, and execute the script
.PHONY: run
run: install ## Install missing tools, prompt for env vars, and execute the script
	@echo "🔎 Checking required environment variables..."
	@if [ -z "$$CORALOGIX_API_KEY" ]; then \
		read -p "Enter CORALOGIX_API_KEY: " api_key; \
		export CORALOGIX_API_KEY=$$api_key; \
	else \
		echo "✔ CORALOGIX_API_KEY already set"; \
	fi; \
	if [ -z "$$CORALOGIX_ENV" ]; then \
		read -p "Enter CORALOGIX_ENV: " env; \
		export CORALOGIX_ENV=$$env; \
	else \
		echo "✔ CORALOGIX_ENV already set"; \
	fi; \
	echo "🚀 Running generate_and_migrate.sh..."; \
	if [ -x /opt/homebrew/bin/bash ]; then \
		/opt/homebrew/bin/bash ./generate_and_migrate.sh; \
	else \
		bash ./generate_and_migrate.sh; \
	fi


# -------------------- CHECKS ONLY --------------------

## check-terraform: Check if Terraform is installed
.PHONY: check-terraform
check-terraform:
	@command -v $(TERRAFORM) >/dev/null 2>&1 && echo "✔ Terraform found" || echo "✖ Terraform not found"

## check-go: Check if Go is installed
.PHONY: check-go
check-go:
	@command -v $(GO) >/dev/null 2>&1 && echo "✔ Go found" || echo "✖ Go not found"

## check-python: Check if Python 3 is installed
.PHONY: check-python
check-python:
	@command -v $(PYTHON) >/dev/null 2>&1 && echo "✔ Python 3 found" || echo "✖ Python 3 not found"

## check-hcl2json: Check if hcl2json is installed
.PHONY: check-hcl2json
check-hcl2json:
	@command -v $(HCL2JSON) >/dev/null 2>&1 && echo "✔ hcl2json found" || echo "✖ hcl2json not found"

# -------------------- INSTALLERS --------------------

## install-terraform: Install Terraform via Homebrew
.PHONY: install-terraform
install-terraform:
	@command -v $(TERRAFORM) >/dev/null 2>&1 || brew install terraform

## install-go: Install Go via Homebrew
.PHONY: install-go
install-go:
	@command -v $(GO) >/dev/null 2>&1 || brew install go

## install-python: Install Python 3 via Homebrew
.PHONY: install-python
install-python:
	@command -v $(PYTHON) >/dev/null 2>&1 || brew install python

## install-hcl2json: Install hcl2json via Homebrew
.PHONY: install-hcl2json
install-hcl2json:
	@command -v $(HCL2JSON) >/dev/null 2>&1 || brew install hcl2json
