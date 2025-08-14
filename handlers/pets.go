package handlers

import (
	"adoptme/config"
	"adoptme/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func GetPets(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(context.Background(),
		"SELECT id, name, species, breed, age, description, location, image_url, status FROM pets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var pet models.Pet
		rows.Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Breed, &pet.Age, &pet.Description, &pet.Location, &pet.ImageURL, &pet.Status)
		pets = append(pets, pet)
	}

	json.NewEncoder(w).Encode(pets)
}

func GetPet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var pet models.Pet
	err := config.DB.QueryRow(context.Background(),
		"SELECT id, name, species, breed, age, description, location, image_url, status FROM pets WHERE id=$1", id,
	).Scan(&pet.ID, &pet.Name, &pet.Species, &pet.Breed, &pet.Age, &pet.Description, &pet.Location, &pet.ImageURL, &pet.Status)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pet)
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	json.NewDecoder(r.Body).Decode(&pet)

	pet.ID = uuid.New().String()
	pet.CreatedBy = r.Context().Value("userID").(string)

	_, err := config.DB.Exec(context.Background(),
		`INSERT INTO pets (id, name, species, breed, age, description, location, image_url, status, created_by) 
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		pet.ID, pet.Name, pet.Species, pet.Breed, pet.Age, pet.Description, pet.Location, pet.ImageURL, pet.Status, pet.CreatedBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var pet models.Pet
	json.NewDecoder(r.Body).Decode(&pet)

	_, err := config.DB.Exec(context.Background(),
		`UPDATE pets SET name=$1, species=$2, breed=$3, age=$4, description=$5, location=$6, image_url=$7, status=$8 WHERE id=$9`,
		pet.Name, pet.Species, pet.Breed, pet.Age, pet.Description, pet.Location, pet.ImageURL, pet.Status, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pet updated"})
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := config.DB.Exec(context.Background(), "DELETE FROM pets WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pet deleted"})
}
