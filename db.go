package main

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Inisialisasi database dan buat tabel siswa
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS siswa (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nama TEXT,
        umur INTEGER,
        kelas TEXT
    );`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
