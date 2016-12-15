package model

import (
	"testing"
	"Golang/ReceitasGo/mensagem"
)

func TestEmail_SetEmail(t *testing.T) {

	email := Email{}

	entrada := "pcrbrandao@gmail.com"

	if err := email.SetAddress(entrada); err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("%s %q %s", mensagem.TUDOCERTO.String(), email.Address, mensagem.VALIDO.String())
}

func TestEmail_SetEmail2(t *testing.T) {

	email := Email{}

	entrada := "pedro"

	if err := email.SetAddress(entrada); err != nil {
		t.Logf("%s %q %s", mensagem.TUDOCERTO.String(), entrada, mensagem.NAOEVALIDO.String())
		return
	}

	t.Errorf("%q %s %s :(", entrada, mensagem.DEVERIASER.String(), mensagem.INVALIDO.String())
}
