package dao

import (
	"fmt"
	d "Golang/ReceitasGo/domain"
	"Golang/ReceitasGo/utils/misc"
	"sync"
)

// A struct é privada. Uma instância deve ser obtida por SharedAlimentoDAO
// para evitar que haja mais de uma instância no aplicativo.
// Todo método que utiliza o db deve abrí-lo e fechá-lo em seguida.
type alimentoDAO struct {
}

// O singleton para a struct
// As variáveis ficam disponíveis para o package dao
// e devem ser únicas para cada singleton
var sharedAlimentoDAO *alimentoDAO
var onceAlimentoDAO sync.Once

// O construtor singleton
func SharedAlimentoDAO() *alimentoDAO {

	onceAlimentoDAO.Do(func() {
		sharedAlimentoDAO = &alimentoDAO{}
	})
	return sharedAlimentoDAO
}

// Adiciona um alimento ao db com uma transação.
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

// Encontra um alimento pela descrição e o retorna.
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
