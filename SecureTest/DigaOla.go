package main

import (
	"net/http"
	"strings"
	"encoding/base64"
)

/*
func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":9090", nil)

}
*/

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 { return false }

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil { return false }

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 { return false }

	return pair[0] == "user" && pair[1] == "pass"
}

func autoriza(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) {
		// w.Write([]byte("HAI!"))
		// myOriginalHandler.ServeHTTP(w, r)
		http.Handle("/static", http.FileServer(http.Dir("static")))
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="localhost"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

func main() {
	http.HandleFunc("/", autoriza)

	http.ListenAndServe(":2424", nil)
}
