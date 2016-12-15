package model

type Email struct {
	ID int
	UsuarioID int `gorm:"index"` // Foreign key
	Email string `gorm:"type:varchar(100);unique_index"`
	Inscrito bool
}