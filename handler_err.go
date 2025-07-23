package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondingWithError(w, 400, "Something Went Wrong")
}
