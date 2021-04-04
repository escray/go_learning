package go_learning

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2");
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error){
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
  t.Log(GetFibonacci(10))
	t.Log(GetFibonacci(-10))

	if v, err:=GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacciAgain(str string){
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
		  fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacciAnother(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}

func TestNewEqual(t *testing.T) {
  // Different allocations should not be equal.
  if errors.New("abc") == errors.New("abc") {
    t.Errorf(`New("abc") == New("abc")`)
  }
  if errors.New("abc") == errors.New("xyz") {
    t.Errorf(`New("abc") == New("xyz")`)
  }

  // Same allocation should be equal to itself (not crash).
  err := errors.New("jkl")
  if err != err {
    t.Errorf(`err != err`)
  }
}
