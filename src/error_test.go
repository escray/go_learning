package go_learning

import (
	"errors"
	"testing"
)



func GetFibonacci(n int) ([]int, error){
	if n < 2 {
		return nil, errors.New("n should be less than 2")
	}

	if n > 100 {
		return nil, errors.New("n should be not larger than 100")
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

	if v, err:=GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
