package groupcontroller

import (
	"cashapi/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Cria grupo de usuários
// @Description Cria grupo de usuários
// @ID group-create
// @Accept json
// @Produce json
// @Param body body model.Group true "Group"
// @Success 200 {object} model.Group
// @Header 200 {string} Token "qwerty"
// @Router /group/create [post]
func Create(c *gin.Context) {
	var groupRequest model.Group
	if err := c.ShouldBindJSON(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResp, err := groupRequest.Insert()
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
// @Success 200 {array} model.Group
// @Failure 500 {string} error
// @Header 200 {string} Token "qwerty"
// @Router /group/list [get]
func List(c *gin.Context) {
	groups, err := model.ListGroups()

	if groups != nil {
		c.JSON(200, groups)
	} else {
		c.JSON(500, err)
	}
}
