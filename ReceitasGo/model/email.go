package model

import (
	"regexp"
	"fmt"
	"Golang/ReceitasGo/mensagem"
)

type Email struct {
	ID int
	UsuarioID int `gorm:"index"` // Foreign key
	Address string `gorm:"type:varchar(100);unique_index"`
	Inscrito bool
}

func (e *Email) SetAddress(address string) error {

	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if Re.MatchString(address) {
		e.Address = address
		return nil
	}

	return fmt.Errorf("%q %s", address, mensagem.NAOEVALIDO)
}