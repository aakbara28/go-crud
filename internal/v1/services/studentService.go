package services

import (
	"go-crud/internal/v1/models"

	"gorm.io/gorm"
)

type StudentService interface {
	Create(s models.Student) error
	GetByStudentID(studentId string) (models.Student, error)
	Update(s models.Student) error
	DeleteByStudentID(studentId string) error
}

type studentService struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) StudentService {
	return &studentService{db}
}

func (s *studentService) Create(st models.Student) error {
	tx := s.db.Begin()
	err := tx.Create(&st).Error
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

func (s *studentService) GetByStudentID(studentId string) (models.Student, error) {
	var st models.Student
	err := s.db.Where("student_id = ?", studentId).First(&st).Error
	return st, err
}

func (s *studentService) Update(st models.Student) error {
	tx := s.db.Begin()
	var existing models.Student
	err := tx.Where("student_id = ?", st.StudentID).First(&existing).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Save(&st).Error
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

func (s *studentService) DeleteByStudentID(studentId string) error {
	tx := s.db.Begin()
	var existing models.Student
	err := tx.Where("student_id = ?", studentId).First(&existing).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("student_id = ?", studentId).Delete(&models.Student{}).Error
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}