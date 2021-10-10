package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	// "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// greeting := os.Getenv("GREET")
	// fmt.Fprintf(w, "%s, 世界\n", greeting)
	url := "http://app2:7777"
	resp, err := http.Get(url)
	if err != nil {
      fmt.Fprintf(w, "ERROR READING APP2\n")
	  panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Fprintf(w, "ERROR READING resp.Body\n")
      panic(err)
    }
    // show the HTML code as a string %s
    fmt.Fprintf(w, "%s\n", html)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6666", nil))
}
