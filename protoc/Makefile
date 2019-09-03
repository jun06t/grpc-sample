.PHONY: proto_basic proto_patah proto_plugin proto_format

proto_basic:
	protoc \
	-I=./basic \
	--go_out=./basic \
	./basic/*.proto

proto_path:
	protoc \
	-I=./path \
	--go_out=./path \
	./path/*.proto

proto_plugin:
	protoc \
	-I=./plugin \
	--go_out=plugins=grpc:./plugin \
	./plugin/*.proto

proto_format:
	find ./basic/ -name "*.proto" | xargs clang-format -i
	find ./path/ -name "*.proto" | xargs clang-format -i
	find ./grpc/ -name "*.proto" | xargs clang-format -i

