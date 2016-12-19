package lib

type Sharer interface {
	SharedControl() *Sharer
}
