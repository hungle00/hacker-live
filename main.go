package main

import (
	"log"
	"net/http"
	"os"

	"huy.rocks/hackerlive/api"
)

// This file isn't needed for deploying on Vercel, it's just for local development.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/feed", api.FeedHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
