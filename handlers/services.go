package handlers

import (
	"fmt"
	"net/http"
)

func Service1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Response from Service 1")
}

func Service2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Response from Service 2")
}
