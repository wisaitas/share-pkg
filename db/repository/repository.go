package repositoryutil

import (
	"gorm.io/gorm"
)

type (
	Condition struct {
		Query interface{}
		Args  []interface{}
	}

	Relation struct {
		Query string
		Args  []interface{}
	}

	PaginationQuery struct {
		Page     *int    `query:"page"`
		PageSize *int    `query:"page_size"`
		Sort     *string `query:"sort"`
		Order    *string `query:"order"`
	}

	repositoryUtil[T any] struct {
		db *gorm.DB
	}
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

func NewCondition(query interface{}, args ...interface{}) *Condition {
	return &Condition{
		Query: query,
		Args:  args,
	}
}

func NewRelation(query string, args ...interface{}) *Relation {
	return &Relation{
		Query: query,
		Args:  args,
	}
}

func NewPaginationQuery(page *int, pageSize *int, sort *string, order *string) *PaginationQuery {
	return &PaginationQuery{
		Page:     page,
		PageSize: pageSize,
		Sort:     sort,
		Order:    order,
	}
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

	return query.Find(items).Error
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

	return query.First(item).Error
}

func (r *repositoryUtil[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *repositoryUtil[T]) CreateMany(items *[]T) error {
	return r.db.Create(items).Error
}

func (r *repositoryUtil[T]) Update(item *T) error {
	return r.db.Updates(item).Error
}

func (r *repositoryUtil[T]) UpdateMany(items *[]T) error {
	return r.db.Updates(items).Error
}

func (r *repositoryUtil[T]) Save(item *T) error {
	return r.db.Save(item).Error
}

func (r *repositoryUtil[T]) SaveMany(items *[]T) error {
	return r.db.Save(items).Error
}

func (r *repositoryUtil[T]) Delete(item *T) error {
	return r.db.Delete(item).Error
}

func (r *repositoryUtil[T]) DeleteMany(items *[]T) error {
	return r.db.Delete(items).Error
}

func (r *repositoryUtil[T]) WithTx(tx *gorm.DB) RepositoryUtil[T] {
	return &repositoryUtil[T]{
		db: tx,
	}
}
