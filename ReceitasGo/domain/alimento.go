package domain

import  (
	"github.com/jinzhu/gorm"
)

type Alimento struct {

	gorm.Model

	Descricao string
	GrupoAlimentar int // criar um Enum com os tipos possíveis

	Ingrediente []Ingrediente
}

func (a *Alimento)GetID() uint {
	return a.ID
}