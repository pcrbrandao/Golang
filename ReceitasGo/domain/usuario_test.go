package domain_test

import (
	"testing"
	"fmt"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/tests"
)

func getUsuario() domain.Usuario {
	return domain.Usuario{}
}

func TestSenhaSemMaiusculaDeveSerInvalida(t *testing.T) {

	usuario := getUsuario()
	entrada := "ped@146"

	if err := usuario.SetSenha(entrada); err != nil {

		tests.PassouEntradaInvalida(entrada, t)
		return
	}

	tests.FalhouEntradaValida(entrada, t)
}

func TestSenhaDeveSerValida(t *testing.T) {

	u := getUsuario()
	entrada := "Ped@146"
	cond := u.VerificaSenha(entrada)

	if len(cond) > 0 {
		fmt.Printf("condicoes.... %s\n", cond)
		tests.FalhouEntradaInvalida(entrada, t)
		return
	}
	fmt.Printf("comprimento de cond... %v\n", len(cond))
	tests.PassouEntradaValida(entrada, t)
}

func TestSenhaMenorQueSeteDeveSerInvalida(t *testing.T) {

	u := getUsuario()
	entrada := "Ped@14"

	if cond := u.VerificaSenha(entrada); len(cond) > 0 {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}

	tests.FalhouEntradaValida(entrada, t)
}

func TestSenhaEmBrancoDeveSerInvalida(t *testing.T) {
	u := getUsuario()
	entrada := ""

	if cond := u.VerificaSenha(entrada); len(cond) > 0 {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}
	tests.FalhouEntradaInvalida(entrada, t)
}

func TestSenhaNumericaDeveSerInvalida(t *testing.T) {
	u := getUsuario()
	entrada := "1234567"

	if cond := u.VerificaSenha(entrada); len(cond) > 0 {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}
	tests.FalhouEntradaInvalida(entrada, t)
}

func TestSenhaTodasMinusculasDeveSerInvalida(t *testing.T) {
	u := getUsuario()
	entrada := "pedrocr"

	if cond := u.VerificaSenha(entrada); len(cond) > 0 {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}
	tests.FalhouEntradaInvalida(entrada, t)
}
