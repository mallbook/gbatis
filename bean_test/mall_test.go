package bean1

import (
	"testing"

	"github.com/mallbook/gbatis/bean"
	. "gopkg.in/check.v1"
)

func TestBean(t *testing.T) { TestingT(t) }

type beanSuite struct{}

var _ = Suite(&beanSuite{})

func (s beanSuite) TestNewBean(c *C) {

	v, err := bean.New("bean.Mall")
	c.Assert(err, IsNil)
	c.Check(v, NotNil)

	v, err = bean.New("bean.Shop")
	c.Assert(err, IsNil)
	c.Check(v, NotNil)

	v, err = bean.New("bean.Brand")
	c.Assert(err, IsNil)
	c.Check(v, NotNil)
}
