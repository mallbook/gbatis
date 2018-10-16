package gbatis

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type sqlClass int8

const (
	selectClass    sqlClass = 0
	insertClass    sqlClass = 1
	updateClass    sqlClass = 2
	deleteClass    sqlClass = 3
	procStoreClass sqlClass = 4
)

// NamedSQLInfo means NamedSQL infomation
type sqlInfor struct {
	class      sqlClass
	sql        string
	resultType string
}

// namedSQL include all namedSQL
// var namedSQL = make(map[string]sqlInfor)

var (
	_sqlMgrInst *sqlMgr
	_sqlMgrOnce sync.Once
)

type sqlMgr struct {
	sync.RWMutex
	sqls map[string]sqlInfor
}

func getSQLMgrInstance() *sqlMgr {

	_sqlMgrOnce.Do(func() {
		_sqlMgrInst = &sqlMgr{
			sqls: make(map[string]sqlInfor),
		}
	})
	return _sqlMgrInst
}

func (s *sqlMgr) getSQL(sqlID string) (si sqlInfor, err error) {
	s.RLock()
	defer s.RUnlock()

	si, ok := s.sqls[sqlID]
	if !ok {
		return si, fmt.Errorf("Not found sql named %s", sqlID)
	}
	return si, nil
}

func (s *sqlMgr) appendSQL(m *mapper) {

	s.Lock()
	defer s.Unlock()

	for _, item := range m.Insert {
		infor := sqlInfor{
			class: insertClass,
			sql:   formatSQL(item.SQL),
		}
		s.sqls[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Delete {
		infor := sqlInfor{
			class: deleteClass,
			sql:   formatSQL(item.SQL),
		}
		s.sqls[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Update {
		infor := sqlInfor{
			class: updateClass,
			sql:   formatSQL(item.SQL),
		}
		s.sqls[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Select {
		infor := sqlInfor{
			class:      selectClass,
			resultType: item.ResultType,
			sql:        formatSQL(item.SQL),
		}
		s.sqls[m.Namespace+"."+item.ID] = infor
	}
}

func (s *sqlMgr) display() {
	s.RLock()
	defer s.RUnlock()

	for id, si := range s.sqls {
		log.Printf("sqlID=%s\n\t%s\n", id, si.sql)
	}
}

func formatSQL(sql string) string {
	segs := strings.Split(sql, "\n")
	for i, seg := range segs {
		segs[i] = strings.Trim(seg, "\t ")
	}
	sql = strings.Join(segs, "\n")
	return strings.Trim(sql, "\r\n")
}

/*
func appendNamedSQL(m *mapper) {

	for _, item := range m.Insert {
		infor := sqlInfor{
			class: insertClass,
			sql:   item.SQL,
		}
		namedSQL[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Delete {
		infor := sqlInfor{
			class: deleteClass,
			sql:   item.SQL,
		}
		namedSQL[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Update {
		infor := sqlInfor{
			class: updateClass,
			sql:   item.SQL,
		}
		namedSQL[m.Namespace+"."+item.ID] = infor
	}
	for _, item := range m.Select {
		infor := sqlInfor{
			class: selectClass,
			sql:   item.SQL,
		}
		namedSQL[m.Namespace+"."+item.ID] = infor
	}
}*/

/*type namedSQLMgr struct {
	namedSQL map[string]sqlInfor
}

func (n *namedSQLMgr) (m *mapper) {

}*/

func walk(path string) ([]string, error) {
	namedFiles := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			// log.Println("nil")
			return err
		}
		if f.IsDir() {
			return nil
		}
		// log.Println(path)
		namedFiles = append(namedFiles, path)
		return nil
	})
	return namedFiles, err
}

func loadNamedSQL(namedPath string) error {
	namedFiles, err := walk(namedPath)
	if err != nil {
		log.Printf("Walk named sql fail, err=%s, namedPath=%s", err, namedPath)
		return err
	}

	s := getSQLMgrInstance()

	for index, namedFile := range namedFiles {
		log.Println(index, namedFile)
		m, err := parseNamedSQL(namedFile)
		if err != nil {
			log.Printf("Parse named sql fail, namedFile=%s, err=%s", namedFile, err)
			continue
		}
		// appendNamedSQL(m)
		s.appendSQL(m)
	}

	// log.Println(s)

	return nil
}

func init() {
	loadNamedSQL("namedsql")
}
