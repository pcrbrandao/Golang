package webservice

import (
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

// Implementa webservice.GetPath
func (g *GuestBook) GetPath() string {

	// Associa esse serviço com http://host:port/guestbook
	return "/guestbook"
}

// Implementa webservice.WebDelete
func (g *GuestBook) WebDelete(params martini.Params) (int, string) {

	length := len(g.GuestBooksData)

	if length <= 0 {
		return http.StatusNotFound, "não há registros a serem apagados"
	}

	if length > 0 && len(params) == 0 {
		// não tem parâmetros. Remove todos os registros
		g.RemoveAllEntries()

		return http.StatusOK, "registros apagados"
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Id não é um número
		return http.StatusBadRequest, "id inválido"
	}

	// se passou das condições acima remova id
	err = g.RemoveEntry(id)
	if err != nil {
		// registro não encontrado
		return http.StatusNotFound, "registro não encontrado"
	}

	return http.StatusOK, "registro apagado"
}

// Implementa webservice.WebGet
func (g *GuestBook) WebGet(params martini.Params) (int, string) {

	// Não há parâmetros? Retorna todos os registros em JSON
	if len(params) == 0 {
		encodedEntries, err := json.Marshal(g.GetAllEntries())
		if err != nil {
			// Falha codificando registros
			return http.StatusInternalServerError, "Não pude converter dados para JSON"
		}

		// Se os registros estão ok
		return http.StatusOK, string(encodedEntries)
	}

	// Se há parametros continua aqui
	// Converte id para inteiro
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Id não é um número
		return http.StatusBadRequest, "id inválido"
	}

	// O id é um inteiro válido. Tente obter o registro
	entry, err := g.GetEntry(id)
	if err != nil {
		// Registro não encontrado
		return http.StatusNotFound, "registro não encontrado"
	}

	// O registro existe. Converta-o para jSON
	encodedEntry, err := json.Marshal(entry)
	if err != nil {
		// o registro não pode ser convertido
		return http.StatusInternalServerError, "Não pude converter dados para JSON"
	}

	// Registro encontrado e convertido para JSON com sucesso
	return http.StatusOK, string(encodedEntry)
}

// Implementa webservice.WebPost
func (g *GuestBook) WebPost(params martini.Params, req *http.Request) (int, string) {
	// Certifica-se que Body estará fechado quando o método terminar
	defer req.Body.Close()

	// Leia o request body
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "Erro tentando ler req.Body"
	}

	// Continua se conseguiu ler req.Body corretamente
	if len(params) != 0 {
		// Não há chaves em params. Não pode ser?
		return http.StatusMethodNotAllowed, "método não permitido"
	}

	// Converte registro enviado pelo usuário
	var guestBookEntry GuestBookEntry
	err = json.Unmarshal(requestBody, &guestBookEntry)
	if err != nil {
		// não foi possível converter o registro
		return http.StatusBadRequest, "dados inválidos no JSON"
	}

	// Registro convertido com sucesso. Acrescenta o registro enviado pelo usuário
	g.AddEntry(guestBookEntry.Email, guestBookEntry.Title, guestBookEntry.Content)

	// Tudo certo!
	mensagem := "novo registro criado.... "
	fmt.Println(mensagem, guestBookEntry)
	return http.StatusOK, mensagem
}