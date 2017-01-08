package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // traduz os parâmetros passados na url
	// Sem isso, não é possível obter os dados inseridos no formulário login.html
	fmt.Println(r.Form) // imprime informação no lado do servidor
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ", "))
	}
	fmt.Fprintf(w, "Olá você!") // escreve dados para response
}

func login (w http.ResponseWriter, r *http.Request) {
	fmt.Println("mothod: ", r.Method) // obtém método request
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// parte lógica do log in
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("index")
	if err != nil {
		p = &Page{Title:"Erro tentando ler index.html"}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler) // o método sayhelloName será executado na raíz '/'
	http.HandleFunc("/login", login)   // e o método login em '/login'

	err := http.ListenAndServe(":9090", nil) // servidor funcionando em localhost:9090
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}