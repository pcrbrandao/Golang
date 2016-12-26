package dao

import (
	"fmt"
	d "Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type alimentoDAO struct {
	db *gorm.DB
}

// O singleton para a struct
// As variáveis ficam disponíveis para o package dao
// e devem ser únicas para cada singleton
var sharedAlimentoDAO *alimentoDAO
var onceAlimentoDAO sync.Once

// O construtor singleton
func SharedAlimentoDAO() *alimentoDAO {

	onceAlimentoDAO.Do(func() {
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

func (ac *alimentoDAO)Add(al *d.Alimento) error {

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

func (ac *alimentoDAO)Find(descricao string) d.Alimento {

	var alimento d.Alimento
	ac.db.Where(&d.Alimento{Descricao:descricao}).First(&alimento)

	return alimento
}

func (ac *alimentoDAO)List(al *d.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac *alimentoDAO)Edit(al *d.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac *alimentoDAO)Delete(al *d.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}
