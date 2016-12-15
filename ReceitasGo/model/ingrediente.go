package model

import "github.com/jinzhu/gorm"

type Ingrediente struct {

	gorm.Model

	Nome string

	Alimento Alimento
	IngredienteNaReceita []IngredienteNaReceita
}
