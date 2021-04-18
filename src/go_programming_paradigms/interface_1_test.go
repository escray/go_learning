package go_paradigms

import (
	"fmt"
	"testing"
)

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr()  {
	fmt.Println(c.Name)
}

func (c City) PrintStr()  {
	fmt.Println(c.Name)
}

func TestInterface01(t *testing.T) {
	c1 := Country {"China"}
	c2 := City {"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}

type Stringable interface {
	ToString() string
}

func (c Country)ToString() string {
	return "Country = " + c.Name
}

func (c City)ToString() string {
	return "City = " + c.Name
}

func PrintStr(p Stringable)  {
	fmt.Println(p.ToString())
}

func TestStringable(t *testing.T) {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
