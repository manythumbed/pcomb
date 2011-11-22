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
	result := Fail.parse("")
	c.Check(result.Success, Equals, false)
	c.Check(result.Errors, Equals, NoErrors())
	c.Check(result.Value, Equals, nil)
}

func (s *S) TestReturn(c *C) {
	value := "test"
	succeed := Return(value)

	result := succeed.parse("")
	c.Check(result.Success, Equals, true)
	c.Check(result.Errors, Equals, NoErrors())
	c.Check(result.Value, Equals, value)

	result = succeed.parse("12345")
	c.Check(result.Success, Equals, true)
	c.Check(result.Errors, Equals, NoErrors())
	c.Check(result.Value, Equals, value)
}

func (s *S) TestSatisfy(c *C) {
	sat := Satisfy(unicode.IsDigit)

	result := sat.parse("")
	c.Check(result.Success, Equals, false)

	result = sat.parse("a")
	c.Check(result.Success, Equals, false)
	c.Check(result.Errors[0].Position, Equals, Position{1, 1, 1})

	result = sat.parse("1")
	c.Check(result.Success, Equals, true)
	c.Check(result.Errors, Equals, NoErrors())
}

func (s *S) TestOr(c *C) {
	letter := Satisfy(unicode.IsLetter)
	number := Satisfy(unicode.IsNumber)
	or := Or(letter, number)

	result := or.parse("")
	c.Check(result.Success, Equals, false)
	c.Check(result.Value, Equals, nil)
	c.Check(result.Errors, NotNil)
	c.Check(len(result.Errors), Equals, 1)
	c.Check(len(result.Errors[0].Errors), Equals, 2)

	result = or.parse("A")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, int('A'))

	result = or.parse("1")
	c.Check(result.Success, Equals, false)
	c.Check(result.Value, Equals, nil)
	c.Check(result.Errors, NotNil)
	c.Check(len(result.Errors), Equals, 1)
}

func (s *S) TestItem(c *C) {
	item := Item()
	result := item.parse("")
	c.Check(result.Success, Equals, false)
	c.Check(result.Value, Equals, nil)

	result = item.parse("123")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, int('1'))

	result = item.parse(result.State.text)
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, int('2'))

	result = item.parse(result.State.text)
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, int('3'))

	result = item.parse(result.State.text)
	c.Check(result.Success, Equals, false)
	c.Check(result.Value, Equals, nil)
}

func (s *S) TestTry(c *C) {
	letter := Try(Satisfy(unicode.IsLetter))

	r := letter.parse("1")
	c.Check(r.Success, Equals, false)
	c.Check(r.State.text, Equals, "1")
	c.Check(r.Value, Equals, nil)

	number := Try(Satisfy(unicode.IsDigit))
	r = number.parse(r.State.text)
	c.Check(r.Success, Equals, true)
	c.Check(r.State.text, Equals, "")
	c.Check(r.Value, Equals, int('1'))
}

/*
func (s *S) TestLiteral(c *gocheck.C) {
	literal := Literal("yes", true)

	result := literal.parse("yes")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, true)
	c.Check(result.Remaining.Input, gocheck.Equals, "")

	result = literal.parse("no")
	c.Check(result.Success, gocheck.Equals, false)
	c.Check(result.Value, gocheck.Equals, nil)

	yesno := Or(Literal("yes", true), Literal("no", false))

	result = yesno.parse("yes")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, true)
	c.Check(result.Remaining.Input, gocheck.Equals, "")

	result = yesno.parse("no")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, false)
	c.Check(result.Remaining.Input, gocheck.Equals, "")
}

func (s *S) TestMany(c *gocheck.C) {
	many := Many(Literal("*", "star"))

	result := many.parse("123")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(len(slice), gocheck.Equals, 0)
	c.Check(result.Remaining.Input, gocheck.Equals, "123")

	result = many.parse("***1*2*3")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok = result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(len(slice), gocheck.Equals, 3)
	c.Check(slice[0], gocheck.Equals, "star")
	c.Check(slice[1], gocheck.Equals, "star")
	c.Check(slice[2], gocheck.Equals, "star")

	c.Check(result.Remaining.Input, gocheck.Equals, "1*2*3")
}

func (s *S) TestMany1(c *gocheck.C) {
	many := Many1(Literal("*", "star"))

	result := many.parse("!23")
	c.Check(result.Success, gocheck.Equals, false)

	result = many.parse("**23")
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
	result := p.parse("1234")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "Y")

	p1 := Succeed("Y").Then(succeed)
	result = p1.parse("1234")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, "Y")
}

func (s *S) TestChain(c *gocheck.C) {
	one := Literal("1", 1)
	two := Literal("2", 2)
	number := Or(one, two)

	result := number.parse("1")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 1)

	result = number.parse("2")
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

	result = expr.parse("1+2")
	c.Check(result.Success, gocheck.Equals, true)
	c.Check(result.Value, gocheck.Equals, 3)

	result = expr.parse("1+2-2")
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

	result := listOfNumbers.parse("1,2,3,4")
	c.Check(result.Success, gocheck.Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, gocheck.Equals, true)
	c.Check(slice, gocheck.Equals, []interface{}{1, 2, 3, 4})
}
*/
