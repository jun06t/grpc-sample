# gRPC Gateway sample

## Compile
```
make proto
```

## Run gRPC server
```
go run server/main.go
```

## Run gRPC Gateway
```
go run gateway/main.go
```

## Send Request
### Alive
```
curl http://localhost:3000/alive
```

This returns
```
{"status":true}
```

### Get User
```
curl http://localhost:3000/user/100
```

This returns
```
{"id":"100","name":"Alice","age":20}
```

### Get Users By Group
```
curl http://localhost:3000/user?group=ADMIN
```

This returns
```
{"group":"ADMIN","users":[{"name":"Alice","age":20},{"name":"Bob","age":24}]}
```

### Update User
```
curl -XPUT http://localhost:3000/user/100 -d '{"name": "bob", "age": 16}'
```

Then the grpc server's stdout shows
```
2017/11/14 14:25:52 update body is {id: 100, name: bob, age: 16}
```
