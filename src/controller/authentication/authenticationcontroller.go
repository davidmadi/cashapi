package authenticationcontroller

import (
	"cashapi/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate godoc
// @Summary Efetua login
// @Description Efetua login
// @ID connect-token
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "LoginRequest"
// @Success 200 {object} model.LoginResponse
// @Header 200 {string} Token "qwerty"
// @Router /connect/token [post]
func Authenticate(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse := loginRequest.Authenticate()

	c.JSON(200, loginResponse)
}
