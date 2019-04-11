package gbatis

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestFlatValue(t *testing.T) { TestingT(t) }

type FlatValueSuite struct{}

var _ = Suite(&FlatValueSuite{})

func (s FlatValueSuite) TestStruct(c *C) {
	type student struct {
		ID   int
		Name string
		Age  int
	}

	stu := student{
		ID:   100,
		Name: "Jack Ju",
		Age:  29,
	}

	fv, err := flatValue(stu)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 3)

	fv, err = flatValue(&stu)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 3)
}

func (s FlatValueSuite) TestMap(c *C) {
	m := make(map[string]string, 0)

	m["a"] = "111"
	m["b"] = "2222"
	m["c"] = "3333"
	m["d"] = "44444"

	fv, err := flatValue(m)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 4)
}

func (s FlatValueSuite) TestBasic(c *C) {
	fv, err := flatValue(100)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 1)

	fv, err = flatValue("Hello World")
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 1)

	fv, err = flatValue(1000.01)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 1)
}

func (s FlatValueSuite) TestArray(c *C) {
	a := []int{100, 200, 300, 400, 500, 600, 700}
	fv, err := flatValue(a)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 7)

	b := a[1:4]
	fv, err = flatValue(b)
	c.Assert(err, IsNil)
	c.Check(fv, HasLen, 3)
}
