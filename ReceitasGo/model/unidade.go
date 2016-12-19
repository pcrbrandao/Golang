package model

import "github.com/jinzhu/gorm"

type Unidade struct {

	gorm.Model

	Descricao string
	simbolo string
}

func (a *Unidade)GetID() uint {
	return a.ID
}