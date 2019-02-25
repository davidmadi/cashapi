package payment

import (
	"github.com/gin-gonic/gin"
	_ "../model"
)

// Create godoc
// @Summary Cria um pagamento
// @Description Cria um pagamento
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.PaymentModel
// @Header 200 {string} Token "qwerty"
// @Router /payment/create [post]
func Create(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
