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

// GetProfessors retrieves all professors.
func GetProfessors(w http.ResponseWriter, r *http.Request) {
    var professors []models.Professor
    collection := config.Client.Database("attendanceDB").Collection("professors")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var professor models.Professor
        if err := cursor.Decode(&professor); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
            return
        }
        professors = append(professors, professor)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(professors)
}

// AddCourse adds a course for a professor and updates all students.
func AddCourse(w http.ResponseWriter, r *http.Request) {
    var courseInfo struct {
        CourseID   string   `json:"course_id"`
        CourseName string   `json:"course_name"`
        StudentIDs []string `json:"student_ids"`
    }

    // Decode request body
    if err := json.NewDecoder(r.Body).Decode(&courseInfo); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }

    // Add course to the professor
    collection := config.Client.Database("attendanceDB").Collection("professors")
    _, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"id": courseInfo.CourseID},
        bson.M{"$addToSet": bson.M{"courses": courseInfo.CourseName}},
    )
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
        return
    }

    // Update each student's courses
    studentCollection := config.Client.Database("attendanceDB").Collection("students")
    for _, studentID := range courseInfo.StudentIDs {
        _, err := studentCollection.UpdateOne(
            context.TODO(),
            bson.M{"id": studentID},
            bson.M{"$addToSet": bson.M{"courses": courseInfo.CourseName}},
        )
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
            return
        }
    }

    // Successful response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(utils.SuccessResponse("Course added successfully!"))
}
