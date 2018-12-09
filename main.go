package main

import (
	"github.com/Bo0km4n/lupus/internal/elf"
	"github.com/k0kubun/pp"
)

func main() {
	elfFile, _ := elf.Open("/Users/bo0km4n/go/src/github.com/Bo0km4n/lupus/example/example")
	pp.Println(elf.GetFuncBytes(elfFile, "main.f2"))
}
