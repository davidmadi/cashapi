package usercontroller

import (
	"cashapi/src/model"
	"net/http"

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

// Groups godoc
// @Summary Lista grupo de usuarios
// @Description Lista de grupo de usuarios
// @ID user-create
// @Accept json
// @Produce json
// @Success 200 {array} model.Group
// @Failure 500 {string} error
// @Header 200 {string} Token "qwerty"
// @Router /user/groups [get]
func Groups(c *gin.Context) {
	groups, err := model.ListGroups()

	if groups != nil {
		c.JSON(200, groups)
	} else {
		c.JSON(500, err)
	}
}
