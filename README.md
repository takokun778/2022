# 2022

# Setup Local PC

```bash
brew install go
brew tap hashicorp/tap
brew install hashicorp/tap/terraform
```

# Terraform

## Terraform Cloud Workspaces Initialize

0. https://app.terraform.io/app
1. `Create a workspace`
2. choose `API-driven workflow`
3. Workspace Name: `2022`  
4. `Settings` > `General` > `Execution Mode` -> select `Local` -> `Save settings`

# CockroachDB

## Create Service Accounts

for terraform deploy

0. https://cockroachlabs.cloud
1. choose `Organization` -> `Access Management`
2. select `Service Accounts`
3. `Create Service Account`
4. `Create API key`
5. copy created token

# Secret

.env

```
COCKROACH_API_KEY=
TF_VAR_sql_user_password=
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASS=postgres
POSTGRES_PORT=15432
CONTAINER_NAME=
DATABASE_URL=
SLACK_TOKEN=
GITHUB_OWNER=
GITHUB_REPOSITORY=
SLACK_TOKEN=
SLACK_CHANNEL_ID=
IMAGE=
```
