.PHONY: proto_before proto_after proto_format clean

BEFORE_PATH=before

proto_before: proto_format
	protoc \
	-I=./${BEFORE_PATH} \
	--go_out=:./${BEFORE_PATH} \
	./${BEFORE_PATH}/team/*.proto \
	./${BEFORE_PATH}/person/*.proto

AFTER_PATH=after

proto_after: proto_format
	protoc \
	-I=./after \
	--go_out=:./${AFTER_PATH} \
	--go_opt=module=github.com/jun06t/grpc-sample/go-package-option/${AFTER_PATH} \
	./${AFTER_PATH}/team/*.proto \
	./${AFTER_PATH}/person/*.proto

proto_format:
	find ./before -name "*.proto" | xargs clang-format -i
	find ./after -name "*.proto" | xargs clang-format -i

clean:
	find ./before -name "*.pb.go" | xargs rm
	find ./after -name "*.pb.go" | xargs rm
