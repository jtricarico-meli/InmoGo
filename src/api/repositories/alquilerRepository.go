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

func (a *AlquilerRepository) GetAllByInmueble(inmuebleID int) []*models.Alquiler {
	var alquiler []*models.Alquiler
	a.db.Where("inmueble_id = ?", inmuebleID).Find(&alquiler)

	return alquiler
}
