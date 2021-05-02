package gopgutil

import "strings"

func AddAliasToColumnName(column, alias string) string {
	if alias != "" {
		return wrapStringInDoubleQuotes(alias) + "." + wrapStringInDoubleQuotes(column)
	}
	return wrapStringInDoubleQuotes(column)
}

func wrapStringInDoubleQuotes(str string) string {
	return `"` + str + `"`
}

func ExtractAlias(columnName string) string {
	if strings.Contains(columnName, ".") {
		return strings.ReplaceAll(strings.Split(columnName, ".")[0], `"`, "")
	}
	return ""
}
