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

	alimento := domain.Alimento{Descricao: "Alimento 2", GrupoAlimentar:0}

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
	alimento, _ := a_dao.Find(entrada)
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
	alimento, _ := a_dao.Find(entrada)
	fmt.Printf("Foi encontrado %q", alimento)

	if alimento.Descricao == "" {
		tests.PassouEntradaInvalida(alimento.Descricao, t)
		return
	}

	tests.FalhouEntradaValida(entrada, t)
}

// Deleta um registro 'Alimento 1' existente no banco. Deve retornar erro == nil
func TestAlimentoDAO_Delete(t *testing.T) {
	entrada := domain.Alimento{Descricao:"Alimento 4", GrupoAlimentar:0}
	a_dao := dao.SharedAlimentoDAO()
	err := a_dao.Add(&entrada)

	alimento, err := a_dao.Find(entrada.Descricao)

	if err != nil {
		tests.FalhouEntradaInvalida(entrada.Descricao,t)
		return
	}
	fmt.Printf("Registro acrescentado e encontrado: %q", alimento)

	err = a_dao.Delete(alimento)
	if err != nil {
		tests.FalhouEntradaValida(entrada.Descricao, t)
		return
	}

	tests.PassouEntradaValida(entrada.Descricao, t)
}