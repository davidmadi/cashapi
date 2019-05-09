package paymentcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Cria um pagamento
// @Description Cria um pagamento
// @ID payment-create
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.PaymentModel
// @Header 200 {string} Token "qwerty"
// @Router /payment/create [post]
func Create(ctx *gin.Context) {
	id, got := ctx.GetPostForm("id")
	if got {
		fmt.Println(id)
	}

	ctx.JSON(200, gin.H{
		"message": id,
	})
}
