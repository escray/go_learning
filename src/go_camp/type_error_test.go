package go_camp

import (
	"errors"
	"fmt"
	"testing"
)

// Create a named type for our new error type
type errorString string

// Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
// 标准库
var ErrStructType = errors.New("EOF")

func TestNamedTypeError(t *testing.T) {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}

type errorStringStruct struct {
	s string
}

func (e errorStringStruct) Error() string  {
	return e.s
}

func NewError(text string) error {
	return errorStringStruct{text}
}

var ErrType = NewError("EOF")

func TestEOFError(t *testing.T) {
	if ErrType == NewError("EOF") {
		fmt.Println("Error: ", ErrType)
	}
}
