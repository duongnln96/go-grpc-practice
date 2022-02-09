# go-grpc-practice
Practicing gPRC with Go

## **Ref**:
[Tech School](https://dev.to/techschoolguru/series/7311)

## **Setup**

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
$ make test
$ make server
$ make client
```

# **Note**

**HTTP2**

Ref:

https://medium.com/@jacobtan/understanding-http-2-and-its-caveats-1e8200519c4c

https://developers.google.com/web/fundamentals/performance/http2#one_connection_per_origin

Difference with HTTP1

- Binary instead of textual
- Fully Multiplexing request
- One connection instead of multiple
- Flow control and prioritization of multiplexed streams
- Uses HPACK, algorithm to compress header for reducing overhead
- Server push —> QUIC (Quick UDP Internet Connection)
****

**There are 4 types of gRPC**

* **Unary**: 1 request - 1 response
* **Client streaming**: client sends stream of multiple messages, and expecting only 1 single response from server
* **Server streaming**: client sends only 1 request message, and the server replies with a stream of multiple responses
* **Bidirectional**:

****
**gRPC vs REST**

