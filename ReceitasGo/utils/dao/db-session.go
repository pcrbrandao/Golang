package dao

import (
	"strings"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"net"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/interfaces"
	"Golang/ReceitasGo/utils/misc"
)

var TABLES = []interfaces.Ider{
	&domain.Alimento{},
	&domain.Email{},
	&domain.Ingrediente{},
	&domain.IngredienteNaReceita{},
	&domain.Ocasiao{},
	&domain.Receita{},
	&domain.Unidade{},
	&domain.Usuario{} }

// Contém os parâmetros para conexao ao db
// os campos são auto-explicativos
// estão privados pois serão validados por sets e gets
type dbSession struct {

	user string
	pass string
	address net.IP
	dbName string
	param string

	db *gorm.DB
}

// O singleton para a struct
var sharedDbSession *dbSession
var onceDbSession sync.Once

// O construtor singleton
func SharedDbSession() *dbSession {

	onceDbSession.Do(func() {
		sharedDbSession = &dbSession{}
	})
	return sharedDbSession
}

// inicializa User se for nulo
func (m *dbSession)User() string {
	if m.user == "" {
		return "root"
	}
	return m.user
}

// Valida a entrada de User
func (m *dbSession)SetUser(u string) error {

	var MENSERRO = fmt.Sprintf("%s valor %q %s", misc.ERRO, u, misc.NAOEVALIDO)

	if eAlfanumerico := misc.ALFANUMERICO.MatchString(u); !eAlfanumerico {
		return fmt.Errorf(MENSERRO)
	}

	if l := len(u); l == 0 || l < 3 || l > 40 {
		return fmt.Errorf(MENSERRO)
	}

	m.user = u
	return nil
}

// inicializa Pass se for nulo
func (m *dbSession)Pass() string {
	if m.pass == "" {
		return "root"
	}
	return m.pass
}

// o banco de dados padrão é o receitas
func (m *dbSession)Dbname() string {
	if m.dbName == "" {
		return "receitas"
	}
	return m.dbName
}

// parâmetros default do GORM
func (m *dbSession)Param() string {
	if m.param == "" {
		return "charset=utf8&parseTime=True&loc=Local"
	}
	return m.param
}

func (m *dbSession)SetParam(param string) {
	m.param = param
}

// o endereço do servidor. Será a máquina local por padrão. Não aceitou localhost.
func (m *dbSession)Address() string {
	if m.address == nil {
		return "127.0.0.1:8889"
	}
	return m.address.String()
}

// O Address é do tipo net.IP, por isso ele não precisa ser validado aqui.
func (m *dbSession)SetAddress(ip net.IP) {
	m.address = ip
}

// Obtém a string no formato que o gorm.Open aceita.
func (m *dbSession) String() string {

	fields := []string{m.User(), ":", m.Pass(), "@tcp(", m.Address(), ")/", m.Dbname(), "?", m.Param()}

	return strings.Join(fields, "")
}

// retorna um db válido
func (m *dbSession) Db() (*gorm.DB, error) {

	if m.db != nil {
		return m.db, nil
	}

	db, err := gorm.Open("mysql", m.String())

	if err != nil {
		fmt.Printf("%s %s", misc.ERROCONECTANDO, err.Error())
		return nil, err
	}

	if err := CreateTablesOnDb(db); err != nil {
		return nil, err
	}

	println(misc.CONECTADO)
	return db, nil
}

// cria tabelas no db
func CreateTablesOnDb(db *gorm.DB) error {

	for _,model := range TABLES {
		if err := db.AutoMigrate(model).Error; err != nil {
			return fmt.Errorf("%s %s", misc.NAOPUDECRIAR, err.Error())
		}
	}

	return nil
}