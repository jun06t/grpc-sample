# Bidirectional Streaming gRPC sample

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

Then it returns the following messages.
```
2017/11/14 14:52:10 send message:  tokyo
2017/11/14 14:52:10 send message:  001
2017/11/14 14:52:10 send message:  Japan
2017/11/14 14:52:10 received message:  TOKYO
2017/11/14 14:52:10 received message:  001
2017/11/14 14:52:10 received message:  JAPAN
```
