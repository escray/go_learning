package effective_go

import (
	"fmt"
	"testing"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a()  {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b()  {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func TestDeferExploit(t *testing.T) {
	b()
}

func TestDeferLIFO(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
