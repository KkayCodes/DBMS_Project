package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "attendance-tracker/controllers" // Import your controllers package
    "attendance-tracker/config"      // Import your config package
)

func main() {
    // Initialize MongoDB connection
    config.ConnectDB() // Assuming you have a ConnectDB function in config package

    // Create a new router
    r := mux.NewRouter()

    // API Routes
    r.HandleFunc("/api/students", controllers.GetStudents).Methods("GET") // Get all students
    r.HandleFunc("/api/professors", controllers.GetProfessors).Methods("GET") // Get all professors
    r.HandleFunc("/api/attendance", controllers.UpdateAttendance).Methods("POST") // Update attendance

    // Serve the frontend files
    http.Handle("/", http.FileServer(http.Dir("./frontend/")))
    http.Handle("/api/", r)

    // Start the server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
