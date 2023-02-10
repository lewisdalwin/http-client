// Filename: main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080"
	// create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/xml")

	client := http.Client{
		Timeout: 1 * time.Second,  // wait a maximum of 1 sec for a response
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}