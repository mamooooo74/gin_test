package user

import (
	"gin_test/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func (u *User) Create() {
	result := database.DB.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}
