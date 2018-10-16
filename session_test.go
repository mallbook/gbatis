package gbatis_test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mallbook/gbatis"
	_ "github.com/mallbook/gbatis/bean"
)

func TestOpenSession(t *testing.T) {
	s, err := gbatis.OpenSession()
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()
}

func TestSelectOne(t *testing.T) {
	//loadNamedSQL("namedsql")
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	row, err := s.SelectRow("mapper.mall.selectMall", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id string
	var name string
	var avatar string
	var story string
	err = row.Scan(&id, &name, &avatar, &story)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("id = %s, name=%s, avatar=%s, story=%s\n", id, name, avatar, story)
}

func TestSelect(t *testing.T) {
	// loadNamedSQL("namedsql")
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	results, err := s.Select("mapper.mall.selectMall", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results)
}

func TestSelect_2(t *testing.T) {
	// loadNamedSQL("namedsql")
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	results, err := s.Select("mapper.mall.selectAllMalls")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results)
}
