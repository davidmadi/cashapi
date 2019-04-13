package authenticationcontroller

import (
	loginmodel "cashapi/src/model/login"

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
	loginRequest := loginmodel.LoginRequest{
		Login:    ctx.PostForm("Login"),
		Password: ctx.PostForm("Password")}

	loginResponse := loginmodel.LoginResponse{
		Token: loginRequest.Password}

	ctx.JSON(200, loginResponse)
}
