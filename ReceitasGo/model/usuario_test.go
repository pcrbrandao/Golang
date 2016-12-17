package model

import (
	"testing"
	"Golang/ReceitasGo/utils"
	"fmt"
)

func TestSenhaSemMaiusculaDeveSerInvalida(t *testing.T) {

	usuario := Usuario{}
	entrada := "ped@146"

	if err := usuario.SetSenha(entrada); err != nil {

		utils.PassouEntradaInvalida(entrada, t)
		return
	}

	utils.FalhouEntradaValida(entrada, t)
}

func TestSenhaDeveSerValida(t *testing.T) {

	u := Usuario{}
	entrada := "Ped@146"
	cond := u.verificaSenha(entrada)

	if len(cond) > 0 {
		fmt.Printf("condicoes.... %s\n", cond)
		utils.FalhouEntradaInvalida(entrada, t)
		return
	}
	fmt.Printf("comprimento de cond... %v\n", len(cond))
	utils.PassouEntradaValida(entrada, t)
}

func TestSenhaMenorQueSeteDeveSerInvalida(t *testing.T) {

	u := Usuario{}
	entrada := "Ped@14"

	if cond := u.verificaSenha(entrada); len(cond) > 0 {
		utils.PassouEntradaInvalida(entrada, t)
		return
	}

	utils.FalhouEntradaValida(entrada, t)
}

func TestSenhaEmBrancoDeveSerInvalida(t *testing.T) {
	u := Usuario{}
	entrada := ""

	if cond := u.verificaSenha(entrada); len(cond) > 0 {
		utils.PassouEntradaInvalida(entrada, t)
		return
	}
	utils.FalhouEntradaInvalida(entrada, t)
}

func TestSenhaNumericaDeveSerInvalida(t *testing.T) {
	u := Usuario{}
	entrada := "1234567"

	if cond := u.verificaSenha(entrada); len(cond) > 0 {
		utils.PassouEntradaInvalida(entrada, t)
		return
	}
	utils.FalhouEntradaInvalida(entrada, t)
}

func TestSenhaTodasMinusculasDeveSerInvalida(t *testing.T) {
	u := Usuario{}
	entrada := "pedrocr"

	if cond := u.verificaSenha(entrada); len(cond) > 0 {
		utils.PassouEntradaInvalida(entrada, t)
		return
	}
	utils.FalhouEntradaInvalida(entrada, t)
}
