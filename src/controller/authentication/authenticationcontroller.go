package authenticationcontroller

import (
	loginmodel "cashapi/src/model/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Efetua login
// @Description Efetua login
// @ID connect-token
// @Accept json
// @Produce json
// @Param body body loginmodel.LoginRequest true "LoginRequest"
// @Success 200 {object} loginmodel.LoginResponse
// @Header 200 {string} Token "qwerty"
// @Router /connect/token [post]
func Create(c *gin.Context) {
	var json loginmodel.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse := loginmodel.LoginResponse{
		Token: json.Password}

	c.JSON(200, loginResponse)
}
