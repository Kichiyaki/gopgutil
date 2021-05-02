package gopgutil

import "fmt"

func BuildCountColumnExpr(column, alias string) string {
	base := fmt.Sprintf("count(%s)", column)
	if alias != "" {
		return base + " as " + alias
	}
	return base
}
