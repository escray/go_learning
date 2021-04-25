package go_camp

import (
	"fmt"
	"testing"
)

type MyError struct {
	Msg string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

func test() error  {
	return &MyError{"Something happened", "server.go", 42}
}

func TestPrintError(t *testing.T) {
	fmt.Println(&MyError{"Something happened", "server.go", 42})
}

func TestMyError(t *testing.T) {
	err := test()
	switch err := err.(type) {
	case nil:
		// call succeeded, nothing to do
		fmt.Println("call succeeded")
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		// unknown error
		fmt.Println("unknown error")
	}
}
