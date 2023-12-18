package main

import (
	"sman1parung/config"
	cobaControllers "sman1parung/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	r.GET("/semua", cobaControllers.Index)
	r.GET("/tampil/:id", cobaControllers.Show)
	r.PUT("/ganti/:id", cobaControllers.Update)
	r.POST("/create", cobaControllers.Create)
	r.DELETE("/delete/:id", cobaControllers.Hapus)
	r.Run()
}
