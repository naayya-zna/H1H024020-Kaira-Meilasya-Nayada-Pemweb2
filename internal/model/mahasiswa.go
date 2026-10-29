package model

// Mahasiswa mendefinisikan struktur data mahasiswa.
type Mahasiswa struct {
	ID       int     `json:"id"`
	NIM      string  `json:"nim"`
	Nama     string  `json:"nama"`
	Prodi    string  `json:"prodi"`
	Angkatan int     `json:"angkatan"`
	IPK      float64 `json:"ipk"`
}

// DataMahasiswa berisi data contoh untuk API.
var DataMahasiswa = []Mahasiswa{
	{ID: 1, NIM: "H1H024001", Nama: "Andi", Prodi: "Teknik Komputer", Angkatan: 2024, IPK: 3.5},
	{ID: 2, NIM: "H1H024002", Nama: "Budi", Prodi: "Teknik Komputer", Angkatan: 2024, IPK: 3.7},
	{ID: 3, NIM: "H1H025001", Nama: "naya", Prodi: "Teknik Komputer", Angkatan: 2025, IPK: 3.8},
}
