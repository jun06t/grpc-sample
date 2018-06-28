# Envoy front proxy sample

## Run servers
```
docker-compose up -d
```

## Scale out backend servers
```
docker-compose up --scale grpc=3
```

## Confirm Admin statistics
http://localhost:8001

## Send Request
### Alive
```
curl http://localhost:3000/alive
```

### Get User
```
curl http://localhost:3000/user/100
```

### Get Users By Group
```
curl http://localhost:3000/user?group=ADMIN
```

### Update User
```
curl -XPUT http://localhost:3000/user/100 -d '{"name": "bob", "age": 16}'
```
