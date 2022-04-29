package repositories

import (
	"InmoGo/src/api/models"
	"gorm.io/gorm"
)

type InmuebleRepository struct {
	db *gorm.DB
}

func NewInmuebleRepository(db *gorm.DB) *InmuebleRepository {
	return &InmuebleRepository{db: db}
}

func (i *InmuebleRepository) Save(inmueble *models.Inmueble) {
	i.db.Save(inmueble)
}

func (i *InmuebleRepository) Get(ID int) *models.Inmueble {
	var inmueble *models.Inmueble
	i.db.First(&inmueble, ID)

	return inmueble
}

func (i *InmuebleRepository) GetAll(propietarioID int) []*models.Inmueble {
	var inmueble []*models.Inmueble
	i.db.Where("propietario = ?", propietarioID).Find(&inmueble)

	return inmueble
}
