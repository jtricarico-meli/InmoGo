package controllers

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/services"
	"strconv"
)

type PropietarioController struct {
	service *services.PropietarioService
}

func NewPropietarioController(service *services.PropietarioService) *PropietarioController {
	return &PropietarioController{service: service}
}

func (p *PropietarioController) Get(ID string) *models.Propietario {
	id, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}
	return p.service.Get(id)
}

func (p *PropietarioController) Save(propietario models.Propietario) *models.Propietario {
	return p.service.Save(&propietario)
}
