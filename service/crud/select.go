package crud

import (
	"database/sql"
)

func Select(db *sql.DB, table string, query string, conditions map[string]interface{}, columns []string, valuesToSelect ...interface{}) error {
	conditionPlaceholders := []string{}
	conditionValues := []interface{}{}

	if conditions != nil {
		// Collect placeholders and values for conditions
		for k, v := range conditions {
			conditionPlaceholders = append(conditionPlaceholders, k)
			conditionValues = append(conditionValues, v)
		}
	}
	if query == "" {
		// Construct the query using your SelectQueryBuilder
		query = SelectQueryBuilder(table, columns, conditionPlaceholders)
	}

	// Execute the query and scan the result into valuesToSelect
	err := db.QueryRow(query, conditionValues...).Scan(valuesToSelect...)
	if err != nil {
		return err
	}
	return nil
}
func SelectMultiple(db *sql.DB, table string, query string, conditions map[string]interface{}, columns []string) (*sql.Rows, error) {
	conditionPlaceholders := []string{}
	conditionValues := []interface{}{}

	if conditions != nil {
		// Collect placeholders and values for conditions
		for k, v := range conditions {
			conditionPlaceholders = append(conditionPlaceholders, k)
			conditionValues = append(conditionValues, v)
		}
	}
	if query == "" {
		// Construct the query using your SelectQueryBuilder
		query = SelectQueryBuilder(table, columns, conditionPlaceholders)
	}
	// Execute the query and return the result set
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil

}
