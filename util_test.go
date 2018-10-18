package gbatis

import (
	"fmt"
	"testing"
)

func TestCutSQL(t *testing.T) {
	sql := `insert into t_mall(id, name, avatar, createdAt, updatedAt, story) 
	values     
	(?, ?, ?, unix_timestamp(now()), unix_timestamp(now()), ?)`
	s, err := cutInsertSQL(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestCutSQL_1(t *testing.T) {
	sql := "hello world"
	s, err := cutInsertSQL(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
