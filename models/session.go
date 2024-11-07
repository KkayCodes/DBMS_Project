package models

import (
	"encoding/json"
    "time"
)

type AcademicYear struct {
    ID        int       `bson:"id" json:"id"`
    StartYear time.Time `bson:"start_year" json:"start_year"`
    EndYear   time.Time `bson:"end_year" json:"end_year"`
}
