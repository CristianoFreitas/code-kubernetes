package main

import "net/http"
	
func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World v2!</h1>"))
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8000", nil)
}