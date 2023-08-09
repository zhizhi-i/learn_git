package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,go")
	})
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}
