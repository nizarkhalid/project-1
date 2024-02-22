package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responeErr(w, 500, "something went wrong")
}
