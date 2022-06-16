
GRAFANA_TARGET_VERSION ?= main
SWAGGER_CODEGEN_CLI_TAG ?= latest

generate:
	docker run --rm \
	--user $$(id -u):$$(id -g) \
	-v $$(pwd):$$(pwd) \
	swaggerapi/swagger-codegen-cli:${SWAGGER_CODEGEN_CLI_TAG} generate \
	-i https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json \
	-l go \
	-o $$(pwd)/goclient \
	--model-name-suffix=Model \
	--additional-properties packageName=goclient \
	-t $$(pwd)/codegen/templates
	goimports -w -v $$(pwd)/goclient

clean:
	rm -rf $$(pwd)/goclient