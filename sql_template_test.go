package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestSQLTemplate(t *testing.T) { TestingT(t) }

type sqlTemplateSuite struct{}

var _ = Suite(sqlTemplateSuite{})

func (s sqlTemplateSuite) TestTemplate(c *C) {
	data := make(map[string]string)
	data["tableName"] = "t_mall"
	t := Template("/mapper/mall.updateMall2", data)
	c.Assert(t.err, IsNil)
	c.Check(t.statement, Equals, "update t_mall set name = ?, updatedAt = unix_timestamp(now()) where id = ?")

	t2 := Template("/mapper/mall.selectAllMalls2", data)
	c.Assert(t2.err, IsNil)
	c.Check(t2.statement, Equals, "SELECT id as ID, name as Name, avatar as Avatar, story as Story FROM t_mall")
}

func (s sqlTemplateSuite) TestTemplateErr(c *C) {
	data := make(map[string]string)
	data["tableName"] = "t_mall"
	t := Template("/mapper/mall.updateMall", data)
	c.Assert(t.err.Error(), Equals, "The sql(/mapper/mall.updateMall) is not a sql template")

	t2 := Template("/mapper/mall.selectAllMalls", data)
	c.Assert(t2.err.Error(), Equals, "The sql(/mapper/mall.selectAllMalls) is not a sql template")
}

func (s sqlTemplateSuite) TestGetTmpl(c *C) {
	t, si, err := newSQLTemplateMgr().getTmpl("/mapper/mall.updateMall2")
	c.Assert(err, IsNil)
	c.Check(si.class, Equals, updateClass)

	t2, si2, err := newSQLTemplateMgr().getTmpl("/mapper/mall.updateMall2")
	c.Assert(err, IsNil)
	c.Check(si2.class, Equals, updateClass)

	c.Check(t, Equals, t2)
}

func (s sqlTemplateSuite) TestJudgeTemplate(c *C) {
	type testCase struct {
		tmpl   string
		result bool
	}

	testCases := []testCase{
		{"${cc}", true},
		{"select * from ${.name}", true},
		{"select id from t_mall", false},
		{"select ${.fields} from ${.name}", true},
	}

	for _, testCase := range testCases {
		c.Check(judgeTemplate(testCase.tmpl), Equals, testCase.result)
	}
}
