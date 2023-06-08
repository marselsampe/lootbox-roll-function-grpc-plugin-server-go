GOLANG_DOCKER_IMAGE := golang:1.19

run:
	@test -n "$(ENV_FILE_PATH)" || (echo "ENV_FILE_PATH is not set" ; exit 1)
	docker run -t --rm -u $$(id -u):$$(id -g) --env-file $(ENV_FILE_PATH) -v $$(pwd):/data/ -w /data/ \
		-e GOCACHE=/data/.cache/go-build \
		-e GOPATH=/data/.cache/mod \
    		 $(GOLANG_DOCKER_IMAGE) go run .

gen-api-client:
	docker run -t --rm -u $$(id -u):$$(id -g) \
    	-v $$(pwd):/workspace -w /workspace \
    	quay.io/goswagger/swagger generate client -f pkg/client/platformservice/swagger.json \
    		--target pkg/client/platformservice/openapi2 \
    		--operation=updateLootBoxPluginConfig --operation=deleteLootBoxPluginConfig \
    		--model=LootBoxPluginConfigUpdate --model=LootBoxPluginConfigInfo --model=ValidationErrorEntity \
    		--model=AppConfig --model=BaseCustomConfig --model=PublicCustomConfigInfo --model=BaseTLSConfig \
    		--model=FieldValidationError --model=ErrorEntity