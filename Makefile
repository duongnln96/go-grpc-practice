.PHONY: gen
gen:
	@echo "====== Generating ======"
	@protoc --proto_path=proto proto/*.proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--fatal_warnings
	@echo "====== Completed ======"

.PHONY: build
build:
	@echo "====== Building app ======"
	@go build -o app
	@echo "====== Completed ======"

.PHONY: clean
clean:
	@echo "====== Cleaning ======"
	@rm -rf pb/*
	@rm -rf swagger/*
	@echo "====== Completed ======"

.PHONY: server
server:
	@go run main.go grpc_server --port 8080

.PHONY: client
client:
	@go run main.go grpc_client --address 0.0.0.0

.PHONY: test
test:
	@echo "====== Start testing ======"
	@go test -cover -race ./...
	@echo "====== Completed ======"

.PHONY: tmp
tmp:
	@echo "====== Cleaning tmp folder ======"
	@rm -rf tmp/*
	@echo "====== Completed ======"

.PHONY: cert
cert:
	@echo "====== Generating ======"
	cd cert; ./gen.sh; cd ..
	@echo "====== Completed ======"
