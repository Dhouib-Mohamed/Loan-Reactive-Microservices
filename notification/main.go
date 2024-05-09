package main

import (
	"github.com/gin-gonic/gin"
	"notification/utils"
)

func main() {
	r := gin.Default()
	r.POST("/notif", func(c *gin.Context) {
		// Read and Send Notification
		utils.SendNotif(c)
		c.JSON(200, gin.H{
			"message": "Notification sent",
		})
	})
	r.Run(":8000")
}
