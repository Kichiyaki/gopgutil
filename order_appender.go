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
	tableAlias := ""
	if tableModel != nil {
		tableAlias = strings.ReplaceAll(string(tableModel.Table().Alias), "\"", "")
	}

outerLoop:
	for _, order := range o.Orders {
		parts := strings.Split(order, ".")
		depth := len(parts)
		lastIndex := depth - 1
		if o.MaxDepth > 0 && depth > o.MaxDepth {
			return q, errors.Errorf("order has depth %d, which exceeds the limit of %d", depth, o.MaxDepth)
		}

		var relation []string
		hasTableAlias := false
		columnAndKeyword := ""
		for i, part := range parts {
			if i == lastIndex {
				ind := strings.Index(part, " ")
				if ind >= 0 {
					columnAndKeyword = strutil.Underscore(part[:ind]) + " " + part[ind+1:]
				} else {
					columnAndKeyword = strutil.Underscore(part)
				}
				continue
			}
			if i == 0 && strutil.Underscore(part) == tableAlias {
				hasTableAlias = true
				continue
			}
			if tableModel == nil {
				continue outerLoop
			}
			relation = append(relation, strutil.PascalCase(part, '_'))
		}

		if len(relation) > 0 && tableModel != nil {
			relationStr := strings.Join(relation, ".")
			alias, err := BuildAliasFromRelationName(tableModel, relationStr)
			if err != nil {
				return q, err
			}
			if join := tableModel.GetJoin(relationStr); join == nil {
				q = q.Relation(relationStr + "._")
			}
			columnAndKeyword = alias + "." + columnAndKeyword
		} else if hasTableAlias {
			columnAndKeyword = tableAlias + "." + columnAndKeyword
		}

		orders = append(orders, columnAndKeyword)
	}

	return q.Order(orders...), nil
}
