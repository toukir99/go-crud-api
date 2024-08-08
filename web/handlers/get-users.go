package handlers

import (
	"go-crud-api/config"
	"log"
	"net/http"
	"encoding/json"
	"go-crud-api/web/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	cfg := config.GetConfig()
	db := cfg.DB
	
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}	
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}