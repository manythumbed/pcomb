package pcomb

/*
import (
	"utf8"
)

type Position struct {
	Column, Line int
}

type State struct {
	Position
	Input string
}

type Error struct {
	Position
	Message string
}

type Result struct {
	Success   bool
	Consumed bool
	Error Error
	Value    interface{}
	Remaining State
}

type Parser func(input State) Result

func (p Parser) parse(text string) Result	{
	return p(State{Input:text})
}

var Fail Parser = func(input State) Result {
	return Result{false, false, Error{}, nil, input}
}

func Succeed(value interface{}) Parser {
	return func(input State) Result {
		return Result{true, false, Error{}, value, input}
	}
}

type Operation func(value interface{}) Parser

func Then(a Parser, f Operation) Parser {
	return func(input State) Result {
		result := a(input)
		if result.Success {
			return f(result.Value)(result.Remaining)
		}
		return result
	}
}

func (p Parser) Then(f Operation) Parser {
	return Then(p, f)
}

func Then_(a, b Parser) Parser {
	return Then(a, func(interface{}) Parser { return b })
}

func (a Parser) Then_(b Parser) Parser {
	return Then_(a, b)
}

func Or(a, b Parser) Parser {
	return func(input State) Result {
		result := a(input)
		if result.Success {
			return result
		}
		return b(input)
	}
}

func Item() Parser {
	return func(input State) Result {
		str := utf8.NewString(input.Input)
		if str.RuneCount() > 0 {
			return Result{true, true, Error{}, str.Slice(0, 1), State{Input:str.Slice(1, str.RuneCount())}}
		}
		return Result{false, false, Error{}, nil, input}
	}
}

func Satisfy(predicate func(rune int) bool) Parser {
	op := func(value interface{}) Parser {
		str, ok := value.(string)
		if ok {
			rune, _ := utf8.DecodeRuneInString(str)
			if predicate(rune) {
				return Succeed(str)
			}
		}
		return Fail
	}
	return Then(Item(), op)
}

func charEq(str string) func(rune int) bool {
	char, _ := utf8.DecodeRuneInString(str)
	return func(rune int) bool {
		return rune == char
	}
}

func Literal(str string, result interface{}) Parser {
	if len(str) > 0 {
		return Then_(Satisfy(charEq(str)), Literal(str[1:], result))
	}
	return Succeed(result)
}

func emptySlice() []interface{} {
	return make([]interface{}, 0)
}

func Many(p Parser) Parser {
	return Or(Many1(p), Succeed(emptySlice()))
}

func cons(x interface{}, xs []interface{}) []interface{} {
	if x != nil {
		return append([]interface{}{x}, xs...)
	}
	return xs
}

func Many1(p Parser) Parser {
	op := func(x interface{}) Parser {
		consOp := func(xs interface{}) Parser {
			slice, _ := xs.([]interface{})
			return Succeed(cons(x, slice))
		}
		return Then(Many(p), consOp)
	}
	return Then(p, op)
}

func chain(x interface{}, p, op Parser) Parser {
	f_func := func(fval interface{}) Parser {
		f, _ := fval.(func(a, b interface{}) interface{})
		y_func := func(y interface{}) Parser {
			return chain(f(x, y), p, op)
		}
		return Then(p, y_func)
	}
	return Or(Then(op, f_func), Succeed(x))
}

func ChainLeft1(p Parser, op Parser) Parser {
	remainder := func(x interface{}) Parser {
		return chain(x, p, op)
	}
	return Then(p, remainder)
}

func SeperatedBy(p, sep Parser) Parser {
	return Or(SeperatedBy1(p, sep), Succeed(emptySlice()))
}

func SeperatedBy1(p, sep Parser) Parser {
	return Then(p, func(x interface{}) Parser {
		return Then(seperated(p, sep), func(xs interface{}) Parser {
			slice, _ := xs.([]interface{})
			return Succeed(cons(x, slice))
		})
	})
}

func seperated(p, sep Parser) Parser {
	return Or(Then_(sep, SeperatedBy1(p, sep)), Succeed(emptySlice()))
}

func Try(p Parser) Parser {
	return p
}

*/

/* TODO
Refactor to use Reader
Refactor to provide error handling
*/
