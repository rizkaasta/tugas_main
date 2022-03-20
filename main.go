package main

import "github.com/gin-gonic/gin"
import handler "github.com/rizkaasta/tugas_handler"

func main() {
	r := gin.Default()
	//no 1
	r.GET("/item", handler.Product)
	r.GET("/id/:id/produk", handler.IdProduct)
	//no 2
	r.GET("/sign_up", handler.SignUp)
	r.GET("/sign_in", handler.SignIn)
	r.GET("/book", handler.Booking)
	r.GET("/return", handler.Return)
	r.GET("/sign_out", handler.SignOut)
	r.Run()
	

}