package main

import (
	"github.com/codegangsta/martini"
	"Golang/GoWebService/models"
	"Golang/GoWebService/webservice"
)

func main() {
	martiniClassic := martini.Classic()
	guestBook := models.NewGuestBook()
	webservice.RegisterWebService(guestBook, martiniClassic)
	martiniClassic.Run()
}
