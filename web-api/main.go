//I go to an endpoint, get some JSON
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on :8080")
	http.Handle("/guesses", &guessHandler{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//Response : huh"
type Response struct {
	Message string `json:"message"`
}

type guessHandler struct {
}

//how to register a handler with new endpoint
func (g guessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hi there"))
	resp := &Response{Message: "Hi there"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Panicf("Whoops! %s", err)
	}
}
