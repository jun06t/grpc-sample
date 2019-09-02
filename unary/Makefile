.PHONY: proto

proto:
	protoc \
	-I=./proto \
	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:./proto \
	./proto/*.proto
