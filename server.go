package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Main
func main() {
	http.HandleFunc("/secret", SecretHandler)
	http.HandleFunc("/configmap", ConfigMapHandler)
	http.HandleFunc("/", HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, %s! You are %s years old.\n", name, age)
}

func ConfigMapHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	fmt.Fprintf(w, "My Family: %s", string(data))
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}
