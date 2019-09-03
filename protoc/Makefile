.PHONY: proto proto_format

proto_basic:
	protoc \
	-I=./basic \
	--go_out=./basic \
	./basic/*.proto

proto_format:
	find ./basic/ -name "*.proto" | xargs clang-format -i
	find ./path/ -name "*.proto" | xargs clang-format -i
	find ./grpc/ -name "*.proto" | xargs clang-format -i

