package repositories

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/utils"
	"fmt"
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

func (a *AlquilerRepository) Get(ID int) (*models.Alquiler, error) {
	var alquiler *models.Alquiler
	a.db.First(&alquiler, ID)

	if alquiler.ID != 0 {
		return alquiler, nil
	}
	return nil, utils.InmoError{
		Code:    404,
		Message: fmt.Sprintf("not found alquiler with id: %v", ID),
	}
}

func (a *AlquilerRepository) GetAllByInmueble(inmuebleID int) ([]*models.Alquiler, error) {
	var alquiler []*models.Alquiler
	a.db.Where("inmueble_id = ?", inmuebleID).Find(&alquiler)

	if len(alquiler) != 0 {
		return alquiler, nil
	}
	return nil, utils.InmoError{
		Code:    404,
		Message: fmt.Sprintf("not found alquileres with inmueble id: %v", inmuebleID),
	}
}
