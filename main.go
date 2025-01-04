package main

import (
	httpHandler "example/hello/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/emails", httpHandler.GetEmails)
	router.Run("localhost:8080")
}
