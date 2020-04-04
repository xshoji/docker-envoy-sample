# app

Hello world application.

```
$ go version
go version go1.13.7 darwin/amd64
```

# build

```
// build for linux
$ cd src
$ GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" app.go
```

> go - How to reduce compiled file size? - Stack Overflow  
> https://stackoverflow.com/questions/3861634/how-to-reduce-compiled-file-size
