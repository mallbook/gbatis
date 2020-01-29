package gbatis

import (
	"database/sql"
	"fmt"

	"github.com/mallbook/gbatis/bean"
)

type sqlSession struct {
	db *sql.DB
}

// SelectRow return one row from database
func (s sqlSession) selectRow(sql string, args ...interface{}) (*sql.Row, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sql)
	}

	sql, err := preparedDyncSQL(sql)
	if err != nil {
		return nil, err
	}

	if containNamedArgs(sql) && len(args) > 0 {
		var err error
		sql, args, err = preparedNamedArgs(sql, args[0])
		if err != nil {
			return nil, err
		}
	}

	return s.db.QueryRow(sql, args...), nil
}

// SelectOne return one object
func (s sqlSession) selectOne(sql, resultType string, args ...interface{}) (interface{}, error) {
	r, err := s.selectList(sql, resultType, args...)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}
	return nil, err
}

// Select return multi objects from database
func (s sqlSession) selectList(sql, resultType string, args ...interface{}) ([]interface{}, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sql)
	}

	result, err := bean.New(resultType)
	if err != nil {
		return nil, fmt.Errorf("NewBean fail, err=%s, resultType=%s, sqlID=%s", err, resultType, sql)
	}

	sql, err = preparedDyncSQL(sql)
	if err != nil {
		return nil, err
	}

	if containNamedArgs(sql) && len(args) > 0 {
		var err error
		sql, args, err = preparedNamedArgs(sql, args[0])
		if err != nil {
			return nil, err
		}
	}

	rows, err := s.db.Query(sql, args...)
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
	typ := result.Elem().Type()
	numField := typ.NumField()
	for i := 0; i < len; i++ {
		j := 0
		for ; j < numField; j++ {
			f := typ.Field(j)
			if f.Tag.Get("json") == columns[i] {
				break
			}
		}
		if j < numField {
			field := result.Elem().Field(j)
			dest[i] = field.Addr().Interface()
		}
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
func (s sqlSession) execute(sql string, args ...interface{}) (sql.Result, error) {

	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sql)
	}

	sql, err := preparedDyncSQL(sql)
	if err != nil {
		return nil, err
	}

	if containNamedArgs(sql) && len(args) > 0 {
		var err error
		sql, args, err = preparedNamedArgs(sql, args[0])
		if err != nil {
			return nil, err
		}
	}

	return s.db.Exec(sql, args...)
}

// BulkInsert means bulk insert data
// sql syntax: insert into t(a1, a2, a3) values (?, ?, ?),(?, ?, ?)...
func (s sqlSession) bulkInsert(sql string, rows []interface{}) (sql.Result, error) {
	if s.db == nil {
		return nil, fmt.Errorf("The session was closed, sqlID=%s", sql)
	}

	if rows == nil || len(rows) == 0 {
		return nil, fmt.Errorf("%s", "The rows is nil or len is zero.")
	}

	if containNamedArgs(sql) {
		var err error
		sql, rows, err = preparedBulkNamedArgs(sql, rows)
		if err != nil {
			return nil, err
		}
	}

	ss, err := cutInsertSQL(sql)
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
func (s sqlSession) close() {
	if s.db != nil {
		s.db = nil
	}
}

// OpenSession Open a sql session
func openSession(dbIDs ...string) (s *sqlSession, err error) {
	dbmgr := getDBMgrInstance()
	if dbIDs == nil {
		defaultID := dbmgr.defaultID
		if db, ok := dbmgr.getDB(defaultID); ok {
			s = &sqlSession{
				db: db,
			}
		} else {
			err = fmt.Errorf("Not found default DB, defaultID=%s", defaultID)
		}
		return
	}

	for _, dbID := range dbIDs {
		if db, ok := dbmgr.getDB(dbID); ok {
			s = &sqlSession{
				db: db,
			}
			return
		}
	}

	err = fmt.Errorf("Not found any DB instance, dbIDs=%p", dbIDs)

	return
}
