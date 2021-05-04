package gopgutil

import (
	"github.com/go-pg/pg/v10"
)

func AddAliasToColumnName(column, alias string) pg.Ident {
	if alias != "" {
		return pg.Ident(alias + "." + column)
	}
	return pg.Ident(alias)
}
