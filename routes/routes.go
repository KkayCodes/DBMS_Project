package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "attendance-tracker/controllers"
)

// InitializeRoutes initializes the API routes.
func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    // Student routes
    router.HandleFunc("/students", controllers.GetStudents).Methods("GET") // Get all students
    router.HandleFunc("/students/{id}/attendance", controllers.GetAttendanceForStudent).Methods("GET") // Get attendance for a specific student

    // Professor routes
    router.HandleFunc("/professors", controllers.GetProfessors).Methods("GET") // Get all professors
    router.HandleFunc("/professors/add-course", controllers.AddCourse).Methods("POST") // Add a new course for a professor
    router.HandleFunc("/professors/{id}/attendance", controllers.UpdateAttendance).Methods("POST") // Update attendance for a specific professor

    // Attendance routes
    router.HandleFunc("/attendance", controllers.UpdateAttendance).Methods("POST") // Update attendance status
    router.HandleFunc("/courses/{courseId}/attendance-summary", controllers.GetAttendanceSummary).Methods("GET") // Get attendance summary for a specific course

    return router
}
