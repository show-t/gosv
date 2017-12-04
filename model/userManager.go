package model

import (
	"github.com/jinzhu/gorm"
)

/* ユーザー情報構造体 */
type User struct {
	gorm.Model
	Name     string `gorm:"size:50"`
	Password string `gorm:"size:50"`
	//Email    string `gorm:"size:255"`
}

/* ユーザー追加 */
func CreateUser(name string, password string) error {
	user := User{
		Name:     name,
		Password: password,
	}
	return DB.Create(&user).Error
}
