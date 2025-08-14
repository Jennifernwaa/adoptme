package models

import "time"

type Favorite struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PetID     string    `json:"pet_id"`
	CreatedAt time.Time `json:"created_at"`
}
