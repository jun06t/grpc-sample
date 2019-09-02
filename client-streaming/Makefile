.PHONY: proto

proto:
	protoc \
	-I=./proto \
	--go_out=plugins=grpc:./proto \
	./proto/*.proto
