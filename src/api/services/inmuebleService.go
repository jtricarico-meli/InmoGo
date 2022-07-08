package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
)

type InmuebleService struct {
	repository *repositories.InmuebleRepository
}

func NewInmuebleService(repository *repositories.InmuebleRepository) *InmuebleService {
	return &InmuebleService{repository: repository}
}

func (i *InmuebleService) Get(ID int) (*models.Inmueble, error) {
	return i.repository.Get(ID)
}

func (i *InmuebleService) Save(propietario *models.Inmueble) *models.Inmueble {
	i.repository.Save(propietario)
	return nil
}

func (i *InmuebleService) GetAll(propietarioID int) (*models.Inmuebles, error) {
	return i.repository.GetAll(propietarioID)
}
