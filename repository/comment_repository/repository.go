package comment_repository

import (
	"portfolio/InstaMini/dto"
	"portfolio/InstaMini/pkg/errs"
)

type Repository interface {
	CreateComment(c dto.NewCommentRequest) (*dto.NewCommentResponse, errs.Error)
	GetUserId(id int) (int, errs.Error)
	GetComments(userId int) (*dto.GetCommentsResponse, errs.Error)
	EditComment(c dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.Error)
	DeleteComment(id int) errs.Error
	PhotoExist(photoId int) (bool, errs.Error)
}
