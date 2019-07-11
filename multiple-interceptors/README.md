# Multiple Interceptors sample

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
client-inter1- before invoker
client-inter2- before invoker
client-inter3- before invoker
client-inter3- after invoker
client-inter2- after invoker
client-inter1- after invoker
2019/07/11 19:44:59 Reply:  Hello alice
```

And server says;
```
server-inter1- before handler
server-inter2- before handler
server-inter3- before handler
handler
server-inter3- after handler
server-inter2- after handler
server-inter1- after handler
```