package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingPongHandler) //Create Handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingPongHandler(
	w http.ResponseWriter, //interface --> Egress
	r *http.Request, // pointer concret type --> Ingress
) {
	resp := struct {
		Message string `json: "message"`
	}{
		Message: "pong",
	}
	json.NewEncoder(w).Encode(&resp)
}
