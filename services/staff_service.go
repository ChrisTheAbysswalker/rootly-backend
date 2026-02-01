package services

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/models"
	"gorm.io/gorm"
)

type StaffService struct {
	DB *gorm.DB
}

func (s *StaffService) GetStaff() ([]models.StaffAnimal, error) {
	var staff []models.StaffAnimal
	
	err := s.DB.Preload("CargoStaff").Find(&staff).Error
	if err != nil {
		return nil, err
	}
	
	return staff, nil
}