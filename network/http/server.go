package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func Welcome(w http.ResponseWriter, res *http.Request) {
	w.Write([]byte("Welcome to go world!"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/welcome", Welcome)
	http.ListenAndServe(":8001", nil)
}
