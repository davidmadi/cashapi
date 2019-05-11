package model

import (
	"cashapi/src/database/connection"
	"encoding/json"
	"fmt"
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
func (model *LoginRequest) Authenticate() LoginResponse {

	response := LoginResponse{
		Token: ""}

	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return response
	}

	query := fmt.Sprintf("select login from users where email='%v' and password='%v'", model.Email, model.Password)
	rows, ok := connection.ExecuteQuery(query, con)
	if ok {
		var loginRecords = []LoginRecord{}
		var loginRec = LoginRecord{}
		for rows.Next() {
			err := rows.Scan(&loginRec.login)
			if err == nil {
				loginRecords = append(loginRecords, loginRec)
			} else {
				log.Fatal(err)
			}
		}
		if len(loginRecords) > 0 {
			response.Token = loginRecords[0].login
			response.Valid = true
		}
	}
	defer con.Close()

	return response
}
