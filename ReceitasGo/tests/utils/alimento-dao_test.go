package tests

import (
	"testing"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/dao"
	"Golang/ReceitasGo/utils/misc"
)

func TestAlimentoDAO_Adiciona(t *testing.T) {

	alimento := domain.Alimento{Descricao: "Alimento 1", GrupoAlimentar:0}

	dao := dao.SharedAlimentoDAO()

	if err := dao.Add(&alimento); err != nil {
		t.Errorf("Erro tentando adicionar: %s %s", misc.ERRO, err.Error())
		return
	}

	t.Logf("%s", misc.TUDOCERTO)
}
