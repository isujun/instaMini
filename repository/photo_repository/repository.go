package photo_repository

import (
	"portfolio/InstaMini/dto"
	"portfolio/InstaMini/pkg/errs"
)

type Repository interface {
	CreatePhoto(p dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.Error)
	GetPhotos() (*dto.GetPhotosResponse, errs.Error)
	EditPhoto(p dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.Error)
	DeletePhoto(id int) errs.Error
	CheckIfPhotoBelongToUser(photoId, userId int) (bool, errs.Error)
}
