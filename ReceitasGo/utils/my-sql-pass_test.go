package utils

import (
	"fmt"
	"testing"
	"Golang/ReceitasGo/mensagem"
)

func ExampleMySqlPass_Param() {

	my := SharedMySqlPass()
	my.param = "parametros"
	fmt.Println(my.Param())
	// OutPut: parametros
}

func ExampleMySqlPass_Param2() {

	my := SharedMySqlPass()
	my.param = "charset=utf8&parseTime=True&loc=Local"
	fmt.Println(my.Param())
	// OutPut: charset=utf8&parseTime=True&loc=Local
}

func ExampleMySqlPass_Dbname() {

	my := SharedMySqlPass()
	fmt.Println(my.Dbname())
	// OutPut: receitas
}

func ExampleMySqlPass_User() {

	my := SharedMySqlPass()
	fmt.Println(my.User())
	// OutPut: root
}

func ExampleMySqlPass_Pass() {

	my := SharedMySqlPass()
	fmt.Println(my.Pass())
	// OutPut: root
}

func ExampleMySqlPass_String() {

	my := SharedMySqlPass()
	fmt.Println(my.String())
	// OutPut: root:root@tcp(127.0.0.1:8889)/receitas?charset=utf8&parseTime=True&loc=Local
}

func TestDoDeveSerInvalido(t *testing.T) {

	m := SharedMySqlPass()

	entrada := "do"

	if err := m.SetUser(entrada); err != nil {
		PassouEntradaInvalida(entrada, t)
		return
	}
	FalhouEntradaValida(entrada, t)
}

func TestUserDeveSerValido(t *testing.T) {

	m := SharedMySqlPass()
	entrada := "user"

	if err := m.SetUser(entrada); err != nil {
		FalhouEntradaInvalida(entrada, t)
		return
	}
	PassouEntradaValida(entrada, t)
}

func TestRoot12DeveSerValido(t *testing.T) {

	m := SharedMySqlPass()

	entrada := "root12"

	if err := m.SetUser(entrada); err != nil {
		FalhouEntradaInvalida(entrada, t)
		return
	}

	PassouEntradaValida(entrada, t)
}

func TestMySqlPass_Db(t *testing.T) {

	m := SharedMySqlPass()

	m.SetUser("root")

	db, err := m.Db()
	defer db.Close()

	if err == nil && db != nil {
		t.Logf("%s %s", mensagem.TUDOCERTO, mensagem.CONECTADO)
		return
	}

	t.Errorf(mensagem.ERROCONECTANDO)
}