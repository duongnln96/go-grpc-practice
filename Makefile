gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative
build:
	go build -o app
clean:
	rm -rf pb/*
server:
	go run main.go server --port 8080
client:
	go run main.go client --adress 0.0.0.0

test:
	go test -cover -race ./...
