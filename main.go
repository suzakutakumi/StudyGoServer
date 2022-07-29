package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Name      string    `json:"name" binding:"required"`
	TimeStamp time.Time `json:"timestamp"`
	Content   string    `json:"content"  binding:"required"`
}

var datas []Data

func main() {
	datas = []Data{}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", index)
	r.POST("/data", setData)
	r.GET("/data", getDatas)
	r.Run(":8080")
}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func setData(c *gin.Context) {
	var d Data
	if err := c.BindJSON(&d); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	d.TimeStamp = time.Now()
	datas = append(datas, d)
	c.Status(http.StatusOK)
}
func getDatas(c *gin.Context) {
	c.JSON(http.StatusOK, datas)
}
