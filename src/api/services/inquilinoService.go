package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
)

type InquilinoService struct {
	repository *repositories.InquilinoRepository
}

func NewInquilinoService(repository *repositories.InquilinoRepository) *InquilinoService {
	return &InquilinoService{repository: repository}
}

func (i *InquilinoService) Get(ID int) (*models.Inquilino, error) {
	return i.repository.Get(ID)
}

func (i *InquilinoService) Save(propietario *models.Inquilino) *models.Inquilino {
	i.repository.Save(propietario)
	return nil
}
