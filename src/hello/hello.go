package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello, world.")
	fmt.Println(morestrings.ReverseRunes("Shawshank"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}


