package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestMapper(t *testing.T) { TestingT(t) }

type MapperSuite struct{}

var _ = Suite(&MapperSuite{})

func (s MapperSuite) TestParseNamedSQL(c *C) {
	mapper, err := parseNamedSQL("namedsql/mysql/test.xml")
	c.Assert(err, IsNil)
	c.Check(mapper.Namespace, Equals, "/mapper/AuthorMapper")
	c.Check(mapper.Select, HasLen, 1)
	c.Check(mapper.Insert, HasLen, 1)
	c.Check(mapper.Delete, HasLen, 1)
	c.Check(mapper.Update, HasLen, 1)
	c.Check(mapper.Anysql, HasLen, 2)
}
