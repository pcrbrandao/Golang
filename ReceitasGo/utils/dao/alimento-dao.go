package dao

import (
	"fmt"
	d "Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"sync"
)

type alimentoDAO struct {
	// não é necessário guardar isso
	// db *gorm.DB
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
	// db, err := SharedDbSession().OpenDb()

	// if err != nil {
	//	log.Fatalf("Não pude iniciar o db: %s", err)
	//	return nil
	//}
	//a := alimentoDAO{db: db}
	a := alimentoDAO{}
	return &a
}

func (ac *alimentoDAO)Add(al *d.Alimento) error {

	db, _ := SharedDbSession().OpenDb()
	defer db.Close()

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

func (ac *alimentoDAO)Find(descricao string) d.Alimento {
	db, _ := SharedDbSession().OpenDb()
	defer db.Close()

	var alimento d.Alimento
	db.Where(&d.Alimento{Descricao:descricao}).First(&alimento)

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
