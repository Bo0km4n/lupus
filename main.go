package main

import (
	"github.com/k0kubun/pp"

	"C"

	"github.com/Bo0km4n/lupus/internal/elf"
)
import (
	"log"
	"unsafe"

	"github.com/Bo0km4n/lupus/internal/mmap"
)

func f1() int {
	return 1
}

func main() {
	elfFile, _ := elf.Open("/Users/bo0km4n/go/src/github.com/Bo0km4n/lupus/example/example")
	newFuncBytes := elf.GetFuncBytes(elfFile, "main.f2")
	// ptr := *(*uintptr)(unsafe.Pointer(&newFuncBytes))

	pp.Println(newFuncBytes)
	pp.Println(*(*uintptr)(unsafe.Pointer(&newFuncBytes)))
	newFuncBytes, err := mmap.WriteFuncVal(newFuncBytes)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(*(*uintptr)(unsafe.Pointer(&newFuncBytes)))
}
