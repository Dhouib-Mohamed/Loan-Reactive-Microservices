package main

import (
	"github.com/gin-gonic/gin"
	"ocr-extraction/utils"
)

func main() {
	// Do something
	r := gin.Default()
	r.POST("/notif", func(c *gin.Context) {
		response := utils.OCR(c)
		c.JSON(200, gin.H{
			"message":  "Notification sent",
			"response": response,
		})
	})
	r.Run(":8000")
}
