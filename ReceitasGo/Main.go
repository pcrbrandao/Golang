package main

import (
	"Golang/ReceitasGo/utils"
	"fmt"
)

func main() {

	mySQL := utils.SharedMySqlPass()
	db, err := mySQL.Db()

	if err != nil {
		fmt.Println(utils.ERROCONECTANDO.String(), err.Error())
	}

	defer db.Close()
}
