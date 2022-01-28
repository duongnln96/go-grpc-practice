# go-grpc-practice
Practicing gPRC with Go


## Setup

1. Install ```protoc```
```
$ brew install protobuf
```

2. Install ```golang grpc``` library and the ```protoc-gen-go``` library
```
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

3. Run to generate codes
```
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
```

4. Create ```Makefile```
```
$ make clean
$ make gen
$ make run
```