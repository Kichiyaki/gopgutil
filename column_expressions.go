package gopgutil

func BuildCountColumnExpr(column, alias string) string {
	base := "count(" + column + ")"
	if alias != "" {
		return base + " as " + alias
	}
	return base
}
