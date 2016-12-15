package webservice

import "github.com/codegangsta/martini"

func main() {
	martiniClassic := martini.Classic()
	guestBook := NewGuestBook()
	RegisterWebService(guestBook, martiniClassic)
	martiniClassic.Run()
}
