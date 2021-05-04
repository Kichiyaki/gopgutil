package gopgutil

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"
	"strings"
)

var ErrRelationNotFound = errors.New("relation not found")

func BuildAliasFromRelationName(model interface{}, relationName string) (string, error) {
	var tableModel orm.TableModel
	switch v := model.(type) {
	case orm.TableModel:
		tableModel = v
	default:
		m, err := orm.NewModel(model)
		if err != nil {
			return "", errors.Wrap(err, "Invalid model")
		}
		tableModel = m.(orm.TableModel)
	}

	alias := buildAliasFromRelationName(tableModel.Table().Relations, "", strings.Split(relationName, ".")...)
	if alias == "" {
		return "", ErrRelationNotFound
	}
	return alias, nil
}

func buildAliasFromRelationName(relations map[string]*orm.Relation, alias string, parts ...string) string {
	if len(parts) == 0 {
		return alias
	}
	current := parts[0]
	var next map[string]*orm.Relation
	for name, relation := range relations {
		if name == current {
			next = relation.JoinTable.Relations
			if alias != "" {
				alias += "__"
			}
			alias += relation.Field.SQLName
		}
	}
	if next == nil {
		return ""
	}
	return buildAliasFromRelationName(next, alias, parts[1:]...)
}
