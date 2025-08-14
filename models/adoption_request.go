package models

import "time"

type AdoptionRequest struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PetID     string    `json:"pet_id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
