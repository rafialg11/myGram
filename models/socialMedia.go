package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Social Media store user's social media links and name
type SocialMedia struct {
	GormModel
	Name      string `json:"name" form:"name" valid:"required~name is Required"`
	SocialUrl string `json:"social_url" form:"social_url" valid:"required~Url is required"`
	UserID    uint
	User      *User
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)

	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
