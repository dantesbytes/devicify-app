package api

import (
	"devicify/backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetDeviceUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID := vars["id"]
	date := r.URL.Query().Get("date")

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	users, err := models.GetUsersLoggedIntoDevice(deviceID, date)
	if err != nil {
		http.Error(w, "Error fetching device users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}