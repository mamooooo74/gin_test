package user

import (
	"gin_test/database"

	"gorm.io/gorm"
)

func init() {
	database.DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `form:"username" binding:"required" gorm:"unique;not null"`
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
