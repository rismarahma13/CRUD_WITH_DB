package main

// Struktur data untuk Siswa
type Siswa struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Umur  int    `json:"umur"`
	Kelas string `json:"kelas"`
}
