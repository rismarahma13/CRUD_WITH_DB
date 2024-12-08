package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	initDB() // Inisialisasi database

	e := echo.New()

	// Routes
	e.POST("/siswa", createSiswa)
	e.GET("/siswa", getAllSiswa)
	e.GET("/siswa/:id", getSiswaByID)
	e.PUT("/siswa/:id", updateSiswa)
	e.DELETE("/siswa/:id", deleteSiswa)

	e.Logger.Fatal(e.Start(":8080"))
}
