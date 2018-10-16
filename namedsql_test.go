package gbatis

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	namedFiles, err := walk("namedsql/mysql")
	if err != nil {
		t.Fail()
	} else {
		t.Log(namedFiles)
		t.Log("OK")
	}
}

func TestLoadNamedSQL(t *testing.T) {
	err := loadNamedSQL("namedsql")
	if err != nil {
		t.Fail()
		return
	}

	s := getSQLMgrInstance()
	s.display()

	si, err := s.getSQL("mapper.AuthorMapper.updateAuthor")
	if err != nil {
		t.Fail()
		return
	}

	fmt.Printf("[%s]", formatSQL(si.sql))
}
