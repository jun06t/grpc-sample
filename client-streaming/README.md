# Client Streaming gRPC sample

## Compile
```
protoeasy --go --grpc ./proto
```

## Run gRPC server
```
cd server
go run main.go
```

## Run gRRC Client
```
cd client
go run main.go
```

Then it returns the following messages every second.
```
2017/11/14 14:37:10 success
```

And the server's directory has had uploaded image now.
```
$ ls
main.go      supercar.jpg
```
