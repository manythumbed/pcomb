package pcomb

import (
	. "launchpad.net/gocheck"
	"os"
	"strings"
)

func (s *S) TestEmptyReader(c *C) {
	r := newReader(strings.NewReader(""))

	rune, err := r.take()
	c.Check(err, NotNil)
	c.Check(rune, Equals, 0)
	c.Check(r.current, Equals, r.previous)
	c.Check(r.taken(), Equals, false)

	err = r.untake()
	c.Check(err, NotNil)
	c.Check(r.current, Equals, r.previous)
	c.Check(r.taken(), Equals, false)
}

func (s *S) TestReader(c *C) {
	r := newReader(strings.NewReader("abc"))

	rune, err := r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('a'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.taken(), Equals, true)

	err = r.untake()
	c.Check(err, Equals, nil)
	c.Check(r.current, Equals, r.previous)
	c.Check(r.taken(), Equals, false)
}

func (s *S) TestLineCounting(c *C)	{
	r := newReader(strings.NewReader("a\nbb\nc"))

	rune, err := r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('a'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 1)
	
	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('\n'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 2)

	err = r.untake()
	c.Check(r.current, Equals, r.previous)
	c.Check(r.current.line, Equals, 1)
	c.Check(r.taken(), Equals, false)

	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('\n'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 2)

	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('b'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 2)

	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('b'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 2)

	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('\n'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 3)

	rune, err = r.take()
	c.Check(err, Equals, nil)
	c.Check(rune, Equals, int('c'))
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 3)

	rune, err = r.take()
	c.Check(err, Equals, os.EOF)
	c.Check(rune, Equals, 0)
	c.Check(r.current, Not(Equals), r.previous)
	c.Check(r.current.line, Equals, 3)
}
