# Error Details example

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

Then it returns
```
error code: INVALID_COUNTRY
2019/09/04 19:49:55 rpc error: code = InvalidArgument desc = some error occurred
exit status 1
```
