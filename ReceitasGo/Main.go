package ReceitasGo

import (
	"Golang/ReceitasGo/utils"
	"fmt"
	"Golang/ReceitasGo/mensagem"
)

func main() {

	mySQL := utils.SharedControl()
	db, err := mySQL.Db()

	if err != nil {
		fmt.Println(mensagem.ERROCONECTANDO, err.Error())
	}

	defer db.Close()
}
