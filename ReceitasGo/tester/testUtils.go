package tester

import (
	"testing"
	"Golang/ReceitasGo/mensagem"
)

func PassouEntradaInvalida(entrada string, t *testing.T) {
	t.Logf("%s %q %s", mensagem.TUDOCERTO, entrada, mensagem.NAOEVALIDO)
}

func PassouEntradaValida(entrada string, t *testing.T) {
	t.Logf("%s %q é %s", mensagem.TUDOCERTO, entrada, mensagem.VALIDO)
}

func FalhouEntradaValida(entrada string, t *testing.T) {
	t.Errorf("%s %q %s %s", mensagem.ERRO, entrada, mensagem.DEVERIASER, mensagem.VALIDO)
}

func FalhouEntradaInvalida(entrada string, t *testing.T) {
	t.Errorf("%s %q %s %s", mensagem.ERRO, entrada, mensagem.DEVERIASER, mensagem.VALIDO)
}
