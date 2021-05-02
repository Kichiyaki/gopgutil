package gopgutil

import (
	"github.com/go-pg/pg/v10/orm"
)

type OrderAppenderRelation struct {
	Name  string
	Apply []func(*orm.Query) (*orm.Query, error)
}

type OrderAppenderJoin struct {
	Query  string
	Params []interface{}
}

type OrderAppender struct {
	Relations map[string]OrderAppenderRelation
	Joins     map[string]OrderAppenderJoin
	Orders    []string
}

func (o OrderAppender) Apply(q *orm.Query) (*orm.Query, error) {
	for _, order := range o.Orders {
		alias := ExtractAlias(order)
		if alias != "" {
			if relation, ok := o.Relations[alias]; ok {
				q = q.Relation(relation.Name, relation.Apply...)
			}
			if join, ok := o.Joins[alias]; ok {
				q = q.Join(join.Query, join.Params...)
			}
		}
	}
	return q.Order(o.Orders...), nil
}
