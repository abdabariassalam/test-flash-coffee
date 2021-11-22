package agent

import (
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return repository{db}
}

type Repository interface {
	GetDBMaster() *gorm.DB
	TxAgent(dbAgent *gorm.DB, fn func(tx *gorm.DB) error) error
}

func (r repository) GetDBMaster() *gorm.DB {
	return r.db
}

func (r repository) TxAgent(dbAgent *gorm.DB, fn func(tx *gorm.DB) error) error {
	err := dbAgent.Transaction(fn)
	return err
}
