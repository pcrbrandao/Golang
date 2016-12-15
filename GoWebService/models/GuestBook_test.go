package models

import (
	"fmt"
	"testing"
	"reflect"
)

func GuestBookTest() *GuestBook {

	guestBook := NewGuestBook()
	guestBook.AddEntry("pcrbrandao@gmail.com", "Título", "Conteúdo")

	return guestBook
}

func GuestBookTest2() *GuestBook {

	guestBook := NewGuestBook()
	guestBook.AddEntry("pcrbrandao@gmail.com", "Título", "Conteúdo")
	guestBook.AddEntry("pcrbrandao@hotmail.com", "Hotmail", "Microsoft")

	return guestBook
}

func ExampleGuestBook_AddEntry() {

	guestBook := GuestBookTest()

	fmt.Println(guestBook.GuestBooksData[0], len(guestBook.GuestBooksData))
	// Output:
	// &{0 pcrbrandao@gmail.com Título Conteúdo} 1
}

func ExampleGuestBooksData_RemoveID() {

	guestBook := GuestBookTest()
	guestBook.GuestBooksData.RemoveID(0)

	fmt.Println(len(guestBook.GuestBooksData))
	// Output: 0
}

func ExampleGuestBooksData_RemoveID_second() {

	guestBook := GuestBookTest2()
	guestBook.GuestBooksData.RemoveID(0)

	fmt.Println(len(guestBook.GuestBooksData))
	// Output: 1
}

func TestGuestBook_RemoveEntry(t *testing.T) {
	g := NewGuestBook()
	retorno := "id -3 invalido..."
	remove := g.RemoveEntry(-3)

	if remove.Error() != retorno {
		t.Errorf("esperado: %s", retorno)
		t.Errorf("retornado: %s", remove)
	} else {
		t.Log("Tudo certo!")
	}
}

func TestGuestBook_GetEntry(t *testing.T) {
	g := GuestBookTest2()
	esperado := &GuestBookEntry{0, "pcrbrandao@gmail.com", "Título", "Conteúdo"}
	obtido,_ := g.GetEntry(0)

	if reflect.DeepEqual(esperado, obtido) {
		t.Logf("Tudo certo! %v == %v", obtido, esperado)
	} else {
		t.Errorf("esperado: %v", esperado)
		t.Errorf("obtido: %v", obtido)
	}
}

func TestGuestBook_GetAllEntries(t *testing.T) {
	guestBook := GuestBookTest()
	guestBookData := guestBook.GetAllEntries()
	obtido := len(guestBookData)

	if obtido == 1 {
		t.Logf("Tudo certo! %s", guestBookData)
	} else {
		t.Errorf("esperado: %d", 1)
		t.Errorf("obtido: %d", obtido)
	}
}

func TestGuestBook_RemoveAllEntries(t *testing.T) {
	guestBook := GuestBookTest2()
	guestBook.RemoveAllEntries()

	obtido := len(guestBook.GuestBooksData)

	if obtido != 0 {
		t.Errorf("esperado: %d", 0)
		t.Errorf("obtido: %d", obtido)
	} else {
		t.Logf("Tudo certo! obtido: %d", obtido)
	}
}
