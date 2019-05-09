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

	if userRequest.Insert() {
		c.JSON(200, 0)
	} else {
		c.JSON(500, 0)
	}
}

// Groups godoc
// @Summary Lista grupo de usuarios
// @Description Lista de grupo de usuarios
// @ID user-create
// @Accept json
// @Produce json
// @Success 200 {array} model.Group
// @Header 200 {string} Token "qwerty"
// @Router /user/groups [get]
func Groups(c *gin.Context) {
	groups := model.ListGroups()

	if groups != nil {
		c.JSON(200, groups)
	} else {
		c.JSON(500, nil)
	}
}
