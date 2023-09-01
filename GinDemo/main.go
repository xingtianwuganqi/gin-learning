package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("123")

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "pong"})
	})

	// AsciiJSON
	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言", "tag": "<br>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	// JSONP
	r.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{"title": "posts/index"})
	})

	r.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{"title": "users/index"})
	})

	r.GET("/some/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello,world"})
	})

	r.GET("/more/json", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "hello, world"
		msg.Age = 19
		context.JSON(http.StatusOK, msg)
	})

	// xml格式的数据
	r.GET("/more/xml", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "hello, world"
		msg.Age = 19
		context.XML(http.StatusOK, msg)
	})

	// YAML格式的数据
	r.GET("/more/yaml", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "hello, world"
		msg.Age = 19
		context.YAML(http.StatusOK, msg)
	})

	// 获取querystring的参数
	r.GET("/user/search", func(c *gin.Context) {
		userName := c.DefaultQuery("username", "小王子") // 有个默认值
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{"message": "ok", "username": userName, "age": age})
	})

	// 获取form参数
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 获取json参数
	r.POST("/user/json", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "fail"})
			return
		}
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	// 获取path参数
	r.GET("/user/path/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{"message": "ok", "username": username, "age": address})
	})

	r.Run()

}
