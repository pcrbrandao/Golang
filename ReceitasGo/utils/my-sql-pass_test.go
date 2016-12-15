package utils

import (
	"fmt"
	"testing"
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

func TestMySqlPass_User(t *testing.T) {

	m := SharedMySqlPass()

	if err := m.SetUser("do"); err != nil {
		t.Logf("%s %s", ERRO.String(), err.Error())
		return
	}

	t.Errorf("%s usuário %q está ok...", TUDOCERTO.String(), m.User())
}

func TestMySqlPass_User2(t *testing.T) {

	m := SharedMySqlPass()

	if err := m.SetUser("user"); err != nil {
		t.Errorf("%s %s", ERRO.String(), err.Error())
		return
	}

	t.Logf("%s usuario %q está ok....", TUDOCERTO.String(), m.User())
}

func TestMySqlPass_User3(t *testing.T) {

	m := SharedMySqlPass()

	if err := m.SetUser("root12"); err != nil {
		t.Errorf("%s %s", ERRO.String(), err.Error())
		return
	}

	t.Logf("%s usuario %q está ok....", TUDOCERTO.String(), m.User())
}

func TestMySqlPass_Db(t *testing.T) {

	m := SharedMySqlPass()

	m.SetUser("root")

	if db, err := m.Db(); err == nil && db != nil {
		t.Logf("%s %s", TUDOCERTO.String(), CONECTADO.String())

		defer db.Close()
		return
	}

	t.Errorf(ERROCONECTANDO.String())
}