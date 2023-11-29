package models

import "github.com/google/uuid"

// Customer struct for the database model
type Customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     int       `json:"phone,omitempty"`
	Contacted bool      `json:"contacted,omitempty"`
}
