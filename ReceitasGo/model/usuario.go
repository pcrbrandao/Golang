package model

import (
	"github.com/jinzhu/gorm"
)

type Usuario struct {

	// Contém ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model

	Nome string `gorm:"size:255"`
	Login string
	Email Email
}
