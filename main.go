package main

import (
	"fmt"
	"os"
	
	"./controllers"
	"./docs"
	
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Shorten URL Go
// @version 1.0
// @description Simple shortener URL.
// @termsOfService http://swagger.io/terms/

// @contact.name Dicky Wijaya
// @contact.url http://www.dickywijayaa.com
// @contact.email dw_authorized@yahoo.co.id

// @securityDefinitions.basic BasicAuth
func main() {
	godotenv.Load()
	fmt.Println("success load env")

	router := gin.Default()
	controllers.ShortenControllerHandler(router)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/documentations/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.Run(serverString)
}