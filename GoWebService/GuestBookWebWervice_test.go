package webservice

import (
	"sync"
	"fmt"
)

func ExampleGuestBook_GetPath() {

	g := &GuestBook{
		make(GuestBooksData, 0),
		new(sync.Mutex),
	}

	path := g.GetPath()
	fmt.Println(path)
	// Output: /guestbook
}

