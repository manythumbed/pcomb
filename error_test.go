package pcomb

import (
	. "launchpad.net/gocheck"
)

func (suite *S) TestNewError(c *C)	{
	s := newState("")
	
	m := "Test an error"
	e := NewError(s.Position, m)
	c.Check(e.Line, Equals, 1)
	c.Check(e.Rune, Equals, 0)
	c.Check(e.Column, Equals, 0)
	c.Check(e.Errors, NotNil)
	c.Check(len(e.Errors), Equals, 0)
}

func (suite *S) TestAppend(c *C)	{
	e := NewError(newState("").Position, "Append")
	c.Check(e.Errors, NotNil)
	c.Check(len(e.Errors), Equals, 0)

	e1 := NewError(newState("").Position, "Append 1")
	e.Append(e1)
	c.Check(e.Errors, NotNil)
	c.Check(len(e.Errors), Equals, 1)
	c.Check(e.Errors[0], Equals, e1)

	e2 := NewError(newState("").Position, "Append 2")
	es := []Error{e2, e1}
	e.Append(es...)
	c.Check(e.Errors, NotNil)
	c.Check(len(e.Errors), Equals, 3)
	c.Check(e.Errors[0], Equals, e1)
	c.Check(e.Errors[1], Equals, e2)
	c.Check(e.Errors[2], Equals, e1)
}
