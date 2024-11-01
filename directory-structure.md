attendance-tracker/
├── go.mod                    # Go module file
├── go.sum                    # Dependency lock file
├── main.go                   # Main application entry point
├── config/
│   └── config.go             # Configuration setup (e.g., MongoDB connection, Firebase setup)
├── models/
│   ├── student.go            # Student schema/model
│   ├── professor.go          # Professor schema/model
│   └── attendance.go         # Attendance schema/model
├── controllers/
│   ├── student_controller.go # Student-specific logic and handlers
│   ├── professor_controller.go # Professor-specific logic and handlers
│   └── attendance_controller.go # Attendance logic and handlers
├── routes/
│   └── routes.go             # API routes
├── utils/
│   └── utils.go              # Utility functions (e.g., helper functions for formatting dates or handling errors)
├── frontend/                 # Frontend (optional for now if focusing on API)
│   ├── student/              # Separate student frontend (e.g., React app)
│   └── professor/            # Separate professor frontend
└── README.md                 # Project documentation



frontend/
├── src/
│   ├── components/
│   │   ├── Navbar.js
│   │   ├── ProfessorDashboard.js
│   │   ├── AddCourseForm.js
│   │   ├── StudentList.js
│   │   └── AttendanceSummary.js
│   ├── App.js
│   ├── index.js
│   ├── App.css
│   └── index.css
└── package.json
