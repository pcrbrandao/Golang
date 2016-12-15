package main

import (
	"Golang/ReceitasGo/utils"
	"fmt"
)

func main() {

	mySQL := utils.SharedMySqlPass()
	db, err := mySQL.Db()

	if err != nil {
		fmt.Println("Não pude conectar ao banco...", err.Error())
	}

	defer db.Close()
}
