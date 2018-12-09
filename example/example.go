package main

import "fmt"

func f1() int {
	return 1
}

func f2() {
	fmt.Println("f2")
}

type S struct {
	Callback func()
}

func (s *S) Sfunc() int {
	return 0
}

func main() {
	s := &S{Callback: f2}
	s.Callback()
}
