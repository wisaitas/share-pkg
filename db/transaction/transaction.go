package transaction

import (
	"gorm.io/gorm"
)

type Transaction interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	GetTransaction() *gorm.DB
	Begin() error
	Commit() error
	Rollback() error
}

type transaction struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) Transaction {
	return &transaction{
		db: db,
	}
}

func (tm *transaction) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tm *transaction) GetTransaction() *gorm.DB {
	if tm.tx != nil {
		return tm.tx
	}
	return tm.db
}

func (tm *transaction) Begin() error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	tm.tx = tx
	return nil
}

func (tm *transaction) Commit() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Commit().Error
	tm.tx = nil
	return err
}

func (tm *transaction) Rollback() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Rollback().Error
	tm.tx = nil
	return err
}
