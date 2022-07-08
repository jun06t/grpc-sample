# Fieldmask sample

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

Then it returns
```
2022/07/08 19:18:09 user:{id:"001"  name:"alice"  email:"alice@gmail.com"  age:20  address:{country:"Japan"  state:"Tokyo"  city:"Shibuya"  zipcode:"150-0000"}}
2022/07/08 19:18:09 user:{name:"alice"  age:20  address:{city:"Shibuya"}}
2022/07/08 19:18:09 user:{id:"001"  name:"bob"  age:20  address:{city:"Ikebukuro"  zipcode:"170-0000"}}
```
