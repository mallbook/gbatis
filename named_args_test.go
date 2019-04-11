package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestNamedArgs(t *testing.T) { TestingT(t) }

type NamedArgSuite struct{}

var _ = Suite(&NamedArgSuite{})

func (s NamedArgSuite) TestParse(c *C) {
	sql := "select id, name, age from student where id=#{.ID} and name=#{.Name}"
	params := parseNamedArgs(sql)
	c.Assert(params, NotNil)
	c.Assert(params, HasLen, 2)
	c.Check(params[0], Equals, ".ID")
	c.Check(params[1], Equals, ".Name")
}

func (s NamedArgSuite) TestReplace(c *C) {
	sql := "select id, name, age from student where id=#{.ID} and name=#{.Name}"
	sql1 := replaceNamedArgs(sql)

	expect := "select id, name, age from student where id=? and name=?"
	c.Check(sql1, Equals, expect)

	sql = "select id, name, age from ${.tableName} where id=#{.ID} and name=#{.Name}"
	sql2 := replaceNamedArgs(sql)
	expect = "select id, name, age from ${.tableName} where id=? and name=?"
	c.Check(sql2, Equals, expect)
}

func (s NamedArgSuite) TestContain(c *C) {
	sql := "select id, name, age from student where id=#{.ID} and name=#{.Name}"
	b := containNamedArgs(sql)
	c.Check(b, Equals, true)

	sql = "select id, name, age from student where id=1"
	b = containNamedArgs(sql)
	c.Check(b, Equals, false)
}

func (s NamedArgSuite) TestPrepared(c *C) {
	sql := "select id, name, age from student where id=#{.ID} and name=#{.Name}"
	data := make(map[string]interface{})
	data["ID"] = 100
	data["Name"] = "Jack Ju"

	statement, args, err := preparedNamedArgs(sql, data)
	c.Assert(err, IsNil)
	c.Check(statement, Equals, "select id, name, age from student where id=? and name=?")
	c.Assert(args, HasLen, 2)
	c.Check(args[0], Equals, 100)
	c.Check(args[1], Equals, "Jack Ju")
}

func (s NamedArgSuite) TestPrepared2(c *C) {
	type student struct {
		ID   int
		Name string
		Age  int
	}
	sql := "select id, name, age from student where id=#{.ID} and name=#{.Name}"
	data := student{
		ID:   100,
		Name: "Jack Ju",
		Age:  29,
	}

	statement, args, err := preparedNamedArgs(sql, data)
	c.Assert(err, IsNil)
	c.Check(statement, Equals, "select id, name, age from student where id=? and name=?")
	c.Assert(args, HasLen, 2)
	c.Check(args[0], Equals, 100)
	c.Check(args[1], Equals, "Jack Ju")
}

func (s NamedArgSuite) TestPreparedBulk(c *C) {
	type student struct {
		ID   int
		Name string
		Age  int
	}

	sql := "insert into student(id, name, age) values(#{.ID}, #{.Name}, #{.Age})"

	datas := make([]interface{}, 0)
	for i := 0; i < 100; i++ {
		stu := student{
			ID:   i,
			Name: "Jack Ju",
			Age:  19,
		}
		datas = append(datas, stu)
	}

	expectSQL := "insert into student(id, name, age) values(?, ?, ?)"

	statement, rows, err := preparedBulkNamedArgs(sql, datas)
	c.Assert(err, IsNil)
	c.Check(statement, Equals, expectSQL)
	c.Assert(rows, HasLen, 100)
	c.Assert(rows[99], HasLen, 3)
	row, ok := rows[99].([]interface{})
	c.Assert(ok, Equals, true)
	c.Check(row[0], Equals, 99)
}
