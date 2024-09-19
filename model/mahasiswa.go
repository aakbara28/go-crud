package model

// Mahasiswa represents the student entity
type Mahasiswa struct {
	NIM     string `gorm:"primaryKey;column:nim" json:"nim"`
	Nama    string `json:"nama"`
	Umur    int    `json:"umur"`
	Jurusan string `json:"jurusan"`
}
