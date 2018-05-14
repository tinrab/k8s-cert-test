package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	forever := make(chan struct{})
	go func() {
		http.ListenAndServe(":8443", nil)
	}()

	<-forever
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
