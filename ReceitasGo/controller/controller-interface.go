package controller

import "Golang/ReceitasGo/model"

type Controller interface {

	Adiciona(m *model.Modelo) error
	Edita(m *model.Modelo) error
	Exclui(m *model.Modelo) error
	Lista(m *model.Modelo) error
}
