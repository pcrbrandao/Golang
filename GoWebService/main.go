package main

import (
	"github.com/codegangsta/martini"
	"Golang/GoWebService/model"
	"Golang/GoWebService/webservice"
)

func main() {
	martiniClassic := martini.Classic()
	guestBook := model.NewGuestBook()
	webservice.RegisterWebService(guestBook, martiniClassic)
	martiniClassic.Run()
}
