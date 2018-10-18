package gbatis_test

import (
	"fmt"
	"testing"

	"github.com/mallbook/gbatis/uuid"

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

	row, err := s.SelectRow("mapper.mall.selectMall", "4")
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

	results, err := s.Select("mapper.mall.selectMall", "4")
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

func TestInsert_1(t *testing.T) {
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	u := uuid.New()

	r, err := s.Execute("mapper.mall.insertMall", u.String(), "xxmall", "mallbook.com/img/xx.png", "hi xxmall")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Insert success, rowsAffected=%d\n", rowsAffected)
}

func TestUpdate_1(t *testing.T) {
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	r, err := s.Execute("mapper.mall.updateMall", "New Name", "4")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Update success, rowsAffected=%d\n", rowsAffected)
}

func TestDelete_1(t *testing.T) {
	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	r, err := s.Execute("mapper.mall.deleteMall", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Delete success, rowsAffected=%d\n", rowsAffected)
}

func TestBulkInsert_1(t *testing.T) {

	rows := make([]interface{}, 0)

	for i := 0; i < 10; i++ {
		row := make([]interface{}, 0)
		u := uuid.New()
		row = append(row, u.String(), "a mall", "a.png", "hi a mall")
		rows = append(rows, row)
	}

	s, err := gbatis.OpenSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	r, err := s.BulkInsert("mapper.mall.insertMall", rows)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("BulkInsert rows affected is %d\n", rowsAffected)
}

// func TestMain(m *testing.M) {
func init() {
	err := gbatis.OpenDB("examples/gbatis.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Open database success.")
}
