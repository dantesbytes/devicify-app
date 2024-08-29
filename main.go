package main

import (
	"devicify/backend/api"
	"devicify/backend/database"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	r.HandleFunc("/", landingPageHandler)

	r.HandleFunc("/api/device/{id}/users", api.GetDeviceUsers).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}