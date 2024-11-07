package models

import (
	"encoding/json"
    "time"
)

type AttendanceRecord struct {
    ID           int       `bson:"id" json:"id"`
    Status       int       `bson:"status" json:"status"`
    CreatedAt    time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
    AttendanceID int       `bson:"attendance_id" json:"attendance_id"`
    StudentID    int       `bson:"student_id" json:"student_id"`
}
