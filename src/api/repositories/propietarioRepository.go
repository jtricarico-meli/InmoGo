package repositories

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/utils"
	"gorm.io/gorm"
)

type PropietarioRepository struct {
	db *gorm.DB
}

func NewPropietarioRepository(db *gorm.DB) *PropietarioRepository {
	return &PropietarioRepository{db: db}
}

func (p *PropietarioRepository) Save(propietario *models.Propietario) {
	p.db.Save(propietario)
}

func (p *PropietarioRepository) Get(ID int) *models.Propietario {
	var propietario *models.Propietario
	p.db.First(&propietario, ID)

	return propietario
}

func (p *PropietarioRepository) Login(mail string) (*models.Propietario, error) {
	var propietario *models.Propietario
	p.db.Where("mail = ?", mail).Find(&propietario)

	if propietario.Mail == "" {
		return nil, utils.NewError(403, "email not Found")
	}

	return propietario, nil
}
