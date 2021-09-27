package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Bill")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
