gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
run:
	go run main.go
clean:
	rm -rf pb/*
server:
	go run main.go server --port 8080
client:
	go run main.go client --adress 0.0.0.0

test:
	go test -cover -race ./...
