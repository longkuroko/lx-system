package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello world")
	case "/longxoan":
		fmt.Fprint(w, "Long Xoan")
	default:
		fmt.Fprint(w, "Error")
	}

	fmt.Printf("Handling function with %s request\n", r.Method)

}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Timeout-------------------------\n")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not happend")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)

	server := http.Server{
		Addr:         "localhost:9999",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	server.ListenAndServe()
}
