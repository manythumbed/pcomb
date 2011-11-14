package pcomb

import (
	"launchpad.net/gocheck"
	"strings"
)

func (s *S) TestReader(c *gocheck.C) {
	r := newReader(strings.NewReader(""))
	c.Check(r.String(), gocheck.Equals, "")
}
