package main

import (
	"log"
	"net/http"
	"pencarian_film/web"
)

func main() {
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/search", web.SearchHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
