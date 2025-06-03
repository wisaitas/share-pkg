package repositoryutil

import (
	"fmt"

	"gorm.io/gorm"
)

type RepositoryUtil[T any] interface {
	GetAll(items *[]T, pagination *PaginationQuery, condition *Condition, relations *[]Relation) error
	GetBy(item *T, condition *Condition, relations *[]Relation) error
	Create(item *T) error
	CreateMany(items *[]T) error
	Update(item *T) error
	UpdateMany(items *[]T) error
	Save(item *T) error
	SaveMany(items *[]T) error
	Delete(item *T) error
	DeleteMany(items *[]T) error
	WithTx(tx *gorm.DB) RepositoryUtil[T]
}

type repositoryUtil[T any] struct {
	db *gorm.DB
}

func New[T any](db *gorm.DB) RepositoryUtil[T] {
	return &repositoryUtil[T]{
		db: db,
	}
}

func (r *repositoryUtil[T]) GetAll(items *[]T, pagination *PaginationQuery, condition *Condition, relations *[]Relation) error {
	query := r.db

	if condition != nil {
		query = query.Where(condition.Query, condition.Args...)
	}

	if pagination != nil {
		if pagination.Page != nil && pagination.PageSize != nil {
			offset := *pagination.Page * *pagination.PageSize
			query = query.Offset(offset).Limit(*pagination.PageSize)
		}

		if pagination.Sort != nil && pagination.Order != nil {
			orderClause := *pagination.Sort + " " + *pagination.Order
			query = query.Order(orderClause)
		}
	}

	if relations != nil {
		for _, relation := range *relations {
			query = query.Preload(relation.Query, relation.Args...)
		}
	}

	if err := query.Find(items).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) GetBy(item *T, condition *Condition, relations *[]Relation) error {
	query := r.db

	if condition != nil {
		query = query.Where(condition.Query, condition.Args...)
	}

	if relations != nil {
		for _, relation := range *relations {
			query = query.Preload(relation.Query, relation.Args...)
		}
	}

	if err := query.First(item).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) Create(item *T) error {
	if err := r.db.Create(item).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) CreateMany(items *[]T) error {
	if err := r.db.Create(items).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) Update(item *T) error {
	if err := r.db.Updates(item).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) UpdateMany(items *[]T) error {
	if err := r.db.Updates(items).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) Save(item *T) error {
	if err := r.db.Save(item).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) SaveMany(items *[]T) error {
	if err := r.db.Save(items).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) Delete(item *T) error {
	if err := r.db.Delete(item).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) DeleteMany(items *[]T) error {
	if err := r.db.Delete(items).Error; err != nil {
		return fmt.Errorf("[Share Package RepositoryUtil] : %w", err)
	}

	return nil
}

func (r *repositoryUtil[T]) WithTx(tx *gorm.DB) RepositoryUtil[T] {
	return &repositoryUtil[T]{
		db: tx,
	}
}
