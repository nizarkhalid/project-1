package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responeWithJSON(w, 200, struct{}{})
}
