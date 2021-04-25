
install-oapi-generator:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

.PHONY: generate-rest
generate-rest: install-oapi-generator
	oapi-codegen --config rest-openapi-gen.cfg.yaml api/rest-openapi.yaml
