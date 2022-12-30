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

.PHONY: db
db:
	@docker run --rm -d \
		-p $(POSTGRES_PORT):5432 \
		-e TZ=UTC \
		-e LANG=ja_JP.UTF-8 \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASS) \
		-e POSTGRES_INITDB_ARGS=--encoding=UTF-8 \
		--name $(CONTAINER_NAME) \
		postgres:14.6-alpine

.PHONY: psql
psql:
	@docker exec -it $(CONTAINER_NAME) psql -U postgres

.PHONY: stop
stop:
	@docker stop $(CONTAINER_NAME)

.PHONY: kapply
kapply:
	@envsubst '$$IMAGE','$$DATABASE_URL','$$GITHUB_OWNER','$$GITHUB_REPOSITORY','$$SLACK_TOKEN','$$SLACK_CHANNEL_ID' < k8s/cronjob.yaml > tmp.yaml
	@kubectl apply -f tmp.yaml
	@rm tmp.yaml

.PHONY: kdelete
kdelete:
	@kubectl delete -f k8s/cronjob.yaml

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
	@(cd terraform && terraform plan -lock=false)

.PHONY: tfapply
tfapply: ## Terraform apply
	@(cd terraform && terraform apply -lock=false)

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
