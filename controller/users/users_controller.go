package users

import (
	"gin_test/model/user"
	"gin_test/util/crypto"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ユーザー登録画面へ
func SignUpPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

// ユーザーを登録して、ログイン画面へ
func RegisterUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	password, _ = crypto.PasswordEncrypt(password)
	data := user.User{Name: name, Password: password}
	data.Create()
	c.Redirect(301, "/login")

}

// ログイン画面へ
func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

//
func LoginUser(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	user := user.GetUser(name)
	dbPassword := user.Password
	err := crypto.CompareHashAndPassword(dbPassword, password)
	if err != nil {
		log.Println("not login")
		c.HTML(400, "login.html", gin.H{})
	} else {
		session := sessions.Default(c)
		session.Set("loginUserID", user.ID)
		log.Println("login OK")
		session.Save()
		c.Redirect(302, "/")
	}
}
func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/")
}
