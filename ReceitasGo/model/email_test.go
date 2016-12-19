package model

import (
	"testing"
	"Golang/ReceitasGo/mensagem"
	"Golang/ReceitasGo/tester"
)

func TestEmailDeveSerValido(t *testing.T) {

	email := Email{}

	entrada := "pcrbrandao@gmail.com"

	if err := email.SetAddress(entrada); err != nil {
		t.Errorf(err.Error())
		return
	}

	tester.PassouEntradaValida(entrada, t)
}

func TestEmailDeveSerInvalido(t *testing.T) {

	email := Email{}

	entrada := "pedro"

	if err := email.SetAddress(entrada); err != nil {
		t.Logf("%s %q %s", mensagem.TUDOCERTO, entrada, mensagem.NAOEVALIDO)
		return
	}

	t.Errorf("%q %s %s :(", entrada, mensagem.DEVERIASER, mensagem.INVALIDO)
}

func TestEmailComAcentoDeveSerInvalido(t *testing.T) {

	mail := Email{}

	entrada := "Brandão"

	if err := mail.SetAddress(entrada); err != nil {
		tester.PassouEntradaInvalida(entrada, t)
		return
	}

	tester.FalhouEntradaValida(entrada, t)
}
