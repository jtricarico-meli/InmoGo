package repositories

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/utils"
	"fmt"
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

func (i *InmuebleRepository) Get(ID int) (*models.Inmueble, error) {
	var inmueble *models.Inmueble
	i.db.First(&inmueble, ID)

	if inmueble.ID != 0 {
		return inmueble, nil
	}
	return nil, utils.InmoError{
		Code:    404,
		Message: fmt.Sprintf("not found inmueble with id: %v", ID),
	}

}

func (i *InmuebleRepository) GetAll(propietarioID int) (*models.Inmuebles, error) {
	var inmueble []*models.Inmueble
	i.db.Where("propietario_id = ?", propietarioID).Find(&inmueble)

	if len(inmueble) == 0 {
		return nil, utils.InmoError{
			Code:    404,
			Message: fmt.Sprintf("not found inmuebles for propietario with id: %v", propietarioID),
		}
	}
	return &models.Inmuebles{
		Inmuebles: inmueble,
	}, nil
}
