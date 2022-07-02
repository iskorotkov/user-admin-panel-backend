openapi-validate:
	openapi-generator-cli validate \
		-i openapi.yml

openapi-generate:
	GO_POST_PROCESS_FILE='gofmt -w' openapi-generator-cli generate \
		-g go-server \
		-i openapi.yml \
		-o ./ \
		-c openapi.config.yml
