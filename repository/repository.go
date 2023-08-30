package repository

import (
	"task-5-vix-btpns-SofyanEgiLesmana/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(db *gorm.DB, user *models.User) (err error)
	Login(db *gorm.DB, email string) (user *models.User, err error)
	Update(db *gorm.DB, u *models.User, userID string) (user *models.User, err error)
	Delete(db *gorm.DB, userID string) (err error)
}

type IPhotoRepository interface {
	Create(db *gorm.DB, p *models.Photo) (photo *models.Photo, err error)
	GetAll(db *gorm.DB) (photos *[]models.Photo, err error)
	Update(db *gorm.DB, p *models.Photo, photoID string) (product *models.Photo, err error)
	Delete(db *gorm.DB, photoID string) (err error)
}
