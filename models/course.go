package models

import (
	"encoding/json"
    "time"
)

type Course struct {
    ID        int       `bson:"id" json:"id"`
    Name      string    `bson:"name" json:"name"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
