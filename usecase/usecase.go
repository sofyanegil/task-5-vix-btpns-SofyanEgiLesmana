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
