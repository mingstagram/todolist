package models

import "time"

type Users struct {
	ID        int    `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	SnsProvider SnsProvider `json:"sns_provider"`
	SnsId string `json:"sns_id"`
	ProfileImage string `json:"profile_image"`
	Registered_at time.Time `json:"registered_at"`
}