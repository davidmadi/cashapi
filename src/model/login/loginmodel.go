package loginmodel

import "encoding/json"

type LoginRequest struct {
	Login    string `form:"login" json:"user"`
	Password string `form:"password" json:"password"`
}

type LoginResponse struct {
	Token string
}

func (model *LoginModel) ToJSON() string {
	by, _ := json.Marshal(model)
	return string(by)
}
