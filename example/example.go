package main

import "fmt"

func f1() int {
	fmt.Println("hoge")
	return 1
}

func f2() int {
	return 2
}

func main() {
	f2()
}
