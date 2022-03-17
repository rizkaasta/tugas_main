package main

import "github.com/gin-gonic/gin"
import handler "github.com/rizkaasta/tugas_handler"

func main() {
	r := gin.Default()
	r.GET("/item", handler.Product)
	r.GET("/id/:id/produk", handler.IdProduct)
	r.Run()
}