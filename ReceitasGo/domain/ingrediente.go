package domain

import "github.com/jinzhu/gorm"

type Ingrediente struct {

	gorm.Model

	Nome string

	Alimento Alimento
	IngredienteNaReceita []IngredienteNaReceita
}

func (a *Ingrediente)GetID() uint {
	return a.ID
}