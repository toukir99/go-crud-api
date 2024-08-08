package handlers

import (
	"go-crud-api/config"
	"net/http"
	"encoding/json"
	"go-crud-api/web/model"
	"strconv"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()
	db := cfg.DB

	idStr := r.URL.Query().Get("id")
	//log.Printf("id: %s", idStr)
	if idStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
    	return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted!")
}
