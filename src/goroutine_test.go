package go_learning

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 值传递，每个协程所拥有的变量地址不一样
		go func (i int){
			fmt.Println(i)
		}(i)

		// go func() {
		// 	// 10
		// 	// 共享变量，竞争，锁
		// 	fmt.Println(i)
		// }()
	}

	time.Sleep(time.Millisecond*50)
}

var wg sync.WaitGroup

func TestIncreaseCounter(t *testing.T) {
  counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
  wg.Wait()
	t.Log(("Counter = " + strconv.Itoa(counter)))
}

func TestSchedule(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func(){
		t.Log("0")
	}()
	for {
		t.Log("1")
	}
}

func doSomething(i int) {
	fmt.Print(i)
}

func TestScheduleGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			doSomething(0)
		}
	}()

	for {
		doSomething(1)
	}
}
