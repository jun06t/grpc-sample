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
Server returns random response.

```
error-details $ go run client/main.go
error code: EXPIRED_RECEIPT
2019/09/05 08:00:07 rpc error: code = InvalidArgument desc = some error occurred
exit status 1
error-details $ go run client/main.go
2019/09/05 08:00:08 Reply:  Hello alice
error-details $ go run client/main.go
handle BadRequest case
2019/09/05 08:00:10 rpc error: code = InvalidArgument desc = some error occurred
exit status 1
error-details $ go run client/main.go
error code: INVALID_COUNTRY
2019/09/05 08:00:12 rpc error: code = InvalidArgument desc = some error occurred
exit status 1
error-details $ go run client/main.go
handle QuotaFailure case
2019/09/05 08:00:13 rpc error: code = InvalidArgument desc = some error occurred
exit status 1
```
