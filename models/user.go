package models

import (
	"myGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represent model for user
type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your Username is Required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your Email is Required,email~Invalid Email Format"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Your password is Required,minstringlength(6)~Password has to have a minimum of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Your Role is Required,range(8|100)~Minimum Age is 8 years old"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
