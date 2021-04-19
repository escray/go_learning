package effective_go

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

type IntCounter int

func (ctr *IntCounter) ServeHTTPByPointer(w http.ResponseWriter, r *http.Request)  {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

type Chan chan *http.Request

func (ch Chan) ServeHTTPByChannel(w http.ResponseWriter, r *http.Request) {
	ch <- r
	fmt.Fprintf(w, "notification sent")
}

func ArgServer()  {
	fmt.Println(os.Args)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f HandlerFunc) ServeHTTPByHandleFunc(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

func ArgServerByHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func TestServeHttp(t *testing.T) {
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	http.Handle("/args", http.HandlerFunc(ArgServerByHandleFunc))
}
