package loginmodel

import (
	"cashapi/src/database/connection"
	"encoding/json"
)

//LoginRequest request struct
type LoginRequest struct {
	Login    string `form:"login" json:"login"`
	Password string `form:"password" json:"password"`
}

//LoginResponse response do login
type LoginResponse struct {
	Token string
}

//ToJSON LoginRequest to json
func (model *LoginRequest) ToJSON() string {
	by, _ := json.Marshal(model)
	return string(by)
}

//Authenticate the user
func (model *LoginRequest) Authenticate() LoginResponse {

	response := LoginResponse{
		Token: ""}

	_, err := connection.Open()
	if err != nil {
		return response
	}

	response.Token = "12345"

	return response

}
