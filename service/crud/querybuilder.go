package crud

import (
	"fmt"
	"strings"
)

func InsertQueryBuilder(table string, columns ...string) string {
	joinedColumns := strings.Join(columns, ",")
	placeholders := ""
	for i := range columns {
		if i != len(columns)-1 {
			placeholders += "?" + ","
		} else {
			placeholders += "?"
		}
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, joinedColumns, placeholders)
	return query
}
func SelectQueryBuilder(table string, columns, arrOfCondition []string, additionnalStr ...string) string {
	var (
		conditionStatement string
		joinedColumns      string
	)
	if len(arrOfCondition) != 0 {
		for i := range arrOfCondition {
			arrOfCondition[i] += "= ?"
		}
		joinedCondition := strings.Join(arrOfCondition, " AND ")
		conditionStatement = fmt.Sprintf("WHERE %s", joinedCondition)
	}
	joinedColumns = strings.Join(columns, ",")
	if len(columns) == 0 {
		joinedColumns = "*"
	}
	query := fmt.Sprintf("SELECT %s FROM %s %s", joinedColumns, table, conditionStatement)
	return query
}
