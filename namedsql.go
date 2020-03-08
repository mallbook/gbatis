package gbatis

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mallbook/commandline"
)

type sqlClass int8

const (
	selectClass sqlClass = iota
	insertClass
	updateClass
	deleteClass
	procStoreClass
	anyClass
)

type sqlInfor struct {
	class      sqlClass
	sql        string
	resultType string
}

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

	for _, item := range m.Anysql {
		infor := sqlInfor{
			class: anyClass,
			sql:   formatSQL(item.SQL),
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
	// TODO: temp for test
	// p := commandline.PrefixPath()
	// loadNamedSQL(p + "/namedsql")
}
