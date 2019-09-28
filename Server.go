package main

import "net/http"

// Server : A simple web server
func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error)
	}
}
