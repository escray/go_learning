package go_paradigms

import (
	"fmt"
	"testing"
)

type WithName struct {
	Name string
}

type Countr struct {
	WithName
}

type Cit struct {
	WithName
}

type Printabl interface {
	PrintStr()
}

func (w WithName)PrintStr()  {
	fmt.Println(w.Name)
}

func TestStructEmbedded(t *testing.T) {
	c1 := Countr { WithName{ "China"} }
	c2 := Cit { WithName{ "beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}
