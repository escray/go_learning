package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func (op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println ("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a,b := returnMultiValues()
	t.Log(a, b)

	c,_ := returnMultiValues()
	t.Log(c)

	rand.Seed(time.Now().UnixNano())
	t.Log(rand.Intn(10), rand.Intn(20))

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarPara(t *testing.T) {
	t.Log(sum(1, 3, 5, 7))
	t.Log(sum(1, 2, 3, 4, 5))
}

func Clear(){
	fmt.Println("Clear resources. ")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
	// can't access
	// fmt.Println("End")
}
