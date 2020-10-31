package main

import "github.com/gin-gonic/gin"

type Data struct {
  message string
}

func main() {
  var data Data
  r := gin.Default()
  data.message = "pong"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
      "code": 200,
      "message": "ok",
      "data": data,
		})
	})
	r.Run(":3000")
}
