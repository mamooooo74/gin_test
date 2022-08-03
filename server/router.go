package server

import (
	"gin_test/controller/users"
	"gin_test/model/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("/css", "./static/css")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.GET("/", func(c *gin.Context) {
		u := user.LoginUser(c)
		c.HTML(200, "index.html", gin.H{"user": u})
	})
	r.GET("/signup", users.SignUpPage)
	r.POST("/signup", users.RegisterUser)
	r.GET("/login", users.LoginPage)
	r.POST("/login", users.LoginUser)
	r.GET("logout", users.LogOut)
	return r
}
