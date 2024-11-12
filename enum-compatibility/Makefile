.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./proto \
	--go_out=./proto \
	--go_opt=module=github.com/jun06t/grpc-sample/enum-compatibility/proto \
	--go-grpc_out=./proto \
	--go-grpc_opt=module=github.com/jun06t/grpc-sample/enum-compatibility/proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

