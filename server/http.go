package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
