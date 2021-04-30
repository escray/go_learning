package go_camp

import (
	"errors"
	"fmt"
	"testing"
)

func handle() (int, error) {
	return 1, nil
}

func TestErrorHandle(t *testing.T) {
	i, err := handle()
	if err != nil {
		return
	}
	fmt.Println(i)
}

// Positive returns true if the number is positive, false if it is negative
func Positive(n int) bool {
	return n > -1
}

func Check(n int)  {
	if Positive(n) {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// Positive returns true if the number is positive, false if it is negative
// The second return value indicates if the result is valid
// which in the case of n == 0, is not valid
func PositiveAgain(n int) (bool, bool) {
	if n == 0 {
		return false, false
	}
	return n > -1, true
}

func CheckAgain(n int) {
	pos, ok := PositiveAgain(n)
	if !ok {
		fmt.Println(n, "is neither")
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// Positive returns true if the number is positive,
// false if it is negative
func PositiveWithError(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("underfined")
	}
	return n > -1, nil
}

func CheckWithError(n int)  {
	pos, err := PositiveWithError(n)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// 非常不建议，返回值有二义性
// If the result not nil, the result is true if the number is positive,
// false if it is negative
func PositivePointer(n int) *bool {
	if n == 0 {
		return nil
	}
	r := n > -1
	return &r
}

func CheckPointer(n int)  {
	pos := PositivePointer(n)
	if pos == nil {
		fmt.Println(n, "is neither")
		return
	}
	if *pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// 强烈不建议
// Positive returns true if the number is positive
// false if it is negative
// In the case that n is 0, Positive will panic
func PositiveWithPanic(n int) bool {
	if n == 0 {
		panic("undefined")
	}
	return n > -1
}

func CheckWithPanic(n int)  {
	defer func() {
		if recover() != nil {
			fmt.Println("is neither")
		}
	}()

	if PositiveWithPanic(n) {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func TestPositiveCheck(t *testing.T) {
	Check(1)
	Check(0)
	Check(-1)

	CheckAgain(1)
	CheckAgain(0)
	CheckAgain(-1)

	CheckWithError(1)
	CheckWithError(0)
	CheckWithError(-1)

	CheckPointer(1)
	CheckPointer(0)
	CheckPointer(-1)

	CheckWithPanic(1)
	CheckWithPanic(0)
	CheckWithPanic(-1)
}

