package api

import (
	"github.com/elSyarif/belajar-golang-rest-api/internal/interface/http/api/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RunServer() {
	router := gin.Default()

	user.UserRoute(router)

	err := router.Run(":8888")
	if err != nil {
		return
	}
}
