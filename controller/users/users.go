package users

import "github.com/gin-gonic/gin"

func SignUpPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}
