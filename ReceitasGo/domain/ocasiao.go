package domain

import "github.com/jinzhu/gorm"

type Ocasiao struct {

	gorm.Model

	Descricao string

	Receitas []Receita
}

func (a *Ocasiao)GetID() uint {
	return a.ID
}