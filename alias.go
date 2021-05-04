package gopgutil

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/types"
)

func AddAliasToColumnName(column, alias string) types.Ident {
	if alias != "" {
		return pg.Ident(alias + "." + column)
	}
	return pg.Ident(alias)
}
