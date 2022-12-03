package user

import (
	"context"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain/model"
	"github.com/elSyarif/belajar-golang-rest-api/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type userHandlerImpl struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) domain.UserHandler {
	return &userHandlerImpl{service: s}
}

func (u *userHandlerImpl) AddUser(c *gin.Context) {
	user := &model.User{}

	user.Id = uuid.New().String()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := c.ShouldBind(user)
	helper.PanicError(err)

	usr, err := u.service.Adduser(ctx, user)
	helper.PanicError(err)
	c.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusCreated,
		"status": "success",
		"data": gin.H{
			"userId": usr.Id,
		},
	})
}
