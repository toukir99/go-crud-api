package handlers

import (
	"go-crud-api/config"
	"net/http"
	"encoding/json"
	"go-crud-api/web/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	cfg := config.GetConfig()
	db := cfg.DB

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user);  err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	
	err := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}