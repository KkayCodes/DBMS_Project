// models/attendance.go
package models

// Attendance represents attendance information for a course
type Attendance struct {
    TotalClasses     int     `bson:"total_classes" json:"total_classes"`
    ClassesAttended  int     `bson:"classes_attended" json:"classes_attended"`
    AttendancePercent float64 `bson:"attendance_percent" json:"attendance_percent"`
}

// UpdateAttendance updates the attendance percentage based on total classes and classes attended
func (a *Attendance) UpdateAttendance() {
    if a.TotalClasses > 0 {
        a.AttendancePercent = (float64(a.ClassesAttended) / float64(a.TotalClasses)) * 100
    } else {
        a.AttendancePercent = 0
    }
}
