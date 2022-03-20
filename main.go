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
	//no 3
	r.GET("/anggota/:no", handler.Anggota)
	r.GET("/buku/:no", handler.Buku)
	r.GET("/jurnal/:no", handler.Jurnal)
	r.GET("/pengarang/:pengarang", handler.Pengarang)
	r.GET("/judul/:judul", handler.Judul)
	//POST
	r.POST("/mahasiswa", handler.DataMahasiswa)
	r.Run()
}