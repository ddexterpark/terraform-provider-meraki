default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

#codegen:
	#brew install swagger-codegen
	#swagger-codegen generate -i https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json -l ruby -o $(pwd)/ruby
	terraform init && OTF_VAR_meraki_SWAGGER_URL="https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json" terraform plan

setup:
	export GOBIN=/Users/$USER/go/bin
	go env GOBIN