package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represent model for user comment
type Comment struct {
	GormModel
	Message string `json:"message" form:"message" valid:"required~message is Required"`
	UserID  uint
	PhotoID uint
	User    *User
	Photo   *Photo
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)

	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
