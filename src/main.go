package hello_world_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Hello World endpoint
	router.GET("/api/v1/hws", func(c *gin.Context) {
		message := GetHelloWorldMessage()
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	// Start the server on port 80
	if err := router.Run(":80"); err != nil {
		panic(err)
	}
}

func GetHelloWorldMessage() string {
	return "Hello World from Go"
}
