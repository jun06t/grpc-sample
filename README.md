# protoeasy-sample

## Install
Install protoeasy.
```
go get -u go.pedge.io/protoeasy/cmd/protoeasy
```

And also other libraries.
```
cat <<EOF | xargs go get -u
github.com/golang/protobuf/proto
github.com/golang/protobuf/protoc-gen-go
google.golang.org/grpc
github.com/gogo/protobuf/proto
github.com/gogo/protobuf/protoc-gen-gogofast
github.com/gogo/protobuf/gogoproto
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
EOF
```

## Compile
```
protoeasy --gogo --go-import-path=github.com/jun06t/protoeasy-sample --grpc --grpc-gateway .
```
