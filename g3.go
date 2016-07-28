package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kixpanganiban/g3/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/blog/:name", handlers.EntryHandler)
	router.Run()
}
