package tests

import (
	"testing"
	"Golang/ReceitasGo/utils/misc"
)

func PassouEntradaInvalida(entrada string, t *testing.T) {
	t.Logf("%s %q %s", misc.TUDOCERTO, entrada, misc.NAOEVALIDO)
}

func PassouEntradaValida(entrada string, t *testing.T) {
	t.Logf("%s %q é %s", misc.TUDOCERTO, entrada, misc.VALIDO)
}

func FalhouEntradaValida(entrada string, t *testing.T) {
	t.Errorf("%s %q %s %s", misc.ERRO, entrada, misc.DEVERIASER, misc.VALIDO)
}

func FalhouEntradaInvalida(entrada string, t *testing.T) {
	t.Errorf("%s %q %s %s", misc.ERRO, entrada, misc.DEVERIASER, misc.VALIDO)
}
