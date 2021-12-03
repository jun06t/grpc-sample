.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./proto \
	--go_out=plugins=grpc:./proto \
	--go_opt=module=github.com/jun06t/grpc-sample/bazel/proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i

