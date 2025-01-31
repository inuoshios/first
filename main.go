package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubURL       string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Email:           "inuoshios@gmail.com",
		CurrentDatetime: time.Now().Format(time.RFC3339),
		GithubURL:       "https://github.com/inuoshios/first",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("runnning your server...")
	http.HandleFunc("GET /", handler)

	http.ListenAndServe(":8080", nil)
}
