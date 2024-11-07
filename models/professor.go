package models

import (
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Professor struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ProfessorID int                `bson:"id" json:"id"`
    Password    string             `bson:"password" json:"password"`
    LastLogin   *time.Time         `bson:"last_login,omitempty" json:"last_login,omitempty"`
    IsSuperuser bool               `bson:"is_superuser" json:"is_superuser"`
    FirstName   string             `bson:"first_name" json:"first_name"`
    LastName    string             `bson:"last_name" json:"last_name"`
    IsStaff     bool               `bson:"is_staff" json:"is_staff"`
    IsActive    bool               `bson:"is_active" json:"is_active"`
    DateJoined  time.Time          `bson:"date_joined" json:"date_joined"`
    Email       string             `bson:"email" json:"email"`
    UserType    string             `bson:"user_type" json:"user_type"`
    Gender      string             `bson:"gender" json:"gender"`
    ProfilePic  string             `bson:"profile_pic" json:"profile_pic"`
    Address     string             `bson:"address" json:"address"`
    FCMToken    string             `bson:"fcm_token" json:"fcm_token"`
    CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
    Courses     []Course           `bson:"courses" json:"courses"`
}

type Course struct {
    CourseID   primitive.ObjectID   `bson:"course_id,omitempty" json:"course_id"`
    CourseName string               `bson:"course_name" json:"course_name"`
    Students   []primitive.ObjectID `bson:"students" json:"students"`
}
