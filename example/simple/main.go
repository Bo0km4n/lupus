package main

import (
	"fmt"
	"log"

	"github.com/Bo0km4n/lupus"
)

func f1() int {
	return 1
}

func main() {
	targetFunc := f1
	if err := lupus.PatchFunction(
		"/home/vagrant/go/src/github.com/Bo0km4n/lupus/example/example",
		"main.f2",
		targetFunc,
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(targetFunc())
}
