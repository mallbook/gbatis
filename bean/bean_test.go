package bean

import (
	"testing"

	. "gopkg.in/check.v1"
)

type student struct {
	id   string
	name string
	age  int
}

func newStudent() *student {
	return &student{}
}

func init() {
	Register("student", newStudent)
}

func TestBean(t *testing.T) { TestingT(t) }

type beanSuite struct{}

var _ = Suite(&beanSuite{})

func (s beanSuite) TestNewBean(c *C) {

	f := getBeanFactoryInstance()
	v, err := f.newBean("student")
	c.Assert(err, IsNil)
	c.Check(v, NotNil)

	_, err = f.newBean("student2")
	c.Assert(err, ErrorMatches, "Not found bean named student2")
}
