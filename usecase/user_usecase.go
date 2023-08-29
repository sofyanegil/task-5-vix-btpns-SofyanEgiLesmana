package usecase

import (
	"errors"
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"
	"task-5-vix-btpns-SofyanEgiLesmana/models"
	"task-5-vix-btpns-SofyanEgiLesmana/repository"

	"gorm.io/gorm"
)

type UserUsecase struct {
	UserRepository repository.IUserRepository
	DB             *gorm.DB
}

func NewUserUsecase(repository repository.IUserRepository, db *gorm.DB) *UserUsecase {
	return &UserUsecase{
		UserRepository: repository,
		DB:             db,
	}
}

func (usecase UserUsecase) Register(u *models.User) (user *models.User, err error) {
	err = usecase.UserRepository.Register(usecase.DB, u)
	if err != nil {
		return nil, err
	}
	user = u
	return user, nil
}

func (usecase UserUsecase) Login(email, password string) (token string, err error) {
	user, err := usecase.UserRepository.Login(usecase.DB, email)
	if err != nil {
		return "", errors.New("account not found")
	}

	if !helpers.ComparePassword([]byte(user.Password), []byte(password)) {
		return "", errors.New("password is wrong")
	}

	token = helpers.GenerateToken(user.ID, user.Email)
	return token, nil
}

func (usecase UserUsecase) Update(u *models.User, userID string) (*models.User, error) {
	user, err := usecase.UserRepository.Update(usecase.DB, u, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usecase UserUsecase) Delete(userID string) (err error) {
	err = usecase.UserRepository.Delete(usecase.DB, userID)
	if err != nil {
		return err
	}
	return nil
}
