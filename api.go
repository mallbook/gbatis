package gbatis

import (
	"database/sql"
)

// SelectRow delegate SQLSession.SelectRow
func SelectRow(sqlID string, args ...interface{}) (*sql.Row, error) {
	s, err := OpenSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	return s.SelectRow(sqlID, args)
}

// Select delegate SQLSession.Select
func Select(sqlID string, args ...interface{}) ([]interface{}, error) {
	s, err := OpenSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	return s.Select(sqlID, args...)
}

// Execute delegate SQLSession.Execute
func Execute(sqlID string, args ...interface{}) (sql.Result, error) {
	s, err := OpenSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	return s.Execute(sqlID, args...)
}

// BulkInsert delegate SQLSession.BulkInsert
func BulkInsert(sqlID string, rows []interface{}) (sql.Result, error) {
	s, err := OpenSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()

	return s.BulkInsert(sqlID, rows)
}
