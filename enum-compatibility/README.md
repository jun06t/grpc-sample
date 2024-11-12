# Unary(Simple) gRPC sample

## Compile
```
make proto
```

## Run gRPC server
Server depends on new enum proto.
```
go run server/main.go
```

## Run gRRC Client
Client depends on old enum proto.
```
go run client/main.go
```

Then it returns
```
2024/11/13 08:50:49 Reply:  Hello alice
2024/11/13 08:50:49 Role:  ULTIMATE
```
This means Enum's rename doesn't break compatibility.