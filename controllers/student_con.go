package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "attendance-tracker/config"
    "attendance-tracker/models"
    "attendance-tracker/utils"
    "go.mongodb.org/mongo-driver/bson"
)

// GetStudents retrieves all students.
func GetStudents(w http.ResponseWriter, r *http.Request) {
    var students []models.Student
    collection := config.Client.Database("attendanceDB").Collection("students")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var student models.Student
        if err := cursor.Decode(&student); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
            return
        }
        students = append(students, student)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}

// GetAttendance retrieves attendance for a specific student.
func GetAttendance(w http.ResponseWriter, r *http.Request) {
    studentID := r.URL.Query().Get("student_id")
    if studentID == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(utils.ErrorResponse("Missing student_id query parameter"))
        return
    }

    var attendance []models.Attendance
    collection := config.Client.Database("attendanceDB").Collection("attendance")
    cursor, err := collection.Find(context.TODO(), bson.M{"student_id": studentID})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var record models.Attendance
        if err := cursor.Decode(&record); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
            return
        }
        attendance = append(attendance, record)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(attendance)
}
