package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}