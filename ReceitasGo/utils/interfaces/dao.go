package interfaces


type Dao interface {

	Add(m *Ider) error
	Edit(m *Ider) error
	Delete(m *Ider) error
	List(m *Ider) error
}
