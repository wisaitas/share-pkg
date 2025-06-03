package transactionutil

import (
	"fmt"

	"gorm.io/gorm"
)

type TransactionUtil interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	GetTransaction() *gorm.DB
	Begin() error
	Commit() error
	Rollback() error
}

type transactionUtil struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) TransactionUtil {
	return &transactionUtil{
		db: db,
	}
}

func (tm *transactionUtil) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	return nil
}

func (tm *transactionUtil) GetTransaction() *gorm.DB {
	if tm.tx != nil {
		return tm.tx
	}
	return tm.db
}

func (tm *transactionUtil) Begin() error {
	tx := tm.db.Begin()

	if tx.Error != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", tx.Error)
	}

	tm.tx = tx
	return nil
}

func (tm *transactionUtil) Commit() error {
	if tm.tx == nil {
		return nil
	}

	if err := tm.tx.Commit().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	tm.tx = nil
	return nil
}

func (tm *transactionUtil) Rollback() error {
	if tm.tx == nil {
		return nil
	}

	if err := tm.tx.Rollback().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	tm.tx = nil
	return nil
}
