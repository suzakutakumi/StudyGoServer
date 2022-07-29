package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.Run(":8080")
}
func index(c *gin.Context) {
	c.String(200, "Hello World")
}
