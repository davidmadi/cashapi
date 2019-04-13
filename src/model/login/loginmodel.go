package loginmodel

import "encoding/json"

//LoginRequest request struct
type LoginRequest struct {
	Login    string `form:"login" json:"user"`
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
