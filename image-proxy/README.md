# Image proxy sample

This is a proxy to convert jpeg/png to webp.

## Compile
```
$ make proto
```

## Prepare
You should install webp tool because this proxy runs ``cwebp`` command directly.
```
$ brew install webp
```

## Run gRPC server
```
$ cd server
$ go run main.go
```

## Run gRRC Client
```
$ cd client
$ go run main.go
```

Then you'll get webp image on client directory.
```
$ ls
001.webp      main.go       testimage.png
```
