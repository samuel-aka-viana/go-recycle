package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	_ = router.Run(":8080")
}
