package model

import "github.com/jinzhu/gorm"

type Ocasiao struct {

	gorm.Model

	Descricao string

	Receitas []Receita
}
