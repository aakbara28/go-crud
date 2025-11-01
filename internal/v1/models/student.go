package models

// Student represents the student entity
type Student struct {
	StudentID string `gorm:"primaryKey;column:student_id" json:"studentId"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Major     string `json:"major"`
}