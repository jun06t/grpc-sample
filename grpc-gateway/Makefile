.PHONY: prepare build

REGISTRY_DOMAIN := localimage

prepare:
	dep ensure -v

build: prepare
	docker build -f Dockerfile-gateway -t ${REGISTRY_DOMAIN}/gateway:1.0.0 .
	docker build -f Dockerfile-grpc -t ${REGISTRY_DOMAIN}/grpc:1.0.0 .
