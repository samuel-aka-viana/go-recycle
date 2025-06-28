package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	r := gin.Default()

	// initialize Router
	InitializeRoutes(r)

	// run the server
	_ = r.Run(":8080")
}
