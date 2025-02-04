package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	r.Use(cors.New(config))
	r.GET("/")
	r.POST("/test", test)
	r.Run(":8080")
}

func root(ctx *gin.Context) {
	fmt.Println("hello")
	ctx.Writer.WriteString("root")
}

func test(ctx *gin.Context) {
	fmt.Println("hello2")
	ctx.Writer.WriteString("hello")
}

