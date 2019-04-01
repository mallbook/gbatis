package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestNamedSQL(t *testing.T) { TestingT(t) }

type namedSQLSuite struct{}

var _ = Suite(&namedSQLSuite{})

func (s namedSQLSuite) TestNamedSql(c *C) {
	sm := getSQLMgrInstance()
	si, err := sm.getSQL("/mapper/mall.selectMall")
	c.Assert(err, IsNil)
	c.Assert(si, NotNil)
	c.Check(si.class, Equals, selectClass)

	_, err = sm.getSQL("/mapper/mall.selectMall2")
	c.Assert(err, ErrorMatches, "Not found sql named /mapper/mall.selectMall2")
}
