package pcomb

import (
	"launchpad.net/gocheck"
	"strings"
)

func (s *S) TestEmptyReader(c *gocheck.C) {
	r := newReader(strings.NewReader(""))

	rune, err := r.take()
	c.Check(err, gocheck.NotNil)
	c.Check(rune, gocheck.Equals, 0)
	c.Check(r.current, gocheck.Equals, r.previous)

	err = r.untake()
	c.Check(err, gocheck.NotNil)
	c.Check(r.current, gocheck.Equals, r.previous)
}

func (s *S) TestReader(c *gocheck.C) {
	r := newReader(strings.NewReader("abc"))

	rune, err := r.take()
	c.Check(err, gocheck.Equals, nil)
	c.Check(rune, gocheck.Equals, 97)
	c.Check(r.current, gocheck.Not(gocheck.Equals), r.previous)

	err = r.untake()
	c.Check(err, gocheck.Equals, nil)
	c.Check(r.current, gocheck.Equals, r.previous)
}
