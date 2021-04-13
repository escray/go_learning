package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello World!")
		fmt.Println("hello world")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		str := "hello"
		w.Write([]byte(str))
	})

	http.HandleFunc("/sub/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " leafTree sub")
	})

	http.HandleFunc("/sub/node", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " leafNode sub")
	})

	http.ListenAndServe(":8081", nil)
}
