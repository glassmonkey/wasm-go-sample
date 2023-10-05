run-go: build-go
	go run main.go

run-rust: build-rust
	go run main.go
# not working
build-go:
	 GOOS=wasip1 GOARCH=wasm go build -o main.wasm wasm-go/main.go

build-rust:
	 rustc wasm-rust/main.rs --target wasm32-wasi