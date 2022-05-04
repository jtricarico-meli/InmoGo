package repositories

import (
	"InmoGo/src/api/models"
	"gorm.io/gorm"
)

type AlquilerRepository struct {
	db *gorm.DB
}

func NewAlquilerRepository(db *gorm.DB) *AlquilerRepository {
	return &AlquilerRepository{db: db}
}

func (a *AlquilerRepository) Save(alquiler *models.Alquiler) {
	a.db.Save(alquiler)
}

func (a *AlquilerRepository) Get(ID int) *models.Alquiler {
	var alquiler *models.Alquiler
	a.db.First(&alquiler, ID)

	return alquiler
}

func (a *AlquilerRepository) GetAll(propietarioID int) []*models.Alquiler {
	var alquiler []*models.Alquiler
	a.db.Where("propietario_id = ?", propietarioID).Find(&alquiler)

	return alquiler
}
