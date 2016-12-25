package dao

import (
	"fmt"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"sync"
)

type alimentoDAO struct {
}

// O singleton para a struct
var sharedAlimentoDAO *alimentoDAO
var onceAlimentoDAO sync.Once

// O construtor singleton
func SharedAlimentoDAO() *alimentoDAO {

	onceAlimentoDAO.Do(func() {
		sharedAlimentoDAO = &alimentoDAO{}
	})
	return sharedAlimentoDAO
}

func (ac alimentoDAO)Add(al *domain.Alimento) error {

	db, err := SharedDbSession().Db()

	if err != nil {
		return err
	}

	tx := db.Begin()

	if !tx.NewRecord(al) {
		return fmt.Errorf("%s", misc.JAEXISTE)
	}

	if err := tx.Create(&al).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ac alimentoDAO)Find(al *domain.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac alimentoDAO)List(al *domain.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac alimentoDAO)Edit(al *domain.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac alimentoDAO)Delete(al *domain.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}
