package main

import (
	"fmt"
	"reflect"
)

type Father struct {
	son *Son
}

type Son struct {
	name string
}

func (s *Son) Say()  {
	fmt.Println(s.name)
	fmt.Println(reflect.TypeOf(s).String())
	fmt.Println(reflect.TypeOf(*s).String())
}

func main() {
	f := &Father{}
	s := new(Son)
	s.name = "dog"
	ss := &Son{
		name: "cat",
	}
	f.son = s
	// f.son.name = "ray"
	f.son.Say()
	f.son = ss
	f.son.Say()
	fmt.Println("done")
}
