package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/utils"
	"strconv"
	"time"
)

type PropietarioService struct {
	repository *repositories.PropietarioRepository
	JWT        utils.Maker
}

const loginDuration = 24 * time.Hour

func NewPropietarioService(repository *repositories.PropietarioRepository, jwt utils.Maker) *PropietarioService {
	return &PropietarioService{repository: repository, JWT: jwt}
}

func (p *PropietarioService) Put(ID int, propNew models.Propietario) (*models.Propietario, error) {
	// Obtengo el propietario
	propitario := p.repository.Get(ID)

	if propitario.ID != 0 {
		propitario.Apellido = propNew.Apellido
		propitario.Nombre = propNew.Nombre
		propitario.Telefono = propNew.Telefono

		p.repository.Save(propitario)
		return propitario, nil
	}

	return nil, utils.InmoError{
		Code:    404,
		Message: "propietario not found",
	}
}

func (p *PropietarioService) Get(ID int) *models.Propietario {
	return p.repository.Get(ID)
}

func (p *PropietarioService) Save(propietario *models.Propietario) (*models.Propietario, error) {
	password, err := utils.HashPassword(propietario.Password)
	if err != nil {
		return nil, err
	}
	propietario.Password = password
	p.repository.Save(propietario)
	return propietario, err
}

func (p *PropietarioService) Login(mail string, pass string) (*models.TokenLogin, error) {

	propietario, err := p.repository.Login(mail)
	if err != nil {
		return nil, err
	}
	if propietario != nil {
		if !utils.CheckPasswordHash(pass, propietario.Password) {
			return nil, utils.NewError(403, "incorrect password")
		}
	}

	token, err := p.JWT.CreateToken(strconv.FormatUint(uint64(propietario.ID), 10), loginDuration)
	if err != nil {
		return nil, err
	}

	return &models.TokenLogin{Token: token}, nil
}
