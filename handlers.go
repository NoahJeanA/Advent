package handlers

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo! Du bist auf dem Webserver gelandet.")
}