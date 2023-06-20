package models

import "time"

type Prospect struct {
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	IPAddress string    `json:"ip_address" bson:"ip_address"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
