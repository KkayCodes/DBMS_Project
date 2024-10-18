# DBMS_Project

## Kkay

To create a **centralized attendance tracker** for a classroom using **Go (Golang)**, **MongoDB**, and **Firebase** with a **clean white and blue light mode themed frontend**, here's an outline of the solution, broken down into **Frontend**, **Backend**, and **Database** layers. The application will include roles for both **teachers** and **students**, with the teacher managing attendance and the student viewing their attendance stats in real-time.

### Tech Stack:
- **Backend**: Go (Golang)
- **Database**: MongoDB
- **Authentication & Realtime Data**: Firebase
- **Frontend**: HTML, CSS, JavaScript (React.js or plain JS for simplicity)
- **Hosting**: Firebase Hosting (for the frontend), a Go server for the backend

---

### 1. **High-Level Application Architecture**

- **Frontend (Client-Side)**:
  - Two interfaces: **Teacher Dashboard** and **Student Dashboard**.
  - The **Teacher Dashboard** allows teachers to:
    - Add student details (name, roll number).
    - Mark attendance (present, absent, leave).
  - The **Student Dashboard** allows students to:
    - View their attendance record (with color-coding: green, orange, or red).
    - Get real-time updates when attendance is marked.
  
- **Backend (Server-Side)**:
  - Golang-based REST API for interaction with MongoDB.
  - Firebase used for **Authentication** (for teachers and students).
  - Firebase **Firestore** can also be used for real-time updates to student apps, though MongoDB can be the main database for structured attendance data.
  - REST endpoints for CRUD operations on student data and attendance.

---

### 2. **Detailed Features**

#### **Frontend (UI/UX Design)**

**Tech**: HTML, CSS (with a light blue and white color theme), and JavaScript (React.js if you prefer component-driven UI).

##### **Teacher Dashboard**:
- **Add Student**:
  - Form fields to enter student `name`, `roll number`.
  - Submit button to add the student to MongoDB.
  
- **Attendance Recording**:
  - List of all students with roll numbers.
  - Each student has buttons to mark attendance as:
    - **Present**
    - **Absent**
    - **Leave**
  - Marking attendance updates MongoDB and pushes the new data to the student's Firebase app.

- **Attendance Summary**:
  - A table showing the attendance records of all students.
  - Color-coded percentage based on the attendance threshold:
    - Green: 75% and above.
    - Orange: 65-74%.
    - Red: Below 65%.

##### **Student Dashboard**:
- **Attendance Stats**:
  - Display attendance stats in a table:
    - Total days attended, total days, and percentage.
    - Color-coded stats based on the attendance (green, orange, red).
  
- **Real-time Updates**:
  - Display the live attendance updates using Firebase (when the teacher marks attendance, it updates immediately on the student app).

##### **Themes**:
- Use CSS for styling with primary colors of **white** and **light blue** to achieve a clean, modern look.

---

### 3. **Backend (API and Business Logic)**

#### **Tech**: Go (Golang)

##### **Authentication (Firebase)**:
- Use Firebase Authentication to authenticate users (teachers and students).
- Roles: Teachers and Students have different permissions.

##### **API Endpoints**:

- **POST /students**: Add a new student (only accessible by the teacher).
  - Request Body: `{ "name": "Student Name", "roll_number": "123" }`
  
- **GET /students**: Fetch the list of all students (only accessible by the teacher).

- **POST /attendance**: Mark attendance for students (only accessible by the teacher).
  - Request Body: `{ "roll_number": "123", "status": "present" }`
  
- **GET /attendance/{roll_number}**: Fetch attendance of a specific student (accessible by the student and teacher).
  
- **GET /attendance-summary**: Fetch the attendance summary for all students (accessible by the teacher).

##### **Attendance Logic**:
- Every time attendance is marked, the **percentage of attendance** should be recalculated.
- MongoDB stores attendance records, and the percentage is computed using the total classes attended vs. total classes held.
- The backend can run a cron job or trigger for recalculating attendance if required periodically.

---

### 4. **Database Design (MongoDB)**

#### **Collections**:

1. **Students**: 
   - Stores student details.
   ```json
   {
     "_id": ObjectId,
     "name": "John Doe",
     "roll_number": "123",
     "attendance": [
       { "date": "2024-10-10", "status": "present" },
       { "date": "2024-10-11", "status": "absent" },
       ...
     ]
   }
   ```

2. **Attendance**: 
   - Stores attendance data for each class.
   ```json
   {
     "_id": ObjectId,
     "class_date": "2024-10-10",
     "attendance_records": [
       { "roll_number": "123", "status": "present" },
       { "roll_number": "124", "status": "absent" },
       ...
     ]
   }
   ```

#### **Attendance Calculation**:
- Each student will have an array of attendance records linked by date.
- Calculate the percentage of attendance for each student based on the number of present days vs. total days.

---

### 5. **Real-time Updates with Firebase**

- Firebase Realtime Database or Firestore can be used to update the student’s dashboard in real-time when the teacher marks attendance.
- Whenever the teacher updates attendance (via Go API), the data can also be pushed to Firebase.
- Firebase listeners on the student’s side will update the attendance stats in real-time.

#### **Firebase Setup**:
- **Firebase Authentication** to handle user login (teacher/student).
- **Firebase Realtime Database** or **Firestore** for live attendance updates.
  
##### Example of Firebase Realtime Database structure:
```json
{
  "attendance": {
    "student_id_123": {
      "2024-10-10": "present",
      "2024-10-11": "absent"
    }
  }
}
```

---

### 6. **Attendance Calculation and Status (75% Rule)**

- The percentage for each student is calculated as:
  \[
  \text{Attendance Percentage} = \frac{\text{Days Present}}{\text{Total Days}} \times 100
  \]
  - **Green**: 75% and above.
  - **Orange**: 65-74%.
  - **Red**: Below 65%.

In the backend, this can be dynamically calculated using the total attendance records for each student.

---

### 7. **Deployment**:
- **Frontend**: Host on **Firebase Hosting** (for static assets).
- **Backend**: Deploy the Go API using platforms like **Heroku**, **Google Cloud**, or **AWS**.
- **Database**: MongoDB hosted on **MongoDB Atlas** or self-hosted, and Firebase for real-time data syncing.

---

Yes, you can absolutely build the frontend using **PHP**. PHP is a server-side language typically used for building dynamic web pages and interacting with databases. While it’s not typically associated with modern front-end development, it can handle the job for your attendance tracker application by generating dynamic HTML, CSS, and JavaScript content.

Here’s how you can approach building the frontend with PHP for the **centralized attendance tracker**:

### 1. **Frontend Structure Overview**:
   - PHP will generate the HTML content dynamically based on data fetched from MongoDB and Firebase.
   - You will still use **HTML**, **CSS** (for styling), and **JavaScript** (for interactivity and real-time updates). PHP will handle backend processing, while the frontend can include JavaScript for things like AJAX calls and real-time data updates.
   - **PHP** is responsible for rendering pages and handling server-side logic like student/teacher authentication, attendance submission, and fetching data from MongoDB.

---

### 2. **Frontend for Teacher and Student Using PHP**

You can structure your application as follows:

#### **Directory Structure**:
```
/attendance-tracker
  /assets
    /css
      - style.css
    /js
      - main.js
  /views
    - teacher-dashboard.php
    - student-dashboard.php
    - login.php
  /includes
    - db.php        // For MongoDB connection
    - auth.php      // For authentication logic
  - index.php       // Main entry point
```

#### **Basic Example of a PHP Script to Render the Teacher Dashboard**:

##### **`teacher-dashboard.php`**:
This page will allow teachers to add students, mark attendance, and view the attendance summary.

```php
<?php
// Include database connection
include('includes/db.php');

// Check if teacher is logged in (you can use Firebase Auth or PHP sessions)
session_start();
if (!isset($_SESSION['user_role']) || $_SESSION['user_role'] != 'teacher') {
    header("Location: login.php");
    exit();
}

// Fetch the list of students from MongoDB
$students = $db->students->find();  // Assuming you're using MongoDB and have a 'students' collection

// Handle form submission for marking attendance
if ($_SERVER['REQUEST_METHOD'] == 'POST' && isset($_POST['attendance'])) {
    // Update attendance in the MongoDB database
    $rollNumber = $_POST['roll_number'];
    $status = $_POST['status'];
    $date = date('Y-m-d');  // Current date
    
    // Add attendance to the student's record
    $db->students->updateOne(
        ['roll_number' => $rollNumber],
        ['$push' => ['attendance' => ['date' => $date, 'status' => $status]]]
    );
    
    echo "Attendance updated successfully!";
}
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Teacher Dashboard</title>
    <link rel="stylesheet" href="assets/css/style.css">
</head>
<body>

    <h1>Teacher Dashboard</h1>
    
    <h2>Add Student</h2>
    <form action="teacher-dashboard.php" method="POST">
        <label for="name">Student Name:</label>
        <input type="text" id="name" name="name" required>
        <label for="roll_number">Roll Number:</label>
        <input type="text" id="roll_number" name="roll_number" required>
        <input type="submit" value="Add Student">
    </form>
    
    <h2>Mark Attendance</h2>
    <form action="teacher-dashboard.php" method="POST">
        <label for="roll_number">Roll Number:</label>
        <input type="text" id="roll_number" name="roll_number" required>
        
        <label for="status">Status:</label>
        <select id="status" name="status" required>
            <option value="present">Present</option>
            <option value="absent">Absent</option>
            <option value="leave">On Leave</option>
        </select>
        
        <input type="submit" name="attendance" value="Mark Attendance">
    </form>

    <h2>Attendance Summary</h2>
    <table>
        <tr>
            <th>Roll Number</th>
            <th>Name</th>
            <th>Attendance Percentage</th>
        </tr>
        <?php
        // Calculate and display attendance percentage for each student
        foreach ($students as $student) {
            $totalClasses = count($student['attendance']);
            $daysPresent = 0;
            foreach ($student['attendance'] as $record) {
                if ($record['status'] == 'present') {
                    $daysPresent++;
                }
            }
            $attendancePercentage = ($totalClasses > 0) ? ($daysPresent / $totalClasses) * 100 : 0;
            $color = $attendancePercentage >= 75 ? 'green' : ($attendancePercentage >= 65 ? 'orange' : 'red');
            echo "<tr style='color: $color;'>
                    <td>{$student['roll_number']}</td>
                    <td>{$student['name']}</td>
                    <td>{$attendancePercentage}%</td>
                </tr>";
        }
        ?>
    </table>

</body>
</html>
```

This page allows teachers to:
1. Add new students (via the `POST` method).
2. Mark attendance for students.
3. View a list of students with their attendance percentage, color-coded based on the rules.

#### **Basic Example of a PHP Script for the Student Dashboard**:

##### **`student-dashboard.php`**:
This page shows the attendance record for a student.

```php
<?php
include('includes/db.php');

session_start();
if (!isset($_SESSION['user_role']) || $_SESSION['user_role'] != 'student') {
    header("Location: login.php");
    exit();
}

// Fetch the logged-in student details using session (assuming session holds roll number)
$rollNumber = $_SESSION['user_roll_number'];
$student = $db->students->findOne(['roll_number' => $rollNumber]);

?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Student Dashboard</title>
    <link rel="stylesheet" href="assets/css/style.css">
</head>
<body>

    <h1>Student Dashboard</h1>
    
    <h2>Attendance Record</h2>
    <table>
        <tr>
            <th>Date</th>
            <th>Status</th>
        </tr>
        <?php
        // Display the student's attendance records
        foreach ($student['attendance'] as $record) {
            echo "<tr>
                    <td>{$record['date']}</td>
                    <td>{$record['status']}</td>
                  </tr>";
        }
        ?>
    </table>

</body>
</html>
```

In this page:
- The student’s attendance records are fetched from MongoDB based on their roll number and displayed.
- Attendance records are shown with the date and status (present, absent, or leave).

### 3. **MongoDB Integration with PHP**:

To interact with MongoDB from PHP, you will need the **MongoDB PHP Driver**.

#### **Installing MongoDB PHP Driver**:
Run the following commands to install the MongoDB PHP driver:
```bash
composer require mongodb/mongodb
```

#### **Database Connection (in `db.php`)**:
```php
<?php
require 'vendor/autoload.php';  // Composer autoload

// MongoDB connection
$mongoClient = new MongoDB\Client("mongodb://localhost:27017");  // Replace with your MongoDB URI
$db = $mongoClient->attendanceTracker;  // Your database name
?>
```

### 4. **Real-time Updates**:
PHP on its own doesn’t offer real-time updates like Firebase, but you can achieve a similar effect using **AJAX** and **polling**. The student’s page could periodically check for attendance updates using AJAX every few seconds and update the DOM accordingly.

---

### Conclusion:

Yes, you can build the frontend using PHP! While it won’t be as modern or flexible as using JavaScript frameworks like React or Vue, PHP is perfectly capable of handling this project. It will generate the dynamic HTML pages, interact with MongoDB, and handle the server-side logic.
