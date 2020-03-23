package main

import (
	"fmt"
	"os"
	
	"./controllers"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("success load env")

	router := gin.Default()
	controllers.ShortenControllerHandler(router)

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.Run(serverString)
}