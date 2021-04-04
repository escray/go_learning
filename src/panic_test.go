package go_learning

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong!"))
	// os.Exit(-1)
	// fmt.Println("End")
}
