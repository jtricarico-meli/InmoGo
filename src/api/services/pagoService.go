package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
)

type PagoService struct {
	repository *repositories.PagoRepository
}

func NewPagoService(repository *repositories.PagoRepository) *PagoService {
	return &PagoService{repository: repository}
}

func (p *PagoService) Get(ID int) (*models.Pago, error) {
	return p.repository.Get(ID)
}

func (p *PagoService) Save(pago *models.Pago) *models.Pago {
	p.repository.Save(pago)
	return nil
}

func (p *PagoService) GetAll(alquilerID int) []*models.Pago {
	return p.repository.GetAll(alquilerID)
}
