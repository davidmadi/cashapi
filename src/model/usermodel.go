package model

import (
	"cashapi/src/database/connection"

	"log"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	name    string
	email   string
	GroupID uint
}

// Group model
type Group struct {
	gorm.Model
	Name string
}

// Insert user
func (u *User) Insert() bool {

	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return false
	}

	con.Create(&u)
	con.Close()
	return u.ID > 0
}

// ListGroups all
func ListGroups() []Group {

	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	groups := make([]Group, 0)
	con.Table("groups").Find(&groups)
	con.Close()

	return groups
}
