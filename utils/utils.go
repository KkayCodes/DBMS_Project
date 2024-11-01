package utils

import (
    "encoding/csv"
    "encoding/json"
    "os"
)

// ErrorResponse struct for standardized error messages
type ErrorResponse struct {
    Error string `json:"error"`
}

// SuccessResponse struct for standardized success messages
type SuccessResponse struct {
    Message string `json:"message"`
}

// ReadCSVFile reads a CSV file and returns records
func ReadCSVFile(filePath string) ([][]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    return records, nil
}

// FormatAttendanceStatus formats the attendance status based on percentage
func FormatAttendanceStatus(percentage float64) string {
    if percentage >= 75 {
        return "green"
    } else if percentage >= 65 {
        return "yellow"
    }
    return "red"
}

// ErrorResponse generates a JSON error response
func ErrorResponse(message string) ErrorResponse {
    return ErrorResponse{Error: message}
}

// SuccessResponse generates a JSON success response
func SuccessResponse(message string) SuccessResponse {
    return SuccessResponse{Message: message}
}
