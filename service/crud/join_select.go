package crud

import "database/sql"

func JoinSelect(db *sql.DB, query string, args ...interface{}) ([]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		var result interface{}
		err := rows.Scan(result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
