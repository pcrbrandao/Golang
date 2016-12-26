package dao_test

import (
	"testing"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/dao"
	"Golang/ReceitasGo/utils/misc"
	"fmt"
	"Golang/ReceitasGo/tests"
)

// Tenta adicionar. Passa se não retornar erro.
func TestAlimentoDAO_Add(t *testing.T) {

	alimento := domain.Alimento{Descricao: "Alimento 1", GrupoAlimentar:0}

	a_dao := dao.SharedAlimentoDAO()

	if err := a_dao.Add(&alimento); err != nil {
		t.Errorf("Erro tentando adicionar: %s %s", misc.ERRO, err.Error())
		return
	}

	t.Logf("%s", misc.TUDOCERTO)
}

// Considerando que existe um 'Alimento 1' no banco deve passar.
func TestAlimentoDAO_Find(t *testing.T) {

	entrada := "Alimento 1"

	a_dao := dao.SharedAlimentoDAO()
	alimento := a_dao.Find(entrada)
	fmt.Printf("Foi encontrado %q", alimento.Descricao)

	if alimento.Descricao == entrada {
		tests.PassouEntradaValida(entrada, t)
		return
	}

	tests.FalhouEntradaInvalida(alimento.Descricao, t)
}

// Considerando que não existe um 'Alimento 2' no banco, não deve encontrar.
func TestAlimentoDAO_Find2(t *testing.T) {
	entrada := "Alimento 2"
	a_dao := dao.SharedAlimentoDAO()
	alimento := a_dao.Find(entrada)
	fmt.Printf("Foi encontrado %q", alimento)

	if alimento.Descricao == "" {
		tests.PassouEntradaInvalida(alimento.Descricao, t)
		return
	}

	tests.FalhouEntradaValida(entrada, t)
}