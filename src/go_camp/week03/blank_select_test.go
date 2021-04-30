package go_camp_week03

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestBlankSelect(t *testing.T) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GopherCon SG")
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}

func TestDoTheWorkYourself(t *testing.T) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GopherCon SG")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// search simulates a function that finds a record based
// on a search term. It takes 200ms to perform this work
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

// process is the work for the program. It finds a record
// then prints it.
func process(term string) error {
	record, err := search(term)
	if err != nil {
		return err
	}

	fmt.Println("Receivd:", record)
	return nil
}

type result struct {
	record string
	err	error
}

// launch a goroutine to find the record. Create a result
// from the returned values to send throught the channel
func processWithGoroutine(term string) error  {
	ch := make(chan result)

	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()

	var	ctx context.Context
	// Block waiting to either receive from the goroutine's
	// channel or for the context to the canceled
	select {
	case <- ctx.Done():
		return errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received: ", result.record)
		return nil
	}
}

func TestTwoPorts(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Hello, Qcon!")
	})
	go http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux) // debug
	http.ListenAndServe("0.0.0.0:8080", mux)	// app traffic
}

func serveApp()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, Qcon!")
	})
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug()  {
	http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func TestServeAppAndDebug(t *testing.T) {
	go serveDebug()
	serveApp()
}

