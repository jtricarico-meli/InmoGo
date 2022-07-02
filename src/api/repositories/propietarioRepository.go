package repositories

import (
	"InmoGo/src/api/models"
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

func (p *PropietarioRepository) Login(mail string, password string) *models.Propietario {
	var propietario *models.Propietario
	p.db.Where("mail = ?", mail).Find(&propietario)

	return propietario
}
