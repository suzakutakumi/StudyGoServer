package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	TimeStamp time.Time `json:"timestamp"`
	Content   string    `json:"content"  binding:"required"`
}

var datas []Data
var cnt int

func main() {
	datas = []Data{}
	cnt = 0
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", index)
	r.POST("/data", setData)
	r.GET("/data", getData)
	r.DELETE("/data", deleteData)
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
	d.Id = cnt
	cnt++
	datas = append(datas, d)
	c.Status(http.StatusOK)
}
func getData(c *gin.Context) {
	idstr := c.Query("id")
	if idstr == "" {
		c.JSON(http.StatusOK, datas)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var d Data
	flg := false
	for _, v := range datas {
		if v.Id == id {
			d = v
			flg = true
		}
	}
	if !flg {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, d)
}
func deleteData(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	flg := false
	for i, v := range datas {
		if v.Id == id {
			// i番目のデータを削除
			datas = append(datas[:i], datas[i+1:]...)
			flg = true
			break
		}
	}
	if !flg {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
