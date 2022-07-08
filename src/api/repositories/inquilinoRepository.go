package repositories

import (
	"InmoGo/src/api/models"
	"InmoGo/src/api/utils"
	"fmt"
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

func (i *InquilinoRepository) Get(ID int) (*models.Inquilino, error) {
	var inquilino *models.Inquilino
	i.db.First(&inquilino, ID)

	if inquilino.ID != 0 {
		return inquilino, nil
	}
	return nil, utils.InmoError{
		Code:    404,
		Message: fmt.Sprintf("not found alquiler with id: %v", ID),
	}
}
