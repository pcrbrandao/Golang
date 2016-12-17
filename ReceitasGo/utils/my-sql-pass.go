package utils

import (
	"strings"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"regexp"
	"net"
	"Golang/ReceitasGo/mensagem"
)

// Usado para validar strings
// alfanumerico.MatchString() bool faz a análise
var alfanumerico = regexp.MustCompile(`[0-9A-Za-z]$`)

// Contém os parâmetros para conexao ao db
// os campos são auto-explicativos
// estão privados pois serão validados por sets e gets
type mySqlPass struct {

	user string
	pass string
	address net.IP
	dbName string
	param string

	db *gorm.DB
}

// O singleton para a struct
var sharedInstance *mySqlPass
var once sync.Once

// O construtor singleton
func SharedMySqlPass() *mySqlPass {
	once.Do(func() {
		sharedInstance = &mySqlPass{}
	})
	return sharedInstance
}

// inicializa User se for nulo
func (m *mySqlPass)User() string {
	if m.user == "" {
		return "root"
	}
	return m.user
}

// Valida a entrada de User
func (m *mySqlPass)SetUser(u string) error {

	var MENSERRO = fmt.Sprintf("%s valor %q %s", mensagem.ERRO, u, mensagem.NAOEVALIDO)

	var alfanumerico = regexp.MustCompile(`[0-9A-Za-z]$`)

	if eAlfanumerico := alfanumerico.MatchString(u); !eAlfanumerico {
		return fmt.Errorf(MENSERRO)
	}

	if l := len(u); l == 0 || l < 3 || l > 40 {
		return fmt.Errorf(MENSERRO)
	}

	m.user = u
	return nil
}

// inicializa Pass se for nulo
func (m *mySqlPass)Pass() string {
	if m.pass == "" {
		return "root"
	}
	return m.pass
}

// o banco de dados padrão é o receitas
func (m *mySqlPass)Dbname() string {
	if m.dbName == "" {
		return "receitas"
	}
	return m.dbName
}

// parâmetros default do GORM
func (m *mySqlPass)Param() string {
	if m.param == "" {
		return "charset=utf8&parseTime=True&loc=Local"
	}
	return m.param
}

// o endereço do servidor. Será a máquina local por padrão. Não aceitou localhost.
func (m *mySqlPass)Address() string {
	if m.address == nil {
		return "127.0.0.1:8889"
	}
	return m.address.String()
}

// O Address é do tipo net.IP, por isso ele não precisa ser validado aqui.
func (m *mySqlPass)SetAddress(ip net.IP) {
	m.address = ip
}

// Obtém a string no formato que o gorm.Open aceita.
func (m *mySqlPass) String() string {

	fields := []string{m.User(), ":", m.Pass(), "@tcp(", m.Address(), ")/", m.Dbname(), "?", m.Param()}

	return strings.Join(fields, "")
}

//
func (m *mySqlPass) Db() (*gorm.DB, error) {

	if m.db != nil {
		return m.db, nil
	}

	db, err := gorm.Open("mysql", m.String())

	if err != nil {
		fmt.Printf("%s %s", mensagem.ERROCONECTANDO, err.Error())
		return nil, err
	}

	println(mensagem.CONECTADO)
	return db, nil
}