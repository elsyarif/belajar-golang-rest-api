package domain

import (
	"context"

	"github.com/elSyarif/belajar-golang-rest-api/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *model.User) error
	VerifyUserEmail(ctx context.Context, email string) (string, error)
	VerifyUserCredential(ctx context.Context, email string, password string) (string, error)
	VerifyUser(ctx context.Context, userId string) (*model.User, error)
	GetUser(ctx context.Context) (*[]model.User, error)
}

type UserService interface {
	Adduser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context) (*[]model.User, error)
}

type UserHandler interface {
	AddUser(c *gin.Context)
	GetUser(c *gin.Context)
}
