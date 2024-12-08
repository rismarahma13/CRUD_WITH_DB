package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler untuk membuat siswa baru
func createSiswa(c echo.Context) error {
	siswa := new(Siswa)
	if err := c.Bind(siswa); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := "INSERT INTO siswa (nama, umur, kelas) VALUES (?, ?, ?)"
	_, err := db.Exec(query, siswa.Nama, siswa.Umur, siswa.Kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, siswa)
}

// Handler untuk mengambil semua siswa
func getAllSiswa(c echo.Context) error {
	rows, err := db.Query("SELECT id, nama, umur, kelas FROM siswa")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	siswaList := []Siswa{}
	for rows.Next() {
		var siswa Siswa
		if err := rows.Scan(&siswa.ID, &siswa.Nama, &siswa.Umur, &siswa.Kelas); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		siswaList = append(siswaList, siswa)
	}

	return c.JSON(http.StatusOK, siswaList)
}

// Handler untuk mengambil siswa berdasarkan ID
func getSiswaByID(c echo.Context) error {
	id := c.Param("id")
	var siswa Siswa
	query := "SELECT id, nama, umur, kelas FROM siswa WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&siswa.ID, &siswa.Nama, &siswa.Umur, &siswa.Kelas)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Siswa not found")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, siswa)
}

// Handler untuk mengupdate siswa berdasarkan ID
func updateSiswa(c echo.Context) error {
	id := c.Param("id")
	siswa := new(Siswa)
	if err := c.Bind(siswa); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := "UPDATE siswa SET nama = ?, umur = ?, kelas = ? WHERE id = ?"
	_, err := db.Exec(query, siswa.Nama, siswa.Umur, siswa.Kelas, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, siswa)
}

// Handler untuk menghapus siswa berdasarkan ID
func deleteSiswa(c echo.Context) error {
	id := c.Param("id")
	query := "DELETE FROM siswa WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Siswa deleted")
}
