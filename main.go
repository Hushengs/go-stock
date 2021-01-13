package main

import (
	"go-stock/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("view/**/*.html")
	router.Routes(r)
	r.Run()
}
