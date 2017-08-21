package main

import (
	"testing"
	"reflect"
)

func TestArquivoIn_MaxCaractPorLinha(t *testing.T) {
	arquivo := NewArquivo()

	esperado := 15
	obtido, err := arquivo.maxCaracteresPorLinha()
	check(err)

	if obtido == esperado {
		t.Logf("Tudo certo! %v == %v", obtido, esperado)
		return
	}

	t.Errorf("esperado: %v", esperado)
	t.Errorf("obtido: %v", obtido)
}

func TestArquivoInLinhas(t *testing.T) {

	arquivo := NewArquivo()

	esperado := "Now is the \n\nwinter of our \n\ndiscontent, \n\nMade glorious \n\nsummer by this \n\nsun of York. \n\n"
	obtido := arquivo.texto()

	if reflect.DeepEqual(esperado, obtido) {
		t.Logf("Tudo certo! %v == %v", obtido, esperado)
		return
	}

	t.Errorf("esperado: %v", esperado)
	t.Errorf("obtido: %v", obtido)
}
