package gbatis

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"text/template"
)

const (
	leftDelim  = "${"
	rightDelim = "}"
)

var (
	re = regexp.MustCompile("\\${[0-9a-zA-Z_\\- \t\\.]+}")

	tempalteMgr  *sqlTemplateMgr
	templateOnce sync.Once
)

type sqlTemplateMgr struct {
	sync.RWMutex
	tmpls map[string]*template.Template
}

func newSQLTemplateMgr() *sqlTemplateMgr {
	templateOnce.Do(func() {
		tempalteMgr = &sqlTemplateMgr{
			tmpls: make(map[string]*template.Template),
		}
	})
	return tempalteMgr
}

func (m *sqlTemplateMgr) getTmpl(sqlID string) (*template.Template, *sqlInfor, error) {

	si, err := getSQLMgrInstance().getSQL(sqlID)
	if err != nil {
		return nil, nil, err
	}

	if !judgeTemplate(si.sql) {
		return nil, nil, fmt.Errorf("The sql(%s) is not a sql template", sqlID)
	}

	if t, ok := m.get(sqlID); ok {
		return t, &si, nil
	}

	funcs := template.FuncMap{
		"in":     in,
		"driver": driver,
	}

	tmpl, err := template.New(sqlID).Delims(leftDelim, rightDelim).Funcs(funcs).Parse(si.sql)
	if err != nil {
		return nil, nil, err
	}

	m.set(sqlID, tmpl)

	return tmpl, &si, err
}

func (m *sqlTemplateMgr) get(sqlID string) (*template.Template, bool) {
	m.RLock()
	defer m.RUnlock()
	t, ok := m.tmpls[sqlID]
	return t, ok
}

func (m *sqlTemplateMgr) set(sqlID string, tmpl *template.Template) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.tmpls[sqlID]; !ok {
		m.tmpls[sqlID] = tmpl
	}
}

// SQLTemplate represents a sql template
type SQLTemplate struct {
	dbID      string
	statement string
	err       error
	*sqlInfor
}

// Template return SQLTemplate object
func Template(sqlID string, data interface{}) *SQLTemplate {

	tmpl, si, err := newSQLTemplateMgr().getTmpl(sqlID)
	if err != nil {
		return &SQLTemplate{
			err: err,
		}
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return &SQLTemplate{
			err: err,
		}
	}

	return &SQLTemplate{
		dbID:      getDBMgrInstance().defaultID,
		statement: buf.String(),
		sqlInfor:  si,
	}
}

// SelectRow return one row from database
func (t SQLTemplate) SelectRow(args ...interface{}) (*sql.Row, error) {
	if t.err != nil {
		return nil, t.err
	}
	s, err := openSession(t.dbID)
	if err != nil {
		return nil, err
	}
	defer s.close()

	return s.selectRow(t.statement, args...)
}

// SelectOne delegate SQLSession.SelectOne
func (t SQLTemplate) SelectOne(args ...interface{}) (interface{}, error) {
	if t.err != nil {
		return nil, t.err
	}
	s, err := openSession(t.dbID)
	if err != nil {
		return nil, err
	}
	defer s.close()

	return s.selectOne(t.statement, t.resultType, args...)
}

// Select delegate SQLSession.Select
func (t SQLTemplate) Select(args ...interface{}) ([]interface{}, error) {
	if t.err != nil {
		return nil, t.err
	}
	s, err := openSession(t.dbID)
	if err != nil {
		return nil, err
	}
	defer s.close()

	return s.selectList(t.statement, t.resultType, args...)
}

// Execute delegate SQLSession.Execute
func (t SQLTemplate) Execute(args ...interface{}) (sql.Result, error) {
	if t.err != nil {
		return nil, t.err
	}
	s, err := openSession(t.dbID)
	if err != nil {
		return nil, err
	}
	defer s.close()

	return s.execute(t.statement, args...)
}

// BulkInsert delegate SQLSession.BulkInsert
func (t SQLTemplate) BulkInsert(rows []interface{}) (sql.Result, error) {
	if t.err != nil {
		return nil, t.err
	}
	s, err := openSession(t.dbID)
	if err != nil {
		return nil, err
	}
	defer s.close()

	return s.bulkInsert(t.statement, rows)
}

// judge if the sql is a sql template
func judgeTemplate(sqlTmpl string) bool {
	return re.FindStringIndex(sqlTmpl) != nil
}

func in(values []interface{}) string {
	mylen := len(values)
	l := make([]string, mylen)
	for i := range values {
		l[i] = "?"
	}
	return strings.Join(l, ",")
}

func driver() string {
	dbmgr := getDBMgrInstance()
	d, _ := dbmgr.getDriver(dbmgr.defaultID)
	return d
}
