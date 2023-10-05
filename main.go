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

	runner := instance.GetFunc(store, "__main_void")
	if runner == nil {
		log.Fatal(fmt.Errorf("failed to find `run` function"))
	}
	_, err = runner.Call(store)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to call `run`: %w", err))
	}
}
