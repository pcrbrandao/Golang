package model

import  (
	"github.com/jinzhu/gorm"

)

type Alimento struct {

	gorm.Model

	Descricao string
	GrupoAlimentar int // criar um Enum com os tipos possíveis

	Ingrediente []Ingrediente
}

func SetDescricao(d string) error {
	// hashedDescricao, err := bcrypt.
}