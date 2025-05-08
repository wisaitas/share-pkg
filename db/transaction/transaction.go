package transaction

import (
	"gorm.io/gorm"
)

type TransactionUtil interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	GetTransaction() *gorm.DB
	Begin() error
	Commit() error
	Rollback() error
}

type util struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) TransactionUtil {
	return &util{
		db: db,
	}
}

func (tm *util) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
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

func (tm *util) GetTransaction() *gorm.DB {
	if tm.tx != nil {
		return tm.tx
	}
	return tm.db
}

func (tm *util) Begin() error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	tm.tx = tx
	return nil
}

func (tm *util) Commit() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Commit().Error
	tm.tx = nil
	return err
}

func (tm *util) Rollback() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Rollback().Error
	tm.tx = nil
	return err
}
