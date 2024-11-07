package models

import (
	"encoding/json"
    "time"
)

type Subject struct {
    ID        int       `bson:"id" json:"id"`
    Name      string    `bson:"name" json:"name"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    CourseID  int       `bson:"course_id" json:"course_id"`
    StaffID   int       `bson:"staff_id" json:"staff_id"`
}
