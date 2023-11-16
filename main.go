package main

import (
	"fmt"
	"net/http"
	"github.com/yourusername/yourproject/handlers" // Verwende den richtigen Pfad zu deinem Projekt
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}