package model

import (
	//"cashapi/src/model"
	"encoding/json"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PaymentModel struct {
	gorm.Model
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

//ToJSON PaymentModel to json
func (model *PaymentModel) ToJSON() string {
	by, _ := json.Marshal(model)
	return string(by)
}
