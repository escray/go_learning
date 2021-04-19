package effective_go

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

type CounterHTTP struct {
	n int
}

func (ctr *CounterHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

type IntCounterHTTP int

func (ctr *IntCounterHTTP) ServeHTTPByPointer(w http.ResponseWriter, r *http.Request)  {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

type ChanHTTP chan *http.Request

func (ch ChanHTTP) ServeHTTPByChannel(w http.ResponseWriter, r *http.Request) {
	ch <- r
	fmt.Fprintf(w, "notification sent")
}

func ArgServerHTTP()  {
	fmt.Println(os.Args)
}

type HandlerFuncHTTP func(w http.ResponseWriter, r *http.Request)

func (f HandlerFuncHTTP) ServeHTTPByHandleFunc(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

func ArgServerHTTPByHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func TestServeHttpAgain(t *testing.T) {
	ctr := new(CounterHTTP)
	http.Handle("/counter", ctr)

	http.Handle("/args", http.HandlerFunc(ArgServerHTTPByHandleFunc))
}
