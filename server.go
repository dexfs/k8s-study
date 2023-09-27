package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

// Main
func main() {
	http.HandleFunc("/healthz", HealthzHandler)
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

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 8 || duration.Seconds() > 30 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}
