package repository

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
