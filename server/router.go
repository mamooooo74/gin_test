package server

import (
	"gin_test/controller/users"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("/css", "./static/css")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/signup", users.SignUpPage)
	r.POST("/signup", users.RegisterUser)
	r.GET("/login", users.LoginPage)
	r.POST("/login", users.LoginUser)
	return r
}
