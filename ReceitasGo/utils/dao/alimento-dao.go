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

	db, db_err := SharedDbSession().OpenDb()
	defer db.Close()

	if db_err != nil {
		return db_err
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

// Encontra um alimento pela descrição e o retorna se deleted_at é NULL.
func (ac *alimentoDAO)Find(descricao string) (*d.Alimento, error) {
	db, db_err := SharedDbSession().OpenDb()
	defer db.Close()

	if db_err != nil {
		return nil, db_err
	}

	var alimento d.Alimento
	db.Where(&d.Alimento{Descricao:descricao}).First(&alimento)

	return &alimento, nil
}

// Apaga um alimento. Para evitar uma consulta duplicada o método
// não ferifica se o registro existe antes de deletar.
// O registro não é apagado no banco, mas sim marcado como deletado.
// Um novo método deve apagar os registros marcados definitivamente.
func (ac *alimentoDAO)Delete(al *d.Alimento) error {
	db, err := SharedDbSession().OpenDb()
	defer db.Close()

	if err != nil {
		return err
	}

	// Unscoped deleta o registro definitivamente.
	// err = db.Unscoped().Delete(al).Error
	err = db.Delete(al).Error
	if err != nil {
		return err
	}

	return nil
}

// Retorna os registros onde deleted_at não é nulo.
func (ac *alimentoDAO)FindDeleted() ([]d.Alimento, error) {
	db, err := SharedDbSession().OpenDb()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	var deletados []d.Alimento

	if err := db.Unscoped().Where("deleted_at <> ?", "NULL").Find(&deletados).Error; err != nil {
		return nil, err
	}
	return deletados, nil
}

func (ac *alimentoDAO)List(al *d.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}

func (ac *alimentoDAO)Edit(al *d.Alimento) error {

	return fmt.Errorf("%s", misc.EMDESENVOLVIMENTO)
}
