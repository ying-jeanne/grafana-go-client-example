
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

run-test-image-locally:
	docker rm -f test-grafana
	docker run --name test-grafana -d -p 3000:3000 grafana/grafana:latest

test: clean-test run-test-image-locally
	sleep 3
	go test -cover -race -vet all -mod readonly ./... 	