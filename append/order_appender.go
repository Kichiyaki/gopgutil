package append

import (
	"strings"

	"github.com/go-pg/pg/v10/orm"
)

type Relation struct {
	Name  string
	Apply []func(*orm.Query) (*orm.Query, error)
}

type Join struct {
	Query  string
	Params []interface{}
}

type Sort struct {
	Relations map[string]Relation
	Joins     map[string]Join
	Orders    []string
}

func (s Sort) Apply(q *orm.Query) (*orm.Query, error) {
	for _, order := range s.Orders {
		alias := s.extractAlias(order)
		if alias != "" {
			if relation, ok := s.Relations[alias]; ok {
				q = q.Relation(relation.Name, relation.Apply...)
			}
			if join, ok := s.Joins[alias]; ok {
				q = q.Join(join.Query, join.Params...)
			}
		}
	}
	return q.Order(s.Orders...), nil
}

func (s Sort) extractAlias(order string) string {
	if strings.Contains(order, ".") {
		return strings.Split(order, ".")[0]
	}
	return ""
}
