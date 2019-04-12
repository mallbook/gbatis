package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestDyncSQL(t *testing.T) { TestingT(t) }

type DyncSQLSuite struct{}

var _ = Suite(&DyncSQLSuite{})

func (s DyncSQLSuite) TestPreparedDyncSQL(c *C) {
	sql := `SELECT * FROM BLOG
<where>  
AND title like #{.title}
AND author_name like #{.author.name}
</where>`

	expect := `SELECT * FROM BLOG
WHERE title like #{.title}
AND author_name like #{.author.name}`

	statement, err := preparedDyncSQL(sql)
	c.Assert(err, IsNil)
	c.Check(statement, Equals, expect)
}

func (s DyncSQLSuite) TestPreparedSetStatement(c *C) {
	sql := `update Author
<set>
	username=#{username},
	password=#{password},
	email=#{email},
</set>
where id=#{id}`

	expect := `update Author
SET username=#{username},
	password=#{password},
	email=#{email}
where id=#{id}`

	statement, err := preparedDyncSQL(sql)
	c.Assert(err, IsNil)
	// fmt.Println(statement)
	c.Check(statement, Equals, expect)
}

func (s DyncSQLSuite) TestPreparedSetStatement2(c *C) {
	sql := `update Author <set>
	username=#{username},
	password=#{password},
	email=#{email},
</set> where id=#{id}`

	expect := `update Author SET username=#{username},
	password=#{password},
	email=#{email} where id=#{id}`

	statement, err := preparedDyncSQL(sql)
	c.Assert(err, IsNil)
	// fmt.Println(statement)
	c.Check(statement, Equals, expect)
}
