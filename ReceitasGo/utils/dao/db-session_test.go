package dao_test

import (
	"fmt"
	"testing"
	"github.com/jinzhu/gorm"
	"Golang/ReceitasGo/utils/dao"
	"Golang/ReceitasGo/utils/misc"
	"Golang/ReceitasGo/tests"
)


func ExampleMySqlPass_String() {

	my := dao.SharedDbSession()
	fmt.Println(my.String())
	// OutPut: root:root@tcp(127.0.0.1:8889)/receitas?charset=utf8&parseTime=True&loc=Local
}

func TestDoDeveSerInvalido(t *testing.T) {

	m := dao.SharedDbSession()

	entrada := "do"

	if err := m.SetUser(entrada); err != nil {
		tests.PassouEntradaInvalida(entrada, t)
		return
	}
	tests.FalhouEntradaValida(entrada, t)
}

func TestUserDeveSerValido(t *testing.T) {

	m := dao.SharedDbSession()
	entrada := "user"

	if err := m.SetUser(entrada); err != nil {
		tests.FalhouEntradaInvalida(entrada, t)
		return
	}
	tests.PassouEntradaValida(entrada, t)
}

func TestRoot12DeveSerValido(t *testing.T) {

	m := dao.SharedDbSession()

	entrada := "root12"

	if err := m.SetUser(entrada); err != nil {
		tests.FalhouEntradaInvalida(entrada, t)
		return
	}

	tests.PassouEntradaValida(entrada, t)
}

func TestMySqlPass_Db(t *testing.T) {

	m := dao.SharedDbSession()

	m.SetUser("root")

	db, err := m.OpenDb()

	if err == nil && db != nil {
		t.Logf("%s %s", misc.TUDOCERTO, misc.CONECTADO)
		return
	}

	t.Errorf("%s", misc.ERROCONECTANDO)
}

// O teste cria as tabelas. As instruções para apagar as tabelas estão comentadas.
func TestCreateTablesOnDb(t *testing.T) {

	m := dao.SharedDbSession()

	db, err := gorm.Open("mysql", m.String())

	if err != nil {
		t.Errorf("%s %s", misc.ERROCONECTANDO, err.Error())
		return
	}

	if err := m.CreateTablesOnDb(db); err != nil {
		t.Errorf("%s %s", misc.NAOPUDECRIAR, err.Error())
		return
	}

	// criadas com sucesso. Agora apaga!
	// for _,model := range dao.TABLES {
	//	db.DropTableIfExists(model)
	//}

	t.Logf("tabelas criadas com sucesso!")
}