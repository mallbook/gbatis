package gbatis

import (
	"database/sql"
	"fmt"
)

// SQLSession hold a sql.DB object
type SQLSession struct {
	*sqlSession
}

// SelectRow return one row from database
func (s SQLSession) SelectRow(sqlID string, args ...interface{}) (*sql.Row, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sqlID)
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}
	if si.class != selectClass {
		return nil, fmt.Errorf("The class of this sql is not selectClass, but %d", si.class)
	}

	return s.selectRow(si.sql, args...)
}

// SelectOne return one object
func (s SQLSession) SelectOne(sqlID string, args ...interface{}) (interface{}, error) {
	r, err := s.Select(sqlID, args...)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}
	return nil, err
}

// Select return multi objects from database
func (s SQLSession) Select(sqlID string, args ...interface{}) ([]interface{}, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sqlID)
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}

	if si.class != selectClass {
		return nil, fmt.Errorf("Sql class is not selectClass, but %d, sqlID=%s", si.class, sqlID)
	}

	if si.resultType == "" {
		return nil, fmt.Errorf("The result type is empty, sqlID=%s", sqlID)
	}

	return s.selectList(si.sql, si.resultType, args...)
}

// Execute write database
func (s SQLSession) Execute(sqlID string, args ...interface{}) (sql.Result, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sqlID)
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}

	if si.class != insertClass && si.class != updateClass && si.class != deleteClass && si.class != anyClass {
		return nil, fmt.Errorf("Sql class is not insertClass, updateClass, deleteClass, but %d, sqlID=%s", si.class, sqlID)
	}

	return s.execute(si.sql, args...)
}

// BulkInsert means bulk insert data
// sql syntax: insert into t(a1, a2, a3) values (?, ?, ?),(?, ?, ?)...
func (s SQLSession) BulkInsert(sqlID string, rows []interface{}) (sql.Result, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sqlID)
	}

	if rows == nil || len(rows) == 0 {
		return nil, fmt.Errorf("%s", "The rows is nil or len is zero.")
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}

	if si.class != insertClass {
		return nil, fmt.Errorf("Sql class is not insertClass, but %d, sqlID=%s", si.class, sqlID)
	}

	return s.bulkInsert(si.sql, rows)
}

// Close free a database connection
func (s SQLSession) Close() {
	s.close()
}

// OpenSession Open a sql session
func OpenSession(dbIDs ...string) (s *SQLSession, err error) {
	ss, err := openSession(dbIDs...)
	if err != nil {
		return nil, err
	}

	return &SQLSession{
		sqlSession: ss,
	}, nil
}