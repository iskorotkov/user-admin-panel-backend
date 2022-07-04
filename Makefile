GO_POST_PROCESS_FILE='gofmt -w'
OPENAPI_SPEC_FILE=../user-admin-panel-common/openapi.yml

openapi:
	GO_POST_PROCESS_FILE=$(GO_POST_PROCESS_FILE) openapi-generator-cli generate \
		-g go-server \
		-i $(OPENAPI_SPEC_FILE) \
		-o ./ \
		-c openapi.config.yml
