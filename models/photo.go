package models

import (
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo_url is required"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		return errCreate
	}
	p.ID = "p-" + helpers.GenerateID()
	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		return errCreate
	}
	return nil
}
