package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhattox/go-rest-api-example/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}