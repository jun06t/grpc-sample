.PHONY: proto proto_format

build:
	docker build -f Dockerfile-client -t ${REGISTRY_DOMAIN}/client-side-lb:client .
	docker build -f Dockerfile-server -t ${REGISTRY_DOMAIN}/client-side-lb:server .

push:
	docker push ${REGISTRY_DOMAIN}/client-side-lb:client
	docker push ${REGISTRY_DOMAIN}/client-side-lb:server


proto: proto_format
	protoc \
	-I=./proto \
	--go_out=plugins=grpc:./proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

