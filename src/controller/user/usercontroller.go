package usercontroller

import (
	"cashapi/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Cria usuario
// @Description Cria usuário
// @ID connect-token
// @Accept json
// @Produce json
// @Param body body usermodel.User true "User"
// @Success 200 {object} usermodel.User
// @Header 200 {string} Token "qwerty"
// @Router /user/create [post]
func Create(c *gin.Context) {
	var userRequest usermodel.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRequest.Insert()

	c.JSON(200, 0)
}
