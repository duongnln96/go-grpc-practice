gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
run:
	go run main.go
clean:
	rm -rf pb/*
server:
	go run cmd/server/main.go -port 8080
client:
	go run cmd/server/main.go -adress 0.0.0.0

test:
	go test -cover -race ./...
