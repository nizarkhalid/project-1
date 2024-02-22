package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responeErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5XX error : ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responeWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func responeWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("could not marshel the json respones:  %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
