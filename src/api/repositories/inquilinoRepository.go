package repositories

import (
	"InmoGo/src/api/models"
	"gorm.io/gorm"
)

type InquilinoRepository struct {
	db *gorm.DB
}

func NewInquilinoRepository(db *gorm.DB) *InquilinoRepository {
	return &InquilinoRepository{db: db}
}

func (i *InquilinoRepository) Save(inquilino *models.Inquilino) {
	i.db.Save(inquilino)
}

func (i *InquilinoRepository) Get(ID int) *models.Inquilino {
	var inquilino *models.Inquilino
	i.db.First(&inquilino, ID)

	return inquilino
}
