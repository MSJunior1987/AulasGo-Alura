// main.go
package main

import "github.com/gin-gonic/gin"

func SetupRotas() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}

func main() {
	r := SetupRotas()
	r.Run(":8080")
}
