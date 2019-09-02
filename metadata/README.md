# Unary(Simple) gRPC Metadata sample

## Compile
```
make proto
```

## Run gRPC server
```
go run server/main.go
```

## Run gRRC Client
```
go run client/main.go
```

Then client says;
```
2019/02/28 12:44:34 Reply:  Hello alice
2019/02/28 12:44:34 Header: map[content-type:[application/grpc] timestamp:[Feb 28 12:44:34.956833000]]
2019/02/28 12:44:34 Trailer: map[timestamp:[Feb 28 12:44:34.956868000]]
```

And server says;
```
request id from metadata:
 0. dwToyxFrRj969BcUDe7dc9
```
