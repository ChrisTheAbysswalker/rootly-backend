package services

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/models"
	"gorm.io/gorm"
	"math"
)

type EspecieService struct {
	DB *gorm.DB
}

func (s *EspecieService) GetInventario() ([]models.Especie, error) {
	var especies []models.Especie
	err := s.DB.Preload("FamiliaBotanica").Preload("EstadoSalud").Find(&especies).Error
	if err != nil {
		return nil, err
	}
	
	return especies, nil
}

func (s *EspecieService) GetEspecie(id string) (models.Especie, error) {
	var especie models.Especie
	err := s.DB.Preload("FamiliaBotanica").Preload("EstadoSalud").First(&especie, id).Error
	if err != nil {
		return models.Especie{}, err
	}
	return especie, nil
}

func (s *EspecieService) GetFamilias() ([]models.FamiliaBotanica, error) {
	var familias []models.FamiliaBotanica
	err := s.DB.Find(&familias).Error
	if err != nil {
		return nil, err
	}
	
	return familias, nil
}

func (s *EspecieService) GetEstados() ([]models.EstadoSalud, error) {
	var estados []models.EstadoSalud
	err := s.DB.Find(&estados).Error
	if err != nil {
		return nil, err
	}
	
	return estados, nil
}

func (s *EspecieService) GetEcosistemaStats() (map[string]interface{}, error) {
	var humedadPromedio float64
	var especiesEnAlerta int64
	var sensoresActivos int64
	var totalEspecies int64
	var especiesSaludables int64

	s.DB.Model(&models.RegistroSalud{}).Select("COALESCE(AVG(humedad_porcentaje), 0)").Scan(&humedadPromedio)
	s.DB.Model(&models.RegistroSalud{}).Distinct("id_especie").Count(&sensoresActivos)
	s.DB.Model(&models.Especie{}).Where("id_estado != ?", 1).Count(&especiesEnAlerta)
	s.DB.Model(&models.Especie{}).Count(&totalEspecies)
	s.DB.Model(&models.Especie{}).Where("id_estado = ?", 1).Count(&especiesSaludables)

	saludVivero := 0.0
    if totalEspecies > 0 {
        saludVivero = (float64(especiesSaludables) / float64(totalEspecies)) * 100
    }

    saludVivero = math.Round(saludVivero*100) / 100
    humedadPromedio = math.Round(humedadPromedio*100) / 100

    stats := map[string]interface{}{
        "salud_vivero":      saludVivero,
        "humedad_promedio":  humedadPromedio,
        "especies_alerta":   especiesEnAlerta,
        "sensores_activos":  sensoresActivos,
    }

	return stats, nil
}

func (s *EspecieService) Create(especie *models.Especie) error {
    return s.DB.Create(especie).Error
}

func (s *EspecieService) Update(id string, data *models.Especie) error {
    var especie models.Especie
    if err := s.DB.First(&especie, id).Error; err != nil {
        return err
    }
    return s.DB.Model(&especie).Updates(data).Error
}

func (s *EspecieService) Delete(id string) error {
    return s.DB.Delete(&models.Especie{}, id).Error
}

func (s *EspecieService) GetRegistroSaludByID(idEspecie string) (models.RegistroSalud, error) {
    var salud models.RegistroSalud
    err := s.DB.Preload("Especie").Where("id_especie = ?", idEspecie).Order("fecha_monitoreo DESC").First(&salud).Error 
    if err != nil {
        return models.RegistroSalud{}, err
    }
    
    return salud, nil
}

func (s *EspecieService) CreateRegistroSalud(data *models.RegistroSalud) error {
    return s.DB.Create(data).Error
}

func (s *EspecieService) UpdateRegistroSalud(idEspecie string, data *models.RegistroSalud) error {
    var registro models.RegistroSalud
    if err := s.DB.Where("id_especie = ?", idEspecie).First(&registro).Error; err != nil {
        return err
    }

    return s.DB.Model(&registro).Select("HumedadPorcentaje", "TempCelsius").Updates(data).Error
}