package go_camp_week03

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func serveAppAgain()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func serveDebugAgain()  {
	if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func test_serve_Debug()  {
	go serveDebug()
	go serveApp()
	select {}
}

func serve(addr string, handler http.Handler, stop <- chan struct{}) error {
	s := http.Server {
		Addr: addr,
		Handler: handler,
	}

	go func() {
		<- stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveDebugChan(stop chan struct{}) error {
	return nil
}

func serveAppChan(stop chan struct{}) error {
	return nil
}

func TestServeApp(t *testing.T) {
	done := make(chan error, 2)
	stop := make(chan struct {})
	go func() {
		done <- serveDebugChan(stop)
	}()
	go func() {
		done <- serveAppChan(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
