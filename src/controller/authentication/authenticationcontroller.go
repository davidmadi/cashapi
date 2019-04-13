package authenticationcontroller

import (
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Efetua login
// @Description Efetua login
// @ID connect-token
// @Accept  json
// @Produce  json
// @Param Login path string true "Login"
// @Param Password path string true "Password"
// @Success 200 {object} loginmodel.LoginResponse
// @Header 200 {string} Token "qwerty"
// @Router /connect/token [post]
func Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Id":      1,
		"Usuario": "davizera"})
}
