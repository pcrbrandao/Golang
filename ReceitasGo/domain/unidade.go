package domain

import "github.com/jinzhu/gorm"

type Unidade struct {

	gorm.Model

	Descricao string
	simbolo string
}
// ver alimento
func (a *Unidade)GetID() uint {
	return a.ID
}