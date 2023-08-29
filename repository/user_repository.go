package repository

import (
	"errors"
	"task-5-vix-btpns-SofyanEgiLesmana/models"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Register(db *gorm.DB, user *models.User) (err error) {
	err = db.Debug().Create(&user).Error
	return err
}

func (r UserRepository) Login(db *gorm.DB, email string) (user *models.User, err error) {
	err = db.Debug().Where("email = ?", email).Take(&user).Error
	return user, err
}

func (r UserRepository) Update(db *gorm.DB, u *models.User, userID string) (*models.User, error) {
	result := db.Model(&models.User{}).Where("id = ?", userID).Updates(models.User{
		Username: u.Username,
		Email:    u.Email,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("user doesn't exist")
	}
	return u, nil
}

func (r UserRepository) Delete(db *gorm.DB, userID string) (err error) {
	result := db.Model(&models.User{}).Where("id = ?", userID).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user doesn't exist")
	}
	return nil
}
