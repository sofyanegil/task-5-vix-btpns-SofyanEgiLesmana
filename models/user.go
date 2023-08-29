package models

import (
	"errors"
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string  `gorm:"not null" json:"username" form:"username" valid:"required~username is required"`
	Email    string  `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required,email~invalid email format"`
	Password string  `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~password has to have a minimum length of 6 characters"`
	Photo    []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
}

func isEmailTaken(tx *gorm.DB, email string) bool {
	var user User
	err := tx.Where("email = ?", email).First(&user).Error
	return err == nil
}

func isUsernameTaken(tx *gorm.DB, username string) bool {
	var user User
	err := tx.Where("username = ?", username).First(&user).Error
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidation := govalidator.ValidateStruct(u)
	if errValidation != nil {
		return errValidation
	}

	if isEmailTaken(tx, u.Email) {
		return errors.New("email is already taken")
	}

	if isUsernameTaken(tx, u.Username) {
		return errors.New("username is already taken")
	}

	u.ID = "u-" + helpers.GenerateID()
	u.Password = helpers.HashPassword(u.Password)
	return nil
}
