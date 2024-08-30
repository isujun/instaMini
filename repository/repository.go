package user_repository

import (
	"portfolio/InstaMini/dto"
	"portfolio/InstaMini/entity"
	"portfolio/InstaMini/pkg/errs"
)

type Repository interface {
	CreateUser(newUser dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
	Login(email string) (*entity.User, errs.Error)
	EditUser(user dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.Error)
	DeleteUser(id int) errs.Error
	CountEmail(email string) (int, errs.Error)
	CountUsername(username string) (int, errs.Error)
}
