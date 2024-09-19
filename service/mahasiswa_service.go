package service

import (
	"go-crud/model"

	"gorm.io/gorm"
)

type MahasiswaService interface {
	Create(m model.Mahasiswa) error
	GetByNIM(nim string) (model.Mahasiswa, error)
	Update(m model.Mahasiswa) error
	DeleteByNIM(nim string) error
	List() ([]model.Mahasiswa, error)
}

type mahasiswaService struct {
	db *gorm.DB
}

func NewMahasiswaService(db *gorm.DB) MahasiswaService {
	return &mahasiswaService{db}
}

func (s *mahasiswaService) Create(m model.Mahasiswa) error {
	result := s.db.Create(&m)
	return result.Error
}

func (s *mahasiswaService) GetByNIM(nim string) (model.Mahasiswa, error) {
	var m model.Mahasiswa
	result := s.db.Where("nim = ?", nim).First(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

func (s *mahasiswaService) Update(m model.Mahasiswa) error {
	result := s.db.Save(&m)
	return result.Error
}

func (s *mahasiswaService) DeleteByNIM(nim string) error {
	result := s.db.Where("nim = ?", nim).Delete(&model.Mahasiswa{})
	return result.Error
}

func (s *mahasiswaService) List() ([]model.Mahasiswa, error) {
	var mahasiswaList []model.Mahasiswa
	result := s.db.Find(&mahasiswaList)
	if result.Error != nil {
		return nil, result.Error
	}
	return mahasiswaList, nil
}
