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

### Final Thoughts:

This system allows for real-time tracking of attendance in a classroom, with color-coded statuses, easy data entry for teachers, and real-time updates for students. You can extend this by adding more features like notifications for students (via Firebase Cloud Messaging) or even integrating a calendar for class timings.
