package user

import (
	"gin_test/database"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	database.DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `form:"name" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func (u *User) Create() {
	result := database.DB.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func GetUser(name string) (u User) {
	database.DB.First(&u, "name = ?", name)
	return
}

func LoginUser(c *gin.Context) User {
	var u User
	session := sessions.Default(c)
	id := session.Get("loginUserID")
	log.Println(id)
	if err := database.DB.First(&u, "id = ?", id); err != nil {
		return u
	}
	return u
}
func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

}
