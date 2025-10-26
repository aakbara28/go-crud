package service

import (
	"go-crud/model"

	"gorm.io/gorm"
)

type StudentService interface {
	Create(s model.Student) error
	GetByStudentID(studentId string) (model.Student, error)
	Update(s model.Student) error
	DeleteByStudentID(studentId string) error
}

type studentService struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) StudentService {
	return &studentService{db}
}

func (s *studentService) Create(st model.Student) error {
	tx := s.db.Begin()
	err := tx.Create(&st).Error
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

func (s *studentService) GetByStudentID(studentId string) (model.Student, error) {
	var st model.Student
	err := s.db.Where("student_id = ?", studentId).First(&st).Error
	return st, err
}

func (s *studentService) Update(st model.Student) error {
	tx := s.db.Begin()
	var existing model.Student
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
	var existing model.Student
	err := tx.Where("student_id = ?", studentId).First(&existing).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("student_id = ?", studentId).Delete(&model.Student{}).Error
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}


