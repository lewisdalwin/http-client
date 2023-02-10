package main

import (
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// simulate a delayed response from the server
	time.Sleep(10 * time.Second)
	
	w.Write([]byte("hello from root.\n"))
}
func main() {
	http.HandleFunc("/", rootHandler)
	log.Print("starting server on port :8080")
	http.ListenAndServe(":8080", nil)
}