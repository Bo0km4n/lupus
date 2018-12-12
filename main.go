package main

import (
	"log"
	"unsafe"

	"github.com/Bo0km4n/lupus/internal/mmap"
	"github.com/Bo0km4n/lupus/internal/patch"
	"github.com/k0kubun/pp"
)

func f1() int {
	return 2
}

func main() {
	// elfFile, _ := elf.Open("/home/vagrant/dev/lupus/example/example")
	// newFuncBytes := elf.GetFuncBytes(elfFile, "main.f2")
	// ptr := *(*uintptr)(unsafe.Pointer(&newFuncBytes))

	// debug
	// mov
	newFuncBytes := []byte{0x49, 0xC7, 0xC2, 0x01, 0x00, 0x00, 0x00}

	pp.Println(len(newFuncBytes))
	newFuncBytes, err := mmap.WriteFuncVal(newFuncBytes)
	pp.Println(*(*uintptr)(unsafe.Pointer(&newFuncBytes)))
	pp.Println(newFuncBytes[0:100])
	if err != nil {
		log.Fatal(err)
	}
	a := f1
	patch.Replace(a, newFuncBytes)
	a()
}
