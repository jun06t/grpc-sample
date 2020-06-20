# Server Reflection sample

## Compile
```
make proto
```

## Run gRPC server
```
go run server/main.go
```

## Use Evans
```
evans --port 8080 -r repl
```
or
```
evans --port 8080 --proto proto/helloworld.proto repl
```

