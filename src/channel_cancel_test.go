package go_learning

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 50; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancel")
		}(i, cancelChan)
	}
	// time.Sleep(time.Millisecond * 10)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}

func TestA(t *testing.T) {
	ch := make(chan string)
	// ch <- "Hi"
	select {
		case ret:= <-ch:
			fmt.Println("channel receives", ret)
		default:
			t.Log("default")
	}
}
