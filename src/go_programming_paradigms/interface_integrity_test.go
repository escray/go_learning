package go_paradigms

import (
	"fmt"
	"testing"
)

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s* Square)Sides()int {
	return 4
}

func TestIterfaceIntegrity(t *testing.T) {
	s := Square{ len: 5}
	fmt.Printf("%d\n", s.Sides())
}

func TestImplementCheck(t *testing.T) {
	// var _ Shape = (*Square)(nil)
}
