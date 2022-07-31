package users

import (
	"gin_test/model/user"
	"gin_test/util/crypto"

	"github.com/gin-gonic/gin"
)

func SignUpPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}
func RegisterUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	password, _ = crypto.PasswordEncrypt(password)
	data := user.User{Name: name, Password: password}
	data.Create()
	c.Redirect(301, "/signup")

}
