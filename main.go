package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, App1")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
