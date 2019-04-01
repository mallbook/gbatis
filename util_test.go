package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

// func TestCutSQL(t *testing.T) {
// 	sql := `insert into t_mall(id, name, avatar, createdAt, updatedAt, story)
// 	values
// 	(?, ?, ?, unix_timestamp(now()), unix_timestamp(now()), ?)`
// 	s, err := cutInsertSQL(sql)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(s)
// }

// func TestCutSQL_1(t *testing.T) {
// 	sql := "hello world"
// 	s, err := cutInsertSQL(sql)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(s)
// }

func TestUtil(t *testing.T) { TestingT(t) }

type utilSuite struct{}

var _ = Suite(&utilSuite{})

func (s utilSuite) TestCutSQL(c *C) {
	sql := `insert into t_mall(id, name, avatar, createdAt, updatedAt, story) 
	values     
	(?, ?, ?, unix_timestamp(now()), unix_timestamp(now()), ?)`
	r, err := cutInsertSQL(sql)
	c.Assert(err, IsNil)
	c.Assert(r, NotNil)
	c.Check(r, HasLen, 2)
}

func (s utilSuite) TestCutSQL_1(c *C) {
	sql := "hello world"
	_, err := cutInsertSQL(sql)
	c.Assert(err, NotNil)
	c.Check(err, ErrorMatches, "Not be a valid insert statement, sql=hello world")
}
