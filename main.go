package main

import (
	"log"
	"net/http"
	// "os"
	// "path/filepath"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src/")))
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
