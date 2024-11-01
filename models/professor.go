// models/professor.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Professor represents a professor in the attendance tracker system
type Professor struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name"`
    Courses  []Course           `bson:"courses" json:"courses"`
}

// Course represents a course taught by a professor
type Course struct {
    CourseID   primitive.ObjectID   `bson:"course_id,omitempty" json:"course_id"`
    CourseName string               `bson:"course_name" json:"course_name"`
    Students   []primitive.ObjectID `bson:"students" json:"students"` // Array of student IDs
}
