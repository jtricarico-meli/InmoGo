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

func (p *PropietarioService) Login(mail string, pass string) (string, error) {

	propietario, err := p.repository.Login(mail)
	if err != nil {
		return "", err
	}
	if propietario != nil {
		if !utils.CheckPasswordHash(pass, propietario.Password) {
			return "", utils.NewError(403, "incorrect password")
		}
	}

	token, err := p.JWT.CreateToken(mail, loginDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}
