package model

import "github.com/jinzhu/gorm"

type IngredienteNaReceita struct {

	gorm.Model

	Receita Receita
	Ingrediente Ingrediente
	Unidade Unidade

	Quantidade int
}