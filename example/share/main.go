package main

import (
	"fmt"
	"log"

	"github.com/Bo0km4n/lupus"
)

func f1() int {
	fmt.Println("hello")
	return 1
}

func main() {
	targetFunc := f1
	if err := lupus.PatchFunction(
		"/home/vagrant/go/src/github.com/Bo0km4n/lupus/example/share/patched.bin",
		"main.newFunc",
		targetFunc,
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(targetFunc())
	newFunc()
}
