package model

import (
	"cashapi/src/database/connection"
	"encoding/json"
	"log"
)

//LoginRequest request struct
type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

//LoginResponse response do login
type LoginResponse struct {
	Valid bool
	Token string
}

type LoginRecord struct {
	login string
}

//ToJSON LoginRequest to json
func (model *LoginRequest) ToJSON() string {
	by, _ := json.Marshal(model)
	return string(by)
}

//Authenticate the user
func (request *LoginRequest) Authenticate() LoginResponse {

	response := LoginResponse{
		Token: ""}

	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return response
	}

	var users []User
	con.Where(map[string]interface{}{"email": request.Email, "password": request.Password}).Find(&users)

	if len(users) > 0 {
		response.Valid = true
		response.Token = "12313123"
	} else {
		response.Valid = false
		response.Token = ""
	}

	defer con.Close()

	return response
}
