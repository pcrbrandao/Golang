package mensagem

// O MensagemaNome implementará Enum
type Enum interface {
	String() string
	Numero() int
	Mensagens() *[]string
}

// Esse é o tipo
type Mensagem uint

// essa é a lista dos tipos
const (
	TUDOCERTO Mensagem = iota
	VALIDO
	NAOEVALIDO
	DEVERIASER
	EMDESENVOLVIMENTO
	ERROCONECTANDO
	CONECTADO
	ERRO
	INVALIDO
	JAEXISTE
)

// essa é a lista das mensagens que deve seguir a ordem da lista dos tipos
const mensagens = []string{
	"tudo certo :)",
	"é válido",
	"não é válido",
	"deveria ser",
	"em desenvolvimento",
	"erro tentando conectar",
	"conectado com sucesso",
	"algo errado aqui",
	"inválido",
	"já existe",
}

// as funções que implementam a interface Enum. m é um uint (inteiro sem sinal)
func (m Mensagem) String() string {
	return mensagens[m]
}

func (m Mensagem) Numero() int {
	return int(m)
}

func (m Mensagem) Mensagens() *[]string {
	return &mensagens
}