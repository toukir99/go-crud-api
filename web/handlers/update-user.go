package handlers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/web/model"
	"net/http"
	"strconv"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()
	db := cfg.DB

	// Get the ID from the query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the request body to get the updated user data
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the user in the database, only if fields are provided
	_, err = db.Exec("UPDATE users SET name = COALESCE(NULLIF($1, ''), name), email = COALESCE(NULLIF($2, ''), email) WHERE id = $3", user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Send the updated user in response
	user.ID = id
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
