package model

import (
	"cashapi/src/database/connection"
	"errors"

	"log"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	GroupID  uint
}

// Group model
type Group struct {
	gorm.Model
	Name string
}

// Insert user
func (u *User) Insert() (*User, error) {

	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	con.Create(&u)
	if u.ID > 0 {
		return u, nil
	}

	return nil, errors.New("")
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
