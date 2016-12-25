package domain

import "github.com/jinzhu/gorm"

type IngredienteNaReceita struct {

	gorm.Model

	Receita Receita
	Ingrediente Ingrediente
	Unidade Unidade

	Quantidade int
}

func (a *IngredienteNaReceita)GetID() uint {
	return a.ID
}