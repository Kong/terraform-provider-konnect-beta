.PHONY: *

all: generate

.PHONY: generate
generate: speakeasy generate-plan-modifiers

speakeasy: check-speakeasy
	@rm -rf examples/resources
	speakeasy run --skip-versioning --output console --minimal
	@go mod tidy
	@go generate .
	@git checkout -- README.md
	@git clean -fd examples/data-sources docs/data-sources examples/README.md USAGE.md > /dev/null
	@bash ./rename-provider.sh

FILES=$(shell find internal/provider -type f | grep data_source | grep -v portallist | grep -v cloudgatewayprovideraccountlist)
remove-data-sources:
	@if [ -n "$(FILES)" ]; then rm $(FILES); fi
	@rm -r examples/data-sources

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }

OS=$(shell uname | tr "[:upper:]" "[:lower:]")
ARCH=$(shell uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/')
test:
	@cd tests/e2e; rm -rf .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup .terraform local-plugins
	mkdir -p tests/e2e/local-plugins/registry.terraform.io/kong/konnect/999.99.9/$(OS)_$(ARCH)
	go mod tidy
	go build -o tests/e2e/local-plugins/registry.terraform.io/kong/konnect/999.99.9/$(OS)_$(ARCH)/terraform-provider-konnect_v999.99.9
	@cd tests/e2e; terraform providers mirror ./local-plugins || true
	@cd tests/e2e; ls -R local-plugins; terraform init -plugin-dir ./local-plugins; terraform apply -auto-approve; terraform destroy -auto-approve


test-cleanup:
	@cd tests/e2e; rm -rf local-plugins .terraform .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup

unit-test:
	go test -v ./internal/sdk/internal/...

dev/use-local-shared-speakeasy:
	go mod edit -replace=github.com/Kong/shared-speakeasy=../shared-speakeasy
	go mod tidy

acceptance:
	@TF_ACC=1 go test -count=1 -v ./tests/resources

.PHONY: generate-plan-modifiers
generate-plan-modifiers:
	# MeshControlPlane is identfied by uuid so we can skip it
	mkdir -p "resouce-plan-modifiers"
	cat internal/provider/mesh*_resource.go \
	| grep "Resource struct" \
	| cut -d ' ' -f 2 \
	| sed 's/Resource$$//' \
	| grep -v "MeshControlPlane" \
	| xargs -n1 -I{} sh -c '\
		go run github.com/Kong/shared-speakeasy/generators/resource_plan_modifier@v0.0.8 \
		internal/provider/$$(echo {} | tr A-Z a-z)_resource_plan_modify.go {} terraform-provider-konnect-beta'
