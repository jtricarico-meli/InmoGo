package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/utils"
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

func (p *PropietarioService) Get(mail string) *models.Propietario {
	prop, _ := p.repository.Login(mail)
	return prop
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

	token, err := p.JWT.CreateToken(mail, loginDuration)
	if err != nil {
		return nil, err
	}

	return &models.TokenLogin{Token: token}, nil
}
