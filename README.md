# **go-grpc-practice**

Practicing gPRC with Go

## **Ref**

[Tech School](https://dev.to/techschoolguru/series/7311)

## **Setup**

1. Install ```protoc```

    ```sh
    brew install protobuf

    brew install clang-format
    ```

2. Install ```golang grpc``` library and the ```protoc-gen-go``` library.

    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

3. Run to generate codes

    ```sh
    protoc --proto_path=proto proto/*.proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --fatal_warnings
    ```

4. Create ```Makefile```

    ```sh
    make clean
    make gen
    make run
    make test
    make server
    make client
    ```

## **Setup VSCode**

1. Extension: vscode-proto3, clang-format

2. setting.json

  ```macos
  command + shift + P
  open setting user settings (JSON)
  ```

  ```json
  "protoc": {
    "path": "/usr/local/bin/protoc",
    "options": [
        "--proto_path=proto"
    ]
  },
  ```

## **Note**

### **HTTP2**

Ref:

<https://medium.com/@jacobtan/understanding-http-2-and-its-caveats-1e8200519c4c>

<https://developers.google.com/web/fundamentals/performance/http2#one_connection_per_origin>

***Difference with HTTP1***

- Binary instead of textual
- Fully Multiplexing request
- One connection instead of multiple
- Flow control and prioritization of multiplexed streams
- Uses HPACK, algorithm to compress header for reducing overhead
- Server push —> QUIC (Quick UDP Internet Connection)

****

### **There are 4 types of gRPC**

- **Unary**: 1 request - 1 response
- **Client streaming**: client sends stream of multiple messages, and expecting only 1 single response from server
- **Server streaming**: client sends only 1 request message, and the server replies with a stream of multiple responses
- **Bidirectional**:

****

### **gRPC vs REST**

- gRPC uses HTTP2
- gRPC uses `protocol buffer` to serialize payload data, which is binary and smaller, while `REST` uses JSON, which is text and larger.
- The API contract in gRPC is strict, and required to be clearly defined in proto file, while REST API using the third-party like OpenAPI or Swagger.
- Streaming is bidirectional in gRPC, while only 1 way request from client to server in REST
- One thing REST is still better, that is the browser support.

### **Where and when to use gRPC?**

- Microservices:
  - Low latency and high throughput communication.
  - Strong API contract
- Polyglot environments:
  - Code generation out of the box for many languages
- Point-to-point realtime comunication
- Network contrained environments. Lightweight message format

## **Proto file**

```proto
syntax = "proto3";

// default generate package name "techschool_pcbook"
// using to defination and importing proto file
package techschool.pcbook;

// rename package to "pb"
option go_package = "pb";
```

## gRPC interceptor

It's like a middleware function that can be added both server-side and client-side

- Server-side interceptor is a function that will be called by the gRPC server before reaching the actual RPC method. It can be used for multiple purposes such as logging, tracing, rate-limiting, authentication and authorization.

- Similarly, client-side interceptor is a function that will be called by the gRPC client before invoking the actual RPC
