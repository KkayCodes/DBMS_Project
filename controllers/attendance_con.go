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

// UpdateAttendance updates attendance status for a student in a specific course.
func UpdateAttendance(w http.ResponseWriter, r *http.Request) {
    var attendanceInfo struct {
        CourseID  string `json:"course_id"`
        StudentID string `json:"student_id"`
        Status    string `json:"status"` // Status should be "green", "yellow", or "red"
    }
    
    // Decode request body
    if err := json.NewDecoder(r.Body).Decode(&attendanceInfo); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }

    // Update attendance status in the database
    attendanceCollection := config.Client.Database("attendanceDB").Collection("attendance")
    _, err := attendanceCollection.UpdateOne(
        context.TODO(),
        bson.M{"course_id": attendanceInfo.CourseID, "student_id": attendanceInfo.StudentID},
        bson.M{"$set": bson.M{"status": attendanceInfo.Status}},
    )
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }

    // Successful response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(utils.SuccessResponse("Attendance updated successfully!"))
}

// GetAttendanceSummary retrieves the attendance summary for a specific course.
func GetAttendanceSummary(w http.ResponseWriter, r *http.Request) {
    courseID := r.URL.Query().Get("course_id")
    var attendanceRecords []models.Attendance
    collection := config.Client.Database("attendanceDB").Collection("attendance")
    
    // Find attendance records for the course
    cursor, err := collection.Find(context.TODO(), bson.M{"course_id": courseID})
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
        attendanceRecords = append(attendanceRecords, record)
    }

    // Successful response with attendance records
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(attendanceRecords)
}
