package gopgutil

func BuildConditionEquals(column string) string {
	return column + " = ?"
}

func BuildConditionNEQ(column string) string {
	return column + " != ?"
}

func BuildConditionLT(column string) string {
	return column + " < ?"
}

func BuildConditionLTE(column string) string {
	return column + " <= ?"
}

func BuildConditionGT(column string) string {
	return column + " > ?"
}

func BuildConditionGTE(column string) string {
	return column + " >= ?"
}

func BuildConditionMatch(column string) string {
	return column + " LIKE ?"
}

func BuildConditionIEQ(column string) string {
	return column + " ILIKE ?"
}

func BuildConditionIn(column string) string {
	return column + " IN (?)"
}

func BuildConditionArray(column string) string {
	return column + " = ANY(?)"
}

func BuildConditionNotInArray(column string) string {
	return "NOT (" + BuildConditionArray(column) + ")"
}
