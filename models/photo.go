package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represent model for user's photos
type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Title of your photo is Required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Url of your photo is required"`
	Caption  string `json:"caption" form:"caption"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)

	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
