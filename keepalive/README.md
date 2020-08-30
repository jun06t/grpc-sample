# Keepalive gRPC sample

## Compile
```
make proto
```

## Run gRPC server
```
GODEBUG=http2debug=2 go run server/main.go
```

## Run gRRC Client
```
GODEBUG=http2debug=2 go run client/main.go
```

Then it returns
```
2017/11/14 14:31:42 Reply:  Hello alice
```
