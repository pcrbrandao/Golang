package domain

import "github.com/jinzhu/gorm"

type Receita struct {

	gorm.Model

	Descricao string
	Foto byte
	Rendimento int

	Ocasiao Ocasiao
	IngredienteNaReceita []IngredienteNaReceita
}

// ver alimento
func (a *Receita)GetID() uint {
	return a.ID
}

// configura nome da tabela
func (Receita) TableName() string {
	return "receitas"
}