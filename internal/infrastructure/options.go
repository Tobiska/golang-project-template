package infrastructure

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type Option interface {
	Apply(query sq.SelectBuilder) sq.SelectBuilder //todo change to sq
}

const (
	Asc  = "Asc"
	Desc = "Desc"
)

type SortOption struct {
	Field, Direction string
}

func (so *SortOption) Apply(query sq.SelectBuilder) sq.SelectBuilder {
	if so.Field == "" && so.Direction == "" {
		return query
	}
	query = query.OrderBy(fmt.Sprintf("%s %s", so.Field, so.Direction))
	return query
}

const (
	LT  = "lt"
	LTE = "lte"
	GT  = "gt"
	GTE = "gte"
	E   = "e"
)

type FilterOption struct {
	Field, Value, Operator string
}

func (fo *FilterOption) Apply(query sq.SelectBuilder) sq.SelectBuilder {
	switch fo.Operator {
	case GTE:
		query = query.Where(sq.GtOrEq{fo.Field: fo.Value})
	case GT:
		query = query.Where(sq.Gt{fo.Field: fo.Value})
	case LTE:
		query = query.Where(sq.LtOrEq{fo.Field: fo.Value})
	case LT:
		query = query.Where(sq.Lt{fo.Field: fo.Value})
	case E:
		query = query.Where(sq.Eq{fo.Field: fo.Value})
	}
	return query
}
