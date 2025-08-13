package repository

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockBaseRepository[T any] struct {
	mock.Mock
}

func NewMockBaseRepository[T any]() *MockBaseRepository[T] {
	return &MockBaseRepository[T]{}
}

func (m *MockBaseRepository[T]) GetAll(items *[]T, pagination *PaginationQuery, condition *Condition, relations *[]Relation) error {
	args := m.Called(items, pagination, condition, relations)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) GetBy(item *T, condition *Condition, relations *[]Relation) error {
	args := m.Called(item, condition, relations)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) Create(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) CreateMany(items *[]T) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) Update(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) UpdateMany(items *[]T) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) Save(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) SaveMany(items *[]T) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) Delete(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) DeleteMany(items *[]T) error {
	args := m.Called(items)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) WithTx(tx *gorm.DB) BaseRepository[T] {
	args := m.Called(tx)
	return args.Get(0).(BaseRepository[T])
}

func (m *MockBaseRepository[T]) GetDB() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockBaseRepository[T]) Begin() BaseRepository[T] {
	args := m.Called()
	return args.Get(0).(BaseRepository[T])
}

func (m *MockBaseRepository[T]) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockBaseRepository[T]) Rollback() error {
	args := m.Called()
	return args.Error(0)
}
