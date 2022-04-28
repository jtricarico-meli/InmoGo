package controllers

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/services"
)

type PropietarioController struct {
	service *services.PropietarioService
}

func NewPropietarioController(service *services.PropietarioService) *PropietarioController {
	return &PropietarioController{service: service}
}

func (p *PropietarioController) Get(ID int) *models.Propietario {
	return p.service.Get(ID)
}

func (p *PropietarioController) Save(propietario interface{}) *models.Propietario {
	return &models.Propietario{}
}
