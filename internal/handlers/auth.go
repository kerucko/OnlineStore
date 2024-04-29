package handlers

import (
	"log"
	"net/http"
)

func StartHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("I am here")
		http.ServeFile(w, r, "public/index.html")
	}
}
