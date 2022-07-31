package server

import (
	"gin_test/controller/users"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("/css", "./static/css")
	r.GET("/signup", users.SignUpPage)
	return r
}
