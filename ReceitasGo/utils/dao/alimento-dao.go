package dao

import (
	"fmt"
	"Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"github.com/jinzhu/gorm"
	"log"
)

type alimentoDAO struct {
	db *gorm.DB
}

// O singleton para a struct
var sharedAlimentoDAO *alimentoDAO

// O construtor singleton
func SharedAlimentoDAO() *alimentoDAO {

	once.Do(func() {
		sharedAlimentoDAO = newAlimentoDAO()
	})
	return sharedAlimentoDAO
}

func newAlimentoDAO() *alimentoDAO {
	db, err := SharedDbSession().Db()

	if err != nil {
		log.Fatalf("Não pude iniciar o db: %s", err)
		return nil
	}
	a := alimentoDAO{db: db}
	return &a
}

func (ac *alimentoDAO)Add(al *domain.Alimento) error {

	tx := ac.db.Begin()

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
