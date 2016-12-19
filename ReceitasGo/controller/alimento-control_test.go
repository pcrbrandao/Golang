package controller

import (
	"testing"
	"Golang/ReceitasGo/model"
	"Golang/ReceitasGo/mensagem"
)

func TestAlimentoControl_Adiciona(t *testing.T) {

	alimento := model.Alimento{Descricao: "Alimento 1", GrupoAlimentar:0}

	var control AlimentoControl

	if err := control.Adiciona(&alimento); err != nil {
		t.Errorf("%s %s", mensagem.ERRO, err.Error())
		return
	}

	t.Logf("%s", mensagem.TUDOCERTO)
}
