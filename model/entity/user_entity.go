package entity

import (
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name" valid:"required"`
	Email     string    `json:"email" valid:"required,email" gorm:"uniqueIndex"`
	Password  string    `json:"password" valid:"required,minstringlength(6)"`
	Role      string    `json:"role" valid:"required"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }
