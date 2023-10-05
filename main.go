package main

import (
	_ "embed"
	"fmt"
	"github.com/bytecodealliance/wasmtime-go/v13"
	"log"
)

func main() {
	store := wasmtime.NewStore(wasmtime.NewEngine())
	cfg := wasmtime.NewWasiConfig()
	cfg.InheritStdout()
	store.SetWasi(cfg)

	module, err := wasmtime.NewModuleFromFile(store.Engine, "main.wasm")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to compile module: %w", err))
	}

	linker := wasmtime.NewLinker(store.Engine)
	err = linker.DefineWasi()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to define wasi: %w", err))
	}

	instance, err := linker.Instantiate(store, module)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate module: %w", err))
	}
	call(instance, store, "__main_void")

	call(instance, store, "add", 1, 2)
}

func call(i *wasmtime.Instance, store wasmtime.Storelike, name string, args ...interface{}) {
	fmt.Printf("call function:%s, args: %v\n", name, args)
	runner := i.GetFunc(store, name)
	if runner == nil {
		log.Fatal(fmt.Errorf("failed to find `run` function"))
	}
	got, err := runner.Call(store, args...)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to call `run`: %w", err))
	}
	fmt.Println("result:", got)
}
