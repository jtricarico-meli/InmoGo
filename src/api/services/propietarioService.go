package services

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/repositories"
	"InmoGo/src/api/utils"
	"errors"
	"fmt"
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

func (p *PropietarioService) Save(propietario *models.Propietario) *models.Propietario {
	password, err := utils.HashPassword(propietario.Password)
	if err != nil {
		fmt.Println(fmt.Sprintf("SAVE ERR: %s", err))
		return nil
	}
	propietario.Password = password
	p.repository.Save(propietario)
	return nil
}

func (p *PropietarioService) Login(mail string, pass string) (string, error) {

	propietario := p.repository.Login(mail, pass)
	if propietario != nil {
		if !utils.CheckPasswordHash(pass, propietario.Password) {
			return "", errors.New("incorrect password")
		}
	} else {
		return "", errors.New(fmt.Sprintf("not found propietario with mail: %s", mail))
	}

	token, err := p.JWT.CreateToken(mail, loginDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}
