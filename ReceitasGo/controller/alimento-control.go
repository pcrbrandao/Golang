package controller

import (
	"fmt"
	"Golang/ReceitasGo/model"
	"Golang/ReceitasGo/utils"
	"Golang/ReceitasGo/mensagem"
)

type AlimentoControl struct { }

func (ac AlimentoControl)Adiciona(al *model.Alimento) error {

	db, err := utils.SharedControl().Db()

	if err != nil {
		return err
	}

	tx := db.Begin()

	if !tx.NewRecord(al) {
		return fmt.Errorf(mensagem.JAEXISTE)
	}

	if err := tx.Create(&al).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ac AlimentoControl)Edita(al *model.Alimento) error {

	return fmt.Errorf(mensagem.EMDESENVOLVIMENTO)
}

func (ac AlimentoControl)Exclui(al *model.Alimento) error {

	return fmt.Errorf(mensagem.EMDESENVOLVIMENTO)
}

func (ac AlimentoControl)Lista(al *model.Alimento) error {

	return fmt.Errorf(mensagem.EMDESENVOLVIMENTO)
}