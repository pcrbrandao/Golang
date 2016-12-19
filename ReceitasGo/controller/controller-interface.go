package controller

import "Golang/ReceitasGo/model"

type Controller interface {

	Adiciona(m *model.Ider) error
	Edita(m *model.Ider) error
	Exclui(m *model.Ider) error
	Lista(m *model.Ider) error
}
