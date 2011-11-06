package pcomb

import (
	. "launchpad.net/gocheck"
	"testing"
	"unicode"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {

}

var _ = Suite(&S{})

func (s *S) TestFail(c *C) {
	result := Fail("")
	c.Check(result.Success, Equals, false)
	c.Check(result, NotNil)
}

func (s *S) TestSucceed(c *C) {
	value := "test"
	succeed := Succeed(value)

	result := succeed("")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, value)

	result = succeed("12345")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, value)
}

func (s *S) TestOr(c *C) {
	or := Or(Fail, Succeed("A"))

	result := or("")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "A")

	or = Or(Succeed("A"), Succeed("B"))
	result = or("")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "A")

	or = Or(Succeed("A"), Fail)
	result = or("")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "A")
}

func (s *S) TestItem(c *C) {
	item := Item()
	result := item("")
	c.Check(result.Success, Equals, false)
	c.Check(result.Result, Equals, nil)
	c.Check(result.Remaining, Equals, "")

	result = item("123")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "1")
	c.Check(result.Remaining, Equals, "23")

	result = item(result.Remaining)
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "2")
	c.Check(result.Remaining, Equals, "3")

	result = item(result.Remaining)
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "3")
	c.Check(result.Remaining, Equals, "")
}

func (s *S) TestSatisfy(c *C) {
	parser := Satisfy(unicode.IsDigit)
	result := parser("A")

	c.Check(result.Success, Equals, false)
	c.Check(result.Result, Equals, nil)

	result = parser("1")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "1")

	letter := Satisfy(unicode.IsLetter)
	number := Satisfy(unicode.IsDigit)
	parser = Or(letter, number)

	result = parser("foo")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "f")
	c.Check(result.Remaining, Equals, "oo")

	result = parser("123")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, "1")
	c.Check(result.Remaining, Equals, "23")
}

func (s *S) TestLiteral(c *C) {
	literal := Literal("yes", true)

	result := literal("yes")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, true)
	c.Check(result.Remaining, Equals, "")

	result = literal("no")
	c.Check(result.Success, Equals, false)
	c.Check(result.Result, Equals, nil)

	yesno := Or(Literal("yes", true), Literal("no", false))

	result = yesno("yes")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, true)
	c.Check(result.Remaining, Equals, "")

	result = yesno("no")
	c.Check(result.Success, Equals, true)
	c.Check(result.Result, Equals, false)
	c.Check(result.Remaining, Equals, "")
}

func (s *S) TestMany(c *C) {
	many := Many(Literal("*", "star"))

	result := many("123")
	c.Check(result.Success, Equals, true)
	slice, ok := result.Result.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 0)
	c.Check(result.Remaining, Equals, "123")

	result = many("***1*2*3")
	c.Check(result.Success, Equals, true)
	slice, ok = result.Result.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 3)
	c.Check(slice[0], Equals, "star")
	c.Check(slice[1], Equals, "star")
	c.Check(slice[2], Equals, "star")

	c.Check(result.Remaining, Equals, "1*2*3")
}

func (s *S) TestMany1(c *C) {
	many := Many1(Literal("*", "star"))

	result := many("!23")
	c.Check(result.Success, Equals, false)

	result = many("**23")
	c.Check(result.Success, Equals, true)
	slice, ok := result.Result.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 2)
	c.Check(slice[0], Equals, "star")
	c.Check(slice[1], Equals, "star")
}
