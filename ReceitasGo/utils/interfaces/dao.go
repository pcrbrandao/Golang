package interfaces


type Dao interface {

	Adiciona(m *Ider) error
	Edita(m *Ider) error
	Exclui(m *Ider) error
	Lista(m *Ider) error
}
