package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authenticationcontroller "cashapi/src/controller/authentication"
	paymentcontroller "cashapi/src/controller/payment"
	usercontroller "cashapi/src/controller/user"

	"github.com/labstack/echo"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//echoSwagger "github.com/swaggo/echo-swagger"
	_ "cashapi/docs" // docs is generated by Swag CLI, you have to import it.

	migration "cashapi/src/database/migration"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func main() {

	migration.Main()

	r := gin.New()

	config := &ginSwagger.Config{
		URL: "http://localhost:8080/swagger/doc.json"} //The url pointing to API definition }

	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	r.POST("/connect/token", func(c *gin.Context) {
		authenticationcontroller.Authenticate(c)
	})

	r.POST("/payment/create", paymentcontroller.Create)
	r.POST("/user/create", usercontroller.Create)
	r.GET("/user/groups", usercontroller.Groups)

	r.Run(":8080")

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
