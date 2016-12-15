package model

import "github.com/jinzhu/gorm"

type Receita struct {

	gorm.Model

	Descricao string
	Foto byte
	Rendimento int

	Ocasiao Ocasiao
	IngredienteNaReceita []IngredienteNaReceita
}