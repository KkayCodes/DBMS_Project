package main

import (
    "html/template"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "your_project_name/config" // Replace with your actual module name
)

func main() {
    // Initialize configuration (MongoDB connection, Firebase setup)
    err := config.Initialize() // Assuming you have a function to initialize your configs
    if err != nil {
        log.Fatalf("Could not initialize configuration: %v", err)
    }

    // Create a new router
    r := mux.NewRouter()

    // Serve static files
    r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
    
    // HTML Templates
    r.HandleFunc("/", renderTemplate("login.html")).Methods("GET")
    r.HandleFunc("/login/professor", professorLogin).Methods("GET")
    r.HandleFunc("/login/student", studentLogin).Methods("GET")
    r.HandleFunc("/professor", renderTemplate("professor_dashboard.html")).Methods("GET")
    r.HandleFunc("/student", renderTemplate("student_dashboard.html")).Methods("GET")
    r.HandleFunc("/logout", logout).Methods("GET")

    // Start the server
    log.Println("Starting server on :8080...")
    err = http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}

func renderTemplate(filename string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/" + filename)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func professorLogin(w http.ResponseWriter, r *http.Request) {
    // Add Firebase authentication logic for professor
    // Redirect to professor dashboard if authenticated
    http.Redirect(w, r, "/professor", http.StatusFound)
}

func studentLogin(w http.ResponseWriter, r *http.Request) {
    // Add Firebase authentication logic for student
    // Redirect to student dashboard if authenticated
    http.Redirect(w, r, "/student", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
    // Handle logout logic (clear session/cookies)
    http.Redirect(w, r, "/", http.StatusFound)
}
