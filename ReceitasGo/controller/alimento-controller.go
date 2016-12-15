package controller

import (
	"fmt"
	"Golang/ReceitasGo/model"
	"Golang/ReceitasGo/utils"
)

func Adiciona(alimento *model.Alimento) error {

	db, err := utils.SharedMySqlPass().Db()

	if err != nil {
		return err
	}

	tx := db.Begin()

	if !tx.NewRecord(alimento) {
		return fmt.Errorf("o registro já existe")
	}

	if err := tx.Create(&alimento).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

