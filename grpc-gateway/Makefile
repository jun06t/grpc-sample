.PHONY: prepare build

REGISTRY_DOMAIN := jun06t

prepare:
	dep ensure -v

build:
	docker build -f Dockerfile-gateway -t ${REGISTRY_DOMAIN}/grpc-gateway:latest .
	docker build -f Dockerfile-grpc -t ${REGISTRY_DOMAIN}/grpc-backend:latest .

push:
	docker push ${REGISTRY_DOMAIN}/grpc-gateway
	docker push ${REGISTRY_DOMAIN}/grpc-backend
