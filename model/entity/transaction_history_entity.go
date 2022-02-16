package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type TransactionHistory struct {
	ID         int       `json:"id"`
	ProductID  int       `json:"product_id" valid:"required"`
	UserID     int       `json:"user_id" valid:"required"`
	Quantity   int       `json:"quantity" valid:"required"`
	TotalPrice int       `json:"total_price" valid:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (t *TransactionHistory) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
