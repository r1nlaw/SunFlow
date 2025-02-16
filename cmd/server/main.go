package main

import (
	"fmt"
	"log"
	"net/http"
	"sunflow/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handlers)
	http.HandleFunc("/sunlight", handlers.SunlightHandler)

	port := ":8080"
	fmt.Println("Server start", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
