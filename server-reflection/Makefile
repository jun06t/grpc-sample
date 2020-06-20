.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./proto \
	--go_out=plugins=grpc:./proto \
	./proto/*.proto

gogoproto: proto_format
	protoc \
	-I=./proto \
	--gogo_out=plugins=grpc:./proto \
	./proto/*.proto

gofastproto: proto_format
	protoc \
	-I=./proto \
	--gofast_out=plugins=grpc:./proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

