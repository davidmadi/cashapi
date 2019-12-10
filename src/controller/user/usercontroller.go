package usercontroller

import (
	"cashapi/src/model"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Cria usuario
// @Description Cria usuário
// @ID user-create
// @Accept json
// @Produce json
// @Param body body model.User true "User"
// @Success 200 {object} model.User
// @Header 200 {string} Token "qwerty"
// @Router /user/create [post]
func Create(c *gin.Context) {
	var userRequest model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResp, err := userRequest.Insert()
	if err == nil {
		c.JSON(200, userResp)
	} else {
		c.JSON(500, err)
	}
}

// List godoc
// @Summary Lista grupo de usuarios
// @Description Lista de grupo de usuarios
// @ID user-create
// @Accept json
// @Produce json
// @Param body body model.User true "User"
// @Success 200 {array} model.User
// @Failure 500 {string} error
// @Header 200 {string} Token "qwerty"
// @Router /user/list [post]
func List(c *gin.Context) {

	var userRequest model.User
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(&userRequest)
	users, err := model.ListUsers(&userRequest)

	if users != nil {
		c.JSON(200, users)
	} else {
		c.JSON(500, err)
	}
}
