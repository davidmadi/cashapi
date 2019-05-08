package usermodel

import (
	"github.com/jinzhu/gorm"
	//"cashapi/src/database/connection"
)

// User model
type User struct {
	gorm.Model
	name string
	email string
}

// Insert user
func (u *User) Insert() {

	//con, err := connection.Open()
	
	

}