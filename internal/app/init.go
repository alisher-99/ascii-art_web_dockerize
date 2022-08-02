package app

import (
	"asciiWeb/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	fs := http.FileServer(http.Dir("ui/"))
	mux.Handle("/ui/", http.StripPrefix("/ui/", fs))

	fmt.Println("Server is listening:8181...go to http://localhost:8181/")
	if err := http.ListenAndServe(":8181", mux); err != nil {
		log.Fatal(err)
	}
}
