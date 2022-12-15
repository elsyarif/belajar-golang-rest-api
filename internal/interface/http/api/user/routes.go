package user

import (
	"github.com/elSyarif/belajar-golang-rest-api/internal/helper"
	"github.com/elSyarif/belajar-golang-rest-api/internal/infrastructure/database/mysql"
	"github.com/elSyarif/belajar-golang-rest-api/internal/infrastructure/repositories"
	"github.com/elSyarif/belajar-golang-rest-api/internal/infrastructure/services"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) *gin.Engine {
	db, err := mysql.NewConnection()
	helper.PanicError(err)

	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	handler := NewUserHandler(service)

	router.POST("/users", handler.AddUser)
	router.GET("/users", handler.GetUser)

	return router
}
