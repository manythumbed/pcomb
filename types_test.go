package pcomb

import (
	. "launchpad.net/gocheck"
	"testing"
	"unicode"
)

func Test(t *testing.T)	{
	TestingT(t)
}

type S struct {
}

var _ = Suite(&S{})

func (s *S) TestFail(c *C) {
	success, result := Fail("")
	c.Check(success, Equals, false)
	c.Check(result, NotNil)
}

func (s *S) TestSucceed(c *C) {
	value := "test"
	succeed := Succeed(value)

	success, result := succeed("")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, value)

	success, result = succeed("12345")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, value)
}

func (s *S) TestOr(c *C) {
	or := Or(Fail, Succeed("A"))

	success, result := or("")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "A")

	or = Or(Succeed("A"), Succeed("B"))
	success, result = or("")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "A")

	or = Or(Succeed("A"), Fail)
	success, result = or("")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "A")
}

func (s *S) TestItem(c *C) {
	item := Item()
	success, result := item("")
	c.Check(success, Equals, false)
	c.Check(result.Result, Equals, nil)
	c.Check(result.Remaining, Equals, "")

	success, result = item("123")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "1")
	c.Check(result.Remaining, Equals, "23")

	success, result = item(result.Remaining)
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "2")
	c.Check(result.Remaining, Equals, "3")

	success, result = item(result.Remaining)
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "3")
	c.Check(result.Remaining, Equals, "")
}

func (s *S) TestSatisfy(c *C) {
	parser := Satisfy(unicode.IsDigit)
	success, result := parser("A")

	c.Check(success, Equals, false)
	c.Check(result.Result, Equals, nil)

	success, result = parser("1")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "1")

	letter := Satisfy(unicode.IsLetter)
	number := Satisfy(unicode.IsDigit)
	parser = Or(letter, number)

	success, result = parser("foo")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "f")
	c.Check(result.Remaining, Equals, "oo")

	success, result = parser("123")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, "1")
	c.Check(result.Remaining, Equals, "23")
}

func (s *S) TestLiteral(c *C) {
	literal := Literal("yes", true)

	success, result := literal("yes")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, true)
	c.Check(result.Remaining, Equals, "")

	success, result = literal("no")
	c.Check(success, Equals, false)
	c.Check(result.Result, Equals, nil)

	yesno := Or(Literal("yes", true), Literal("no", false))

	success, result = yesno("yes")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, true)
	c.Check(result.Remaining, Equals, "")

	success, result = yesno("no")
	c.Check(success, Equals, true)
	c.Check(result.Result, Equals, false)
	c.Check(result.Remaining, Equals, "")

}
