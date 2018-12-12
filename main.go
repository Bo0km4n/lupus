package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/Bo0km4n/lupus/internal/elf"
	"github.com/Bo0km4n/lupus/internal/mmap"
	"github.com/Bo0km4n/lupus/internal/patch"
	"github.com/k0kubun/pp"
)

func f1() int {
	return 1
}

func main() {
	elfFile, _ := elf.Open("/home/vagrant/go/src/github.com/Bo0km4n/lupus/example/example")
	newFuncBytes := elf.GetFuncBytes(elfFile, "main.f2")
	// ptr := *(*uintptr)(unsafe.Pointer(&newFuncBytes))

	// debug
	// mov
	// newFuncBytes := []byte{0x48, 0xC7, 0xC2, 0x01, 0x00, 0x00, 0x00}

	pp.Println(len(newFuncBytes))
	newFuncBytes, err := mmap.WriteFuncVal(newFuncBytes)
	pp.Println(*(*uintptr)(unsafe.Pointer(&newFuncBytes)))
	pp.Println(newFuncBytes[0:10])
	if err != nil {
		log.Fatal(err)
	}
	a := f1
	patch.Replace(a, newFuncBytes)
	// runtime.Breakpoint()
	fmt.Println(a())
}
