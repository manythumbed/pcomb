package pcomb

import (
	. "launchpad.net/gocheck"
	"testing"
	"unicode"
	"fmt"
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

func (s *S) TestTag(c *C) {
	label := "letter"
	letter := Tag(Try(Satisfy(unicode.IsLetter)), label)

	r := letter.parse("1")
	c.Check(r.Success, Equals, false)
	c.Check(r.Errors, NotNil)
	c.Check(len(r.Errors), Equals, 1)
	c.Check(r.Errors[0].Message, Equals, fmt.Sprintf(labelFormat, label))

	r = letter.parse("v")
	c.Check(r.Success, Equals, true)
	c.Check(r.Value, Equals, int('v'))
}

func (s *S) TestSequence(c *C) {
	succeed := func(x interface{}) Parser {
		return Return(x)
	}

	p := Sequence(Return("Y"), succeed)
	result := p.parse("1234")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, "Y")

	letter := Satisfy(unicode.IsLetter)
	number := Satisfy(unicode.IsNumber)

	combine := func(x interface{}) Parser {
		return number
	}
	p = Sequence(letter, combine)
	r := p.parse("a1")
	c.Check(r.Success, Equals, true)
	c.Check(r.Value, Equals, int('1'))
}

func (s *S) TestLiteral(c *C) {
	literal := Literal("yes", true)

	result := literal.parse("yes")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, true)
	c.Check(result.State.text, Equals, "")

	result = literal.parse("no")
	c.Check(result.Success, Equals, false)
	c.Check(result.Value, Equals, nil)

	yesno := Or(Try(Literal("yes", true)), Try(Literal("no", false)))

	result = yesno.parse("yes")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, true)
	c.Check(result.State.text, Equals, "")

	result = yesno.parse("no")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, false)
	c.Check(result.State.text, Equals, "")
}

func (s *S) TestZeroOrMore(c *C) {
	many := ZeroOrMore(Try(Literal("*", "star")))

	result := many.parse("123")
	c.Check(result.Success, Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 0)
	c.Check(result.State.text, Equals, "123")

	result = many.parse("***1*2*3")
	c.Check(result.Success, Equals, true)
	slice, ok = result.Value.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 3)
	c.Check(slice[0], Equals, "star")
	c.Check(slice[1], Equals, "star")
	c.Check(slice[2], Equals, "star")

	c.Check(result.State.text, Equals, "1*2*3")
}

func (s *S) TestOneOrMore(c *C) {
	many := OneOrMore(Try(Literal("*", "star")))

	result := many.parse("!23")
	c.Check(result.Success, Equals, false)

	result = many.parse("**23")
	c.Check(result.Success, Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(len(slice), Equals, 2)
	c.Check(slice[0], Equals, "star")
	c.Check(slice[1], Equals, "star")
}

func (s *S) TestChain(c *C) {
	one := Literal("1", 1)
	two := Literal("2", 2)
	number := Or(Try(one), Try(two))

	result := number.parse("1")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, 1)

	result = number.parse("2")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, 2)

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
	add := Sequence_(Literal("+", nil), Return(addFunc))
	minus := Sequence_(Literal("-", nil), Return(minusFunc))
	op := Or(Try(add), Try(minus))

	expr := ChainLeft1(number, op)

	result = expr.parse("1+2")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, 3)

	result = expr.parse("1+2-2")
	c.Check(result.Success, Equals, true)
	c.Check(result.Value, Equals, 1)
}

func (s *S) TestSepBy(c *C) {
	one := Literal("1", 1)
	two := Literal("2", 2)
	three := Literal("3", 3)
	four := Literal("4", 4)
	number := Or(Try(Or(Try(one), Try(two))), Or(Try(three), Try(four)))

	listOfNumbers := SeperatedBy(number, Literal(",", nil))

	result := listOfNumbers.parse("1,2,3,4")
	c.Check(result.Success, Equals, true)
	slice, ok := result.Value.([]interface{})
	c.Check(ok, Equals, true)
	c.Check(slice, Equals, []interface{}{1, 2, 3, 4})
}

func (s *S) TestNiceErrors(c *C) {
	one := Tag(Literal("1", 1), "one")
	two := Literal("2", 2)
	three := Literal("3", 3)
	four := Literal("4", 4)
	number := Or(Try(Or(Try(one), Try(two))), Or(Try(three), Try(four)))

	listOfNumbers := SeperatedBy(number, Literal(",", nil))

	result := listOfNumbers.parse("1,a,3,4")
	c.Check(result.Success, Equals, false)
	fmt.Println(result)
}
