package user

import (
	"context"
	"net/http"

	"github.com/elSyarif/belajar-golang-rest-api/internal/domain"
	"github.com/elSyarif/belajar-golang-rest-api/internal/domain/model"
	"github.com/elSyarif/belajar-golang-rest-api/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandlerImpl struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) domain.UserHandler {
	return &userHandlerImpl{service: s}
}

func (u *userHandlerImpl) AddUser(c *gin.Context) {
	user := model.User{}

	user.Id = uuid.New().String()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := c.ShouldBindJSON(&user)
	helper.PanicError(err)

	usr, err := u.service.Adduser(ctx, &user)
	helper.PanicError(err)
	c.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusCreated,
		"status": "success",
		"data":   usr,
	})
}

func (handler *userHandlerImpl) GetUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	users, err := handler.service.GetUser(ctx)
	helper.PanicError(err)

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "success",
		"data":   users,
	})
}
