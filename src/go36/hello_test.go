package go_36

import (
	"flag"
	"fmt"
	"testing"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func TestHello(t *testing.T) {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
