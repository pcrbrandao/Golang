package controller

type Controller interface {

	Adiciona(objeto interface{}) error
	Edita(objeto interface{}) error
	Exclui(objeto interface{}) error
	Lista(objeto interface{}) error
}
