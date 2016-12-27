package dao

import (
	"strings"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
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
	address string
	dbName string
	param string

	// não preciso armazenar isso
	// db *gorm.DB
}

func newDbSession() *dbSession {
	d := dbSession{user:"root", pass:"root",
		address:"127.0.0.1:8889", dbName:"receitas",
		param:"charset=utf8&parseTime=True&loc=Local"}
	return &d
}

// O singleton para a struct
var sharedDbSession *dbSession
var onceDbSession sync.Once

// O construtor singleton
func SharedDbSession() *dbSession {

	onceDbSession.Do(func() {
		sharedDbSession = newDbSession()
	})
	return sharedDbSession
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

func (m *dbSession)SetParam(param string) {
	m.param = param
}

// O Address é do tipo net.IP, por isso ele não precisa ser validado aqui.
func (m *dbSession)SetAddress(address string) {
	m.address = address
}

// Obtém a string no formato que o gorm.Open aceita.
func (m *dbSession) String() string {

	fields := []string{m.user, ":", m.pass, "@tcp(", m.address, ")/", m.dbName, "?", m.param}

	return strings.Join(fields, "")
}

// retorna um db válido
func (m *dbSession) OpenDb() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", m.String())

	if err != nil {
		fmt.Printf("%s %s", misc.ERROCONECTANDO, err.Error())
		return nil, err
	}

	if err := m.CreateTablesOnDb(db); err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", misc.CONECTADO)
	return db, nil
}

// cria tabelas no db com o nome acrescido do 's' no final
// as nomes incorretos devem estar configurados na struc em domain
func (m *dbSession) CreateTablesOnDb(db *gorm.DB) error {
	db.SingularTable(false)

	for _,model := range TABLES {
		if err := db.AutoMigrate(model).Error; err != nil {
			return fmt.Errorf("%s %s", misc.NAOPUDECRIAR, err.Error())
		}
	}

	return nil
}