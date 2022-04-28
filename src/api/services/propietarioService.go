package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
)

type PropietarioService struct {
	repository *repositories.PropietarioRepository
}

func NewPropietarioService(repository *repositories.PropietarioRepository) *PropietarioService {
	return &PropietarioService{repository: repository}
}

func (p *PropietarioService) Get(ID int) *models.Propietario {
	return p.repository.Get(ID)
}

func (p *PropietarioService) Save(propietario *models.Propietario) {
	p.repository.Save(propietario)
}
