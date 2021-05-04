package gopgutil

func AddAliasToColumnName(column, alias string) string {
	if alias != "" {
		return alias + "." + column
	}
	return column
}
