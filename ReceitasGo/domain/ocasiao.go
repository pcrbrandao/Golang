package domain

import "github.com/jinzhu/gorm"

type Ocasiao struct {

	gorm.Model

	Descricao string

	Receitas []Receita
}

// ver alimento
func (a *Ocasiao)GetID() uint {
	return a.ID
}

// nome da tabela no banco
func (Ocasiao) TableName() string {
	return "ocasioes"
}