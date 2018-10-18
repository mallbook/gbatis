package gbatis

import (
	"database/sql"
	"fmt"
)

// SQLSession hold a sql.DB object
type SQLSession struct {
	db *sql.DB
}

// SelectRow query one record from database
func (s SQLSession) SelectRow(sqlID string, args ...interface{}) (*sql.Row, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The sql.DB object is nil")
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}

	if si.class != selectClass {
		return nil, fmt.Errorf("The class of this sql is not selectClass, but %d", si.class)
	}

	return s.db.QueryRow(si.sql, args...), nil
}

// Select means select multi records from database
func (s SQLSession) Select(sqlID string, args ...interface{}) ([]interface{}, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The sql.DB object is nil, sqlID=%s", sqlID)
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

	result, err := NewBean(si.resultType)
	if err != nil {
		return nil, fmt.Errorf("NewBean fail, err=%s, resultType=%s, sqlID=%s", err, si.resultType, sqlID)
	}

	rows, err := s.db.Query(si.sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	len := len(columns)
	dest := make([]interface{}, len)
	for i := 0; i < len; i++ {
		field := result.Elem().FieldByName(columns[i])
		dest[i] = field.Addr().Interface()
	}

	results := make([]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		results = append(results, result.Elem().Interface())
	}

	return results, nil
}

// Execute write database
func (s SQLSession) Execute(sqlID string, args ...interface{}) (sql.Result, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The sql.DB object is nil, sqlID=%s", sqlID)
	}

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, err
	}

	if si.class != insertClass && si.class != updateClass && si.class != deleteClass {
		return nil, fmt.Errorf("Sql class is not insertClass, updateClass, deleteClass, but %d, sqlID=%s", si.class, sqlID)
	}

	return s.db.Exec(si.sql, args...)
}

// BulkInsert means bulk insert data
// sql syntax: insert into t(a1, a2, a3) values (?, ?, ?),(?, ?, ?)...
func (s SQLSession) BulkInsert(sqlID string, rows []interface{}) (sql.Result, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The sql.DB object is nil, sqlID=%s", sqlID)
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

	ss, err := cutInsertSQL(si.sql)
	if err != nil {
		return nil, err
	}

	sqlStr := ss[0]
	args := make([]interface{}, 0)
	for _, row := range rows {
		sqlStr += ss[1] + ","
		vals, ok := row.([]interface{})
		if ok {
			args = append(args, vals...)
		}
	}

	// trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, err := s.db.Prepare(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("Prepare sql fail, err=%v, sql=%s", err, sqlStr)
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// Close free a database connection
func (s SQLSession) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

// OpenSession Open a sql session
func OpenSession(dbIDs ...string) (s *SQLSession, err error) {
	dbmgr := getDBMgrInstance()
	if dbIDs == nil {
		defaultID := dbmgr.defaultID
		if db, ok := dbmgr.getDB(defaultID); ok {
			s = new(SQLSession)
			s.db = db
		} else {
			err = fmt.Errorf("Not found default DB, defaultID=%s", defaultID)
		}
		return
	}

	for _, dbID := range dbIDs {
		if db, ok := dbmgr.getDB(dbID); ok {
			s = new(SQLSession)
			s.db = db
			return
		}
	}

	err = fmt.Errorf("Not found any DB instance, dbIDs=%p", dbIDs)

	return
}
