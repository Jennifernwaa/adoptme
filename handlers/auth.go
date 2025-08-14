package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"adoptme/config"
	"adoptme/models"
	"adoptme/utils"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.ID = uuid.New().String()

	if user.Name == "" {
		// Return an error to the user
		http.Error(w, "Name is a required field", http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec(context.Background(),
		"INSERT INTO users (id, email, password, name) VALUES ($1, $2, $3, $4)",
		user.ID, user.Email, string(hashedPassword), user.Name,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	var dbUser models.User
	err := config.DB.QueryRow(context.Background(),
		"SELECT id, password FROM users WHERE email=$1", input.Email,
	).Scan(&dbUser.ID, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(dbUser.ID, dbUser.Role)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
