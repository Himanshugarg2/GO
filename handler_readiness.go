package main
import (
	"net/http"	
)
	func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{"status": "ok"})
	}								