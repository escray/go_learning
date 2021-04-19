package effective_go

import (
	"fmt"
	"sort"
	"testing"
)

func TestChannel(t *testing.T) {
	c := make(chan int)
	list := make(Sequence, 0, 10)
	go func() {
		sort.Sort(list)
		c <- 1
	}()
	doSomethingForAWhile()
	<- c
}

func doSomethingForAWhile()  {

}

func process(r *Request) {

}

type Request struct {
	args []int
	f func([]int) int
	resultChan chan int
}


var MaxOutstanding int = 10
var sem = make(chan int, MaxOutstanding)

func handle(r *Request)  {
	sem <- 1
	process(r)
	<- sem
}

func Serve(queue chan *Request)  {
	for {
		req := <- queue
		go handle(req)
	}
}

func handleAgain(queue chan *Request)  {
	for r := range queue {
		process(r)
	}
}

func ServeAgain(clientRequest chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handleAgain(clientRequest)
	}
	<- quit
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func TestSum(t *testing.T) {
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	var clientRequest chan *Request
	clientRequest <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

func handleTriple(queue chan *Request)  {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}
