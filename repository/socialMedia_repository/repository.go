package social_media_repository

import (
	"portfolio/InstaMini/dto"
	"portfolio/InstaMini/pkg/errs"
)

type Repository interface {
	CreateSocialMedia(sm dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.Error)
	GetSocialMedias(userId int) (*dto.GetSocialMediaResponse, errs.Error)
	UpdateSocialMedia(sm dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.Error)
	DeleteSocialMedia(id int) errs.Error
	GetUserID(id int) (int, errs.Error)
}
