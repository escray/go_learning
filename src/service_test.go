package go_learning

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	// retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

// classic channel
// working on something else
// returned result.
// Task is done.
// service exited.
// Done

// buffer channel
// working on something else
// returned result.
// service exited.
// Task is done.
// Done

func TestBufferChannel(t *testing.T) {
	// c := make(chan int)
	// panic: test timed out after 30s
	// fatal error: all goroutines are asleep - deadlock!
	c := make(chan int, 1)
	c <- 12
	time.Sleep( 2 * time.Second)
	ch := <-c
	fmt.Println(ch)
}

func TestDeadLock(t *testing.T) {
	arr:=[]int{1,2,3,4,5}

	c:=make(chan int,5)
	fin:=make(chan int,3)
	cl:=make(chan int,1)

	defer close(fin)
	defer close(cl)
	defer close(c)

	go func() {
    for _,ch:=range arr{
        c<-ch
    }
	}()

	for i:=0;i<3;i++{
    go func() {
        for{
            classId:=<-c
            time.Sleep(time.Second * 1)
            fin<-classId
        }
    }()
	}

	go func() {
    for classId:=range fin{
        fmt.Printf("classId %d finished \n",classId)
    }
    cl<-1
	}()

	<-cl

	fmt.Println("all done")
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
