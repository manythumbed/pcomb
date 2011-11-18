package pcomb

import (
	. "launchpad.net/gocheck"
)

func (suite *S) TestNext(c *C)	{
	s := newState("")

	rune, next, ok := s.Next()
	c.Check(rune, Equals, 0)
	c.Check(next, Equals, s)
	c.Check(ok, Equals, false)

	s = newState("\nab\nc\n")
	rune, next, ok = s.Next()
	c.Check(rune, Equals, int('\n'))
	c.Check(next.Line, Equals, 2)
	c.Check(next.Rune, Equals, 0)
	c.Check(next.Column, Equals, 0)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, int('a'))
	c.Check(next.Line, Equals, 2)
	c.Check(next.Rune, Equals, 1)
	c.Check(next.Column, Equals, 1)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, int('b'))
	c.Check(next.Line, Equals, 2)
	c.Check(next.Rune, Equals, 2)
	c.Check(next.Column, Equals, 2)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, int('\n'))
	c.Check(next.Line, Equals, 3)
	c.Check(next.Rune, Equals, 0)
	c.Check(next.Column, Equals, 0)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, int('c'))
	c.Check(next.Line, Equals, 3)
	c.Check(next.Rune, Equals, 1)
	c.Check(next.Column, Equals, 1)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, int('\n'))
	c.Check(next.Line, Equals, 4)
	c.Check(next.Rune, Equals, 0)
	c.Check(next.Column, Equals, 0)
	c.Check(ok, Equals, true)

	rune, next, ok = next.Next()
	c.Check(rune, Equals, 0)
	c.Check(next.Line, Equals, 4)
	c.Check(next.Rune, Equals, 0)
	c.Check(next.Column, Equals, 0)
	c.Check(ok, Equals, false)
}
