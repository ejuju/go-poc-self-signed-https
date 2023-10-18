package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting server on port 8080")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", http.HandlerFunc(handleHTTP))
	if err != nil {
		panic(err)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<!DOCTYPE html><h1>Hello world</h1>"))
}
