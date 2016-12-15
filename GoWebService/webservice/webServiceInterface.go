package webservice

import (
	"net/http"
	"github.com/codegangsta/martini"
)

// deve ser implementada por tipo que querem prover web services
type WebService interface {
	// retorna o caminho a ser associado com o serviço
	GetPath()string

	// DELETE request
	// retorna a resposta do servidor, código e descrição ?
	WebDelete(params martini.Params) (int, string)

	// GET request
	WebGet(params martini.Params) (int, string)

	// POST request
	WebPost(params martini.Params, req *http.Request) (int, string)
}

// acrescenta rotas martini para métodos webservice relevante baseado no path
// retornado por GetPath
func RegisterWebService(webService WebService, classicMartini *martini.ClassicMartini) {

	path := webService.GetPath()

	classicMartini.Get(path, webService.WebGet)
	classicMartini.Get(path+"/:id", webService.WebGet)

	classicMartini.Post(path, webService.WebPost)
	classicMartini.Post(path+"/:id", webService.WebPost)

	classicMartini.Delete(path, webService.WebDelete)
	classicMartini.Delete(path+"/:id", webService.WebDelete)
}