package main

import (
	"log"
	"net/http"
)

// Main
func main() {
	http.HandleFunc("/", HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, There!</h1>"))
}
