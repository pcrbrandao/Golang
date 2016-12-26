package domain_test

import (
	"testing"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"Golang/ReceitasGo/tests"
)

func TestEmailDeveSerValido(t *testing.T) {

	email := domain.Email{}

	entrada := "pcrbrandao@gmail.com"

	if err := email.SetAddress(entrada); err != nil {
		t.Errorf(err.Error())
		return
	}

	tests.PassouEntradaValida(entrada, t)
}

func TestEmailDeveSerInvalido(t *testing.T) {

	email := domain.Email{}

	entrada := "pedro"

	if err := email.SetAddress(entrada); err != nil {
		t.Logf("%s %q %s", misc.TUDOCERTO, entrada, misc.NAOEVALIDO)
		return
	}

	t.Errorf("%q %s %s :(", entrada, misc.DEVERIASER, misc.INVALIDO)
}

func TestEmailComAcentoDeveSerInvalido(t *testing.T) {

	mail := domain.Email{}

	entrada := "Brandão"

	if err := mail.SetAddress(entrada); err != nil {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}

	tests.FalhouEntradaValida(entrada, t)
}
