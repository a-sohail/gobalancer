package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just got load balanced! %s", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7070", nil))
}
