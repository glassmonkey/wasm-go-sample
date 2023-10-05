run:
	go run main.go

build:
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm wasm/main.go