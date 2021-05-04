package gopgutil

import (
	"github.com/Kichiyaki/goutil/strutil"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"
	"strings"
)

type OrderAppender struct {
	Orders   []string
	MaxDepth int
}

func (o OrderAppender) Apply(q *orm.Query) (*orm.Query, error) {
	var orders []string
	tableModel := q.TableModel()
	if tableModel != nil {
		table := tableModel.Table()
		tableAlias := strings.ReplaceAll(string(table.Alias), "\"", "")
		for _, order := range o.Orders {
			parts := strings.Split(order, ".")
			depth := len(parts)
			if o.MaxDepth > 0 && depth > o.MaxDepth {
				return q, errors.Errorf("order has depth %d, which exceeds the limit of %d", depth, o.MaxDepth)
			}

			var relation []string
			hasTableAlias := false
			columnAndKeyword := ""
			for i, part := range parts {
				if i == 0 && strutil.Underscore(part) == tableAlias {
					hasTableAlias = true
					continue
				}
				if i == depth-1 {
					ind := strings.Index(part, " ")
					if ind >= 0 {
						columnAndKeyword = strutil.Underscore(part[:ind]) + " " + part[ind+1:]
					} else {
						columnAndKeyword = strutil.Underscore(part)
					}
					continue
				}
				relation = append(relation, strutil.PascalCase(part, '_'))
			}

			if len(relation) > 0 {
				relationStr := strings.Join(relation, ".")
				alias, err := BuildAliasFromRelationName(tableModel, relationStr)
				if err != nil {
					return q, err
				}
				q = q.Relation(relationStr)
				columnAndKeyword = alias + "." + columnAndKeyword
			} else if hasTableAlias {
				columnAndKeyword = tableAlias + "." + columnAndKeyword
			}

			orders = append(orders, columnAndKeyword)
		}
	}

	return q.Order(orders...), nil
}
