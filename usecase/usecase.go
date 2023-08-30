package usecase

import (
	"task-5-vix-btpns-SofyanEgiLesmana/models"
)

type IUserUseCase interface {
	Register(u *models.User) (user *models.User, err error)
	Login(email, password string) (token string, err error)
	Update(u *models.User, userID string) (user *models.User, err error)
	Delete(userID string) (err error)
}

type IPhotoUseCase interface {
	Create(p *models.Photo, userID string) (photo *models.Photo, err error)
	GetAll() (photos *[]models.Photo, err error)
	Update(p *models.Photo, photoID string) (product *models.Photo, err error)
	Delete(photoID string) (err error)
}
