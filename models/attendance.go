package models

import (
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Attendance struct {
    ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    AttendanceID      int                `bson:"id" json:"id"`
    Date              string             `bson:"date" json:"date"`
    CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
    SessionID         int                `bson:"session_id" json:"session_id"`
    SubjectID         int                `bson:"subject_id" json:"subject_id"`
    TotalClasses      int                `bson:"total_classes" json:"total_classes"`
    ClassesAttended   int                `bson:"classes_attended" json:"classes_attended"`
    AttendancePercent float64            `bson:"attendance_percent" json:"attendance_percent"`
}

func (a *Attendance) UpdateAttendance() {
    if a.TotalClasses > 0 {
        a.AttendancePercent = (float64(a.ClassesAttended) / float64(a.TotalClasses)) * 100
    } else {
        a.AttendancePercent = 0
    }
}
