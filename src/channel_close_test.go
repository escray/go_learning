package go_learning

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// panic: send on closed channel
		// ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		//for i := 0; i < 10; i++ {
		//  data := <-ch
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func dataReceiveAgain(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i := 0; i < 11; i++ {
			data := <-ch
			fmt.Println(data)
	}
	wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
		var wg sync.WaitGroup
		ch := make(chan int)
		wg.Add(1)
		dataProducer(ch, &wg)
		// receive plus
		// wg.Add(1)
		// dataReceiveAgain(ch, &wg)
		// multiple receive
		wg.Add(1)
		dataReceiver(ch, &wg)
		wg.Add(1)
		dataReceiver(ch, &wg)
		wg.Add(1)
		dataReceiver(ch, &wg)
		wg.Wait()
}

func serviceOne() string {
    time.Sleep(time.Millisecond * 50)
    return "service 1 done"
}

func channelserviceone_2nd() chan string {
    retCh := make(chan string)
    go func() {
        ret := serviceOne()
        fmt.Println("returned result.")
        retCh <- ret
        fmt.Println("service exited.")
    }()
    return retCh
}

func TestSelectAgain(t *testing.T) {
    select {
    case ret := <-channelserviceone_2nd():
        t.Log(ret)
    case <-time.After(time.Millisecond * 49):
        t.Log("time out")
    }
		// time.Sleep(time.Second * 1)
}
