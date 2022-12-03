package services

import (
	"context"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain/model"
	"github.com/elSyarif/belajar-golang-rest-api/internal/helper"
)

type userService struct {
	repository domain.UserRepository
}

func NewUserService(r domain.UserRepository) domain.UserService {
	return &userService{repository: r}
}

func (u *userService) Adduser(ctx context.Context, user *model.User) (model.User, error) {
	err := u.repository.AddUser(ctx, user)
	helper.PanicError(err)

	return *user, nil
}
