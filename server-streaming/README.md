# Server Streaming gRPC sample

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

Then it returns the following messages every second.
```
message:"article1" 
message:"article2" 
message:"article3" 
```
