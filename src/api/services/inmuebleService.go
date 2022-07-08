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

func (i *InmuebleService) Put(ID int, inmuebleNew models.Inmueble) (*models.Inmueble, error) {
	inmueble, err := i.repository.Get(ID)
	if err != nil {
		return nil, err
	}
	inmueble.Precio = inmuebleNew.Precio
	inmueble.Direccion = inmuebleNew.Direccion
	inmueble.Ambientes = inmuebleNew.Ambientes
	inmueble.Uso = inmuebleNew.Uso
	inmueble.Tipo = inmuebleNew.Tipo
	inmueble.Latitud = inmuebleNew.Latitud
	inmueble.Longitud = inmuebleNew.Longitud

	i.repository.Save(inmueble)
	return inmueble, nil
}

func (i *InmuebleService) Delete(ID int) error {
	return i.repository.Delete(ID)
}

func (i *InmuebleService) Save(propietario *models.Inmueble) *models.Inmueble {
	i.repository.Save(propietario)
	return nil
}

func (i *InmuebleService) GetAll(propietarioID int) (*models.Inmuebles, error) {
	return i.repository.GetAll(propietarioID)
}
