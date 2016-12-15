package webservice

import (
	"sync"
	"fmt"
)

// represents a single entry in a Guest Book
type GuestBookEntry struct {
	Id int
	Email string
	Title string
	Content string
}

// acrescentei aqui para poder criar um método
type GuestBooksData []*GuestBookEntry

// represents a Guest Book instance. Guarda a associação GuestBookEntries
type GuestBook struct {
	// guestBookData []*GuestBookEntry
	GuestBooksData
	Mutex *sync.Mutex
}

// return a new GuestBook instance
func NewGuestBook() *GuestBook {
	return &GuestBook{
		// make([]*GuestBookEntry, 0),
		make(GuestBooksData, 0),
		new(sync.Mutex),
	}
}

// acrescenta um registro GuestBookEntry com os dados fornecidos.
func (guestBook *GuestBook) AddEntry(email, title, content string) int {

	// Adquire a tranca e certifica-se de que estará liberada
	guestBook.Mutex.Lock()
	defer guestBook.Mutex.Unlock()

	// Obtém um id para essa entrada
	newId := len(guestBook.GuestBooksData)

	// Cria uma nova entrada com os dados fornecidos e o newId computado
	newEntry := &GuestBookEntry{
		newId,
		email,
		title,
		content,
	}

	// Acrescenta a entrada ao Guest Book.
	guestBook.GuestBooksData = append(guestBook.GuestBooksData, newEntry)

	// Retorna o Id para uma nova entrada
	return newId
}

// remove um registro a partir de um id. Retorna nil caso sucesso ou
// um erro específico em caso de falha
func (guestBook *GuestBook) RemoveEntry(id int) error {
	// adquire a chave e certifica-se que ela seja liberada
	guestBook.Mutex.Lock()
	defer guestBook.Mutex.Unlock()

	// Verifica se o id é válido
	if id < 0 || id >= len(guestBook.GuestBooksData) ||
		guestBook.GuestBooksData[id] == nil {
		return fmt.Errorf("id %d invalido...", id)
	}

	// se o id é válido continua aqui
	// atribui nil para o registro, segundo o tutorial, mas eu vou tentar excluir
	// guestBook.GuestBooksData[id] = nil
	guestBook.RemoveID(id) // ok, funcionou!
	println("id %s removido", id)

	return nil
}

// criei esse método aqui pois acho que ficou melhor.
// func RemoveID(slice []*GuestBookEntry, id int) {
func (slice *GuestBooksData) RemoveID(id int) {

	s := *slice
	s = append(s[:id], s[id+1:]...)
	*slice = s
}

// retorna um registro por id ou um erro caso não encontre
func (g *GuestBook) GetEntry(id int) (*GuestBookEntry, error) {
	// verifica se temos um id válido
	if id < 0 || id >= len(g.GuestBooksData) ||
		g.GuestBooksData[id] == nil {

		return nil, fmt.Errorf("id inválido")
	} else {
		return g.GuestBooksData[id], nil
	}
}

// retorna todas as entidades no Guest Book
func (g *GuestBook) GetAllEntries() []*GuestBookEntry {
	return g.GuestBooksData
}

// remove todos os registros
func (g *GuestBook) RemoveAllEntries() {
	// certifica-se que está livre
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	// atribui um GuestBookEntry novo
	g.GuestBooksData = GuestBooksData{}
}