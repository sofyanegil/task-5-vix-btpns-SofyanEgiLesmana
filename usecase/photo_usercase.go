package usecase

import (
	"task-5-vix-btpns-SofyanEgiLesmana/models"
	"task-5-vix-btpns-SofyanEgiLesmana/repository"

	"gorm.io/gorm"
)

type PhotoUsecase struct {
	PhotoRepository repository.IPhotoRepository
	DB              *gorm.DB
}

func NewPhotoUsecase(photoRepository repository.IPhotoRepository, db *gorm.DB) *PhotoUsecase {
	return &PhotoUsecase{
		PhotoRepository: photoRepository,
		DB:              db,
	}
}

func (u PhotoUsecase) Create(p *models.Photo, userID string) (photo *models.Photo, err error) {
	p.UserID = userID

	photo, err = u.PhotoRepository.Create(u.DB, p)

	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (u PhotoUsecase) GetAll() (photos *[]models.Photo, err error) {
	photos, err = u.PhotoRepository.GetAll(u.DB)

	if err != nil {
		return nil, err
	}
	return photos, nil
}

func (u PhotoUsecase) Update(p *models.Photo, photoID string) (photo *models.Photo, err error) {
	photo, err = u.PhotoRepository.Update(u.DB, p, photoID)

	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (u PhotoUsecase) Delete(photoID string) (err error) {
	err = u.PhotoRepository.Delete(u.DB, photoID)

	if err != nil {
		return err
	}
	return nil
}
