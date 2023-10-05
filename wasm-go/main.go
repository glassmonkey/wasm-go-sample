package main

import "fmt"

//go:wasmexport main
func main() {
	fmt.Println("Hello, world! from Go")
}
