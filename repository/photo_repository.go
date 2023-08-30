package repository

import (
	"errors"
	"task-5-vix-btpns-SofyanEgiLesmana/models"

	"gorm.io/gorm"
)

type PhotoRepository struct{}

func NewPhotoRepository() *PhotoRepository {
	return &PhotoRepository{}
}

func (r PhotoRepository) Create(db *gorm.DB, p *models.Photo) (photo *models.Photo, err error) {
	err = db.Debug().Create(&p).Error
	photo = p
	return photo, err
}

func (r PhotoRepository) GetAll(db *gorm.DB) (photos *[]models.Photo, err error) {
	err = db.Debug().Find(&photos).Error
	return
}

func (r PhotoRepository) Update(db *gorm.DB, p *models.Photo, photoID string) (photo *models.Photo, err error) {
	result := db.Model(&models.Photo{}).Where("id = ?", photoID).Updates(models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoUrl: p.PhotoUrl,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("user doesn't exist")
	}
	return p, nil
}

func (r PhotoRepository) Delete(db *gorm.DB, photoID string) (err error) {
	dbResult := db.Model(&models.Photo{}).Where("id = ?", photoID).Delete(&models.Photo{})
	if dbResult.Error != nil {
		return dbResult.Error
	}

	if dbResult.RowsAffected == 0 {
		return errors.New("photo doesn't exist")
	}
	return nil
}
