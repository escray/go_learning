package go_learning

import (
	"fmt"
	"math/rand"
	"os"
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

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，后进先出）
func fdefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

// 变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的
func fdefer2() (ret int) {
	defer func() {
		ret ++
	}()
	return 1
}

func TestDeferAgain(t *testing.T) {
	fdefer()
	fmt.Println(fdefer2())
}

// func TestDefer(t *testing.T) {
// 	defer Clear()
// 	fmt.Println("Start")
// 	panic("err")
// 	// can't access
// 	// fmt.Println("End")
// }

func TestDeferOSExit(t *testing.T) {
	defer func() {
		fmt.Println("clear resources.")
		t.Log("test log")
	}()

	fmt.Println("Started")
	t.Log("test start")
	os.Exit(0)
}

func GetFunc() func() {
	fmt.Print("[outside]")
	return func() {
		fmt.Print("[inside]")
	}
}

// [outside][here][inside]
func TestDeferGetFunc(t *testing.T) {
	defer GetFunc()()
	fmt.Print("[here]")

}

// [there][outside]
func TestDeferGetFuncAgain(t *testing.T) {
  defer GetFunc()
	fmt.Print("[there]")
}

func ClearAgain(par1 int) {
  fmt.Println("3. Clear resources!", par1)
}

func FuncDefer() {
  fmt.Println("2. 被 defer 的函数里面还有被 defer 的化，会被先执行!")
  defer ClearAgain(2)
}

func TestDefer(t *testing.T) {
  defer ClearAgain(1)
  defer FuncDefer()
	defer ClearAgain(3)
  fmt.Println("1. Start")
  panic("err")
}
