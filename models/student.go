// models/student.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student represents a student in the attendance tracker system
type Student struct {
    ID         primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
    Name       string                `bson:"name" json:"name"`
    RollNo     string                `bson:"roll_no" json:"roll_no"`
    Courses    []CourseEnrollment    `bson:"courses" json:"courses"`
    Attendance map[string]Attendance `bson:"attendance" json:"attendance"` // Map of courseID to attendance
}

// CourseEnrollment represents a course that the student is enrolled in
type CourseEnrollment struct {
    CourseID   primitive.ObjectID `bson:"course_id" json:"course_id"`
    CourseName string             `bson:"course_name" json:"course_name"`
}
