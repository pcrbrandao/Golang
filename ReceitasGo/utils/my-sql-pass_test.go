package utils

import (
	"fmt"
	"testing"
	"Golang/ReceitasGo/mensagem"
	"github.com/jinzhu/gorm"
	"Golang/ReceitasGo/tester"
)

func ExampleMySqlPass_Param() {

	my := SharedControl()
	my.param = "parametros"
	fmt.Println(my.Param())
	// OutPut: parametros
}

func ExampleMySqlPass_Param2() {

	my := SharedControl()
	my.param = "charset=utf8&parseTime=True&loc=Local"
	fmt.Println(my.Param())
	// OutPut: charset=utf8&parseTime=True&loc=Local
}

func ExampleMySqlPass_Dbname() {

	my := SharedControl()
	fmt.Println(my.Dbname())
	// OutPut: receitas
}

func ExampleMySqlPass_User() {

	my := SharedControl()
	fmt.Println(my.User())
	// OutPut: root
}

func ExampleMySqlPass_Pass() {

	my := SharedControl()
	fmt.Println(my.Pass())
	// OutPut: root
}

func ExampleMySqlPass_String() {

	my := SharedControl()
	fmt.Println(my.String())
	// OutPut: root:root@tcp(127.0.0.1:8889)/receitas?charset=utf8&parseTime=True&loc=Local
}

func TestDoDeveSerInvalido(t *testing.T) {

	m := SharedControl()

	entrada := "do"

	if err := m.SetUser(entrada); err != nil {
		tester.PassouEntradaInvalida(entrada, t)
		return
	}
	tester.FalhouEntradaValida(entrada, t)
}

func TestUserDeveSerValido(t *testing.T) {

	m := SharedControl()
	entrada := "user"

	if err := m.SetUser(entrada); err != nil {
		tester.FalhouEntradaInvalida(entrada, t)
		return
	}
	tester.PassouEntradaValida(entrada, t)
}

func TestRoot12DeveSerValido(t *testing.T) {

	m := SharedControl()

	entrada := "root12"

	if err := m.SetUser(entrada); err != nil {
		tester.FalhouEntradaInvalida(entrada, t)
		return
	}

	tester.PassouEntradaValida(entrada, t)
}

func TestMySqlPass_Db(t *testing.T) {

	m := SharedControl()

	m.SetUser("root")

	db, err := m.Db()
	defer db.Close()

	if err == nil && db != nil {
		t.Logf("%s %s", mensagem.TUDOCERTO, mensagem.CONECTADO)
		return
	}

	t.Errorf(mensagem.ERROCONECTANDO)
}

func TestCreateTablesOnDb(t *testing.T) {

	mysql := SharedControl()

	db, err := gorm.Open("mysql", mysql.String())
	defer db.Close()

	if err != nil {
		t.Errorf("%s %s", mensagem.ERROCONECTANDO, err.Error())
		return
	}

	if err := createTablesOnDb(db); err != nil {
		t.Errorf("%s %s", mensagem.NAOPUDECRIAR, err.Error())
		return
	}

	// criadas com sucesso. Agora apaga!
	for _,model := range TABLES {
		db.DropTableIfExists(model)
	}

	t.Logf("tabelas criadas com sucesso!")
}