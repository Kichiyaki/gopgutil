package gopgutil

import (
	"github.com/Kichiyaki/goutil/strutil"
	"regexp"
	"strings"
)

var (
	orderRegex = regexp.MustCompile(`^[\p{L}_.]+$`)
)

func SanitizeOrder(order string) string {
	parts := strings.Split(strings.TrimSpace(order), " ")
	length := len(parts)

	if length != 2 || !orderRegex.Match([]byte(parts[0])) {
		return ""
	}

	table := ""
	column := parts[0]
	if strings.Contains(parts[0], ".") {
		columnAndTable := strings.Split(parts[0], ".")
		table = strutil.Underscore(columnAndTable[0]) + "."
		column = columnAndTable[1]
	}

	direction := "ASC"
	if strings.ToUpper(parts[1]) == "DESC" {
		direction = "DESC"
	}

	return strings.ToLower(table+strutil.Underscore(column)) + " " + direction
}

func SanitizeOrders(orders []string) []string {
	var sanitizedOrders []string
	for _, sort := range orders {
		sanitized := SanitizeOrder(sort)
		if sanitized != "" {
			sanitizedOrders = append(sanitizedOrders, sanitized)
		}
	}
	return sanitizedOrders
}
