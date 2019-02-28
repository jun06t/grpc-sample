# Unary(Simple) gRPC Metadata sample

## Compile
```
protoeasy --go --grpc ./proto
```

## Run gRPC server
```
go run server/main.go
```

## Run gRRC Client
```
go run client/main.go
```

Then it returns
```
2017/11/14 14:31:42 Reply:  Hello alice
```
