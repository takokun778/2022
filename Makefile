include .env
export

.PHONY: aqua
aqua: ## insatll aqua
	@brew install aquaproj/aqua/aqua

.PHONY: tool
tool: ## install tool
	@aqua i

.PHONY: ymlfmt
ymlfmt: ## yaml file format
	@yamlfmt

.PHONY: tfinit
tfinit: ## Terraform initialize
	@(cd terraform && terraform init)

.PHONY: tffmt
tffmt: ## Terraform format
	@(cd terraform && terraform fmt -recursive)

.PHONY: tflint
tflint: ## Terraform format check and terraform validate
	@(cd terraform && terraform fmt -recursive -check)
	@(cd terraform && terraform validate)

.PHONY: tfplan
tfplan: ## Terraform plan
	@(cd terraform && terraform plan)

.PHONY: tfapply
tfapply: ## Terraform apply
	@(cd terraform && terraform apply)

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
