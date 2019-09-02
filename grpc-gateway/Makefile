.PHONY: prepare build proto

REGISTRY_DOMAIN := jun06t

prepare:
	dep ensure -v

build:
	docker build -f Dockerfile-gateway -t ${REGISTRY_DOMAIN}/grpc-gateway:latest .
	docker build -f Dockerfile-grpc -t ${REGISTRY_DOMAIN}/grpc-backend:latest .

push:
	docker push ${REGISTRY_DOMAIN}/grpc-gateway
	docker push ${REGISTRY_DOMAIN}/grpc-backend

proto:
	protoc \
	-I=./proto \
	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:./proto \
	--grpc-gateway_out=logtostderr=true:./proto \
	./proto/*.proto
