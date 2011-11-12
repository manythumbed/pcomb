package pcomb

import (
	"launchpad.net/gocheck"
	"testing"
	"unicode"
)

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type S struct {

}

var _ = gocheck.Suite(&S{})

func (s *S) TestFail(c *gocheck.C) {
	result := Fail("")
	c.Check(result.Success, gocheck.Equals, false)
	c.Check(result, gocheck.NotNil)
}

func (s *S) TestSucceed(c *gocheck.C) {
	value := "test"
	succeed := Succeed(value)

	result := succeed("")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, value)

	result = succeed("12345")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, value)
}

func (s *S) TestOr(c *gocheck.C) {
	or := Or(Fail, Succeed("A"))

	result := or("")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "A")

	or = Or(Succeed("A"), Succeed("B"))
	result = or("")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "A")

	or = Or(Succeed("A"), Fail)
	result = or("")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "A")
}

func (s *S) TestItem(c *gocheck.C) {
	item := Item()
	result := item("")
	c.Check(result.Success, gocheck.Equals, false)
	c.Check(result.Value, gocheck.Equals, nil)
	c.Check(result.Remaining, gocheck.Equals, "")

	result = item("123")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "1")
	c.Check(result.Remaining, gocheck.Equals, "23")

	result = item(result.Remaining)
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "2")
	c.Check(result.Remaining, gocheck.Equals, "3")

	result = item(result.Remaining)
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "3")
	c.Check(result.Remaining, gocheck.Equals, "")
}

func (s *S) TestSatisfy(c *gocheck.C) {
	parser := Satisfy(unicode.IsDigit)
	result := parser("A")

	c.Check(result.Success, gocheck.Equals, false)
	c.Check(result.Value, gocheck.Equals, nil)

	result = parser("1")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "1")

	letter := Satisfy(unicode.IsLetter)
	number := Satisfy(unicode.IsDigit)
	parser = Or(letter, number)

	result = parser("foo")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "f")
	c.Check(result.Remaining, gocheck.Equals, "oo")

	result = parser("123")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "1")
	c.Check(result.Remaining, gocheck.Equals, "23")
}

func (s *S) TestLiteral(c *gocheck.C) {
	literal := Literal("yes", true)

	result := literal("yes")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, true)
	c.Check(result.Remaining, gocheck.Equals, "")

	result = literal("no")
	c.Check(result.Success, gocheck.Equals, false)
	c.Check(result.Value, gocheck.Equals, nil)

	yesno := Or(Literal("yes", true), Literal("no", false))

	result = yesno("yes")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, true)
	c.Check(result.Remaining, gocheck.Equals, "")

	result = yesno("no")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, false)
	c.Check(result.Remaining, gocheck.Equals, "")
}

func (s *S) TestMany(c *gocheck.C) {
	many := Many(Literal("*", "star"))

	result := many("123")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(len(slice), gocheck.Equals, 0)
	c.Check(result.Remaining, gocheck.Equals, "123")

	result = many("***1*2*3")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok = result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(len(slice), gocheck.Equals, 3)
	c.Check(slice[0], gocheck.Equals, "star")
	c.Check(slice[1], gocheck.Equals, "star")
	c.Check(slice[2], gocheck.Equals, "star")

	c.Check(result.Remaining, gocheck.Equals, "1*2*3")
}

func (s *S) TestMany1(c *gocheck.C) {
	many := Many1(Literal("*", "star"))

	result := many("!23")
	c.Check(result.Success, gocheck.Equals, false)

	result = many("**23")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(len(slice), gocheck.Equals, 2)
	c.Check(slice[0], gocheck.Equals, "star")
	c.Check(slice[1], gocheck.Equals, "star")
}

func (s *S) TestThen(c *gocheck.C) {
	succeed := func(x interface{}) Parser {
		return Succeed(x)
	}

	p := Then(Succeed("Y"), succeed)
	result := p("1234")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "Y")

	p1 := Succeed("Y").Then(succeed)
	result = p1("1234")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "Y")
}

func (s *S) TestChain(c *gocheck.C) {
	one := Literal("1", 1)
	two := Literal("2", 2)
	number := Or(one, two)

	result := number("1")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 1)

	result = number("2")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 2)

	addFunc := func(x_val, y_val interface{}) interface{} {
		x, _ := x_val.(int)
		y, _ := y_val.(int)
		return x + y
	}
	minusFunc := func(x_val, y_val interface{}) interface{} {
		x, _ := x_val.(int)
		y, _ := y_val.(int)
		return x - y
	}
	add := Then_(Literal("+", nil), Succeed(addFunc))
	minus := Then_(Literal("-", nil), Succeed(minusFunc))
	op := Or(add, minus)

	expr := ChainLeft1(number, op)

	result = expr("1+2")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 3)

	result = expr("1+2-2")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 1)
}

func (s *S) TestSepBy(c *gocheck.C) {
	one := Literal("1", 1)
	two := Literal("2", 2)
	three := Literal("3", 3)
	four := Literal("4", 4)
	number := Or(Or(one, two), Or(three, four))

	listOfNumbers := SeperatedBy(number, Literal(",", nil))

	result := listOfNumbers("1,2,3,4")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(slice, gocheck.Equals, []interface{}{1, 2, 3, 4})
}

func (s *S) TestConsumed(c *gocheck.C) {
	consumed := Consumed{}

	c.Check(consumed.ConsumedInput, gocheck.Equals, false)
}
