package pcomb

import(
	"utf8"
)

type ParseResult struct	{
	Result interface{}
	Remaining string
}

type Parser func(input string) (success bool, result ParseResult)

var Fail Parser = func(input string) (bool, ParseResult)	{
	return false, ParseResult{nil, input}
}

func Succeed(value interface{}) Parser {
	return func(input string) (bool, ParseResult)	{
		return true, ParseResult{value, input}
	}
}

type Operation func(value interface{}) Parser

func Then(a Parser, f Operation) Parser {
	return func(input string) (bool, ParseResult) {
		success, result := a(input)
		if success {
			return f(result.Result)(result.Remaining)
		}
		return success, result
	}
}

func Then_(a, b Parser) Parser {
	return Then(a, func(interface{}) Parser { return b })
}

func Or(a, b Parser) Parser {
	return func(input string) (bool, ParseResult) {
		success, result := a(input)
		if success {
			return success, result
		}
		return b(input)
	}
}

func Item() Parser {
	return func (input string) (bool, ParseResult)	{
		str := utf8.NewString(input)
		if str.RuneCount() > 0 {
			return true, ParseResult{str.Slice(0, 1), str.Slice(1, str.RuneCount())}
		}
		return false, ParseResult{nil, input}
	}
}

func Satisfy(predicate func(rune int) bool) Parser {
	op := func (value interface{}) Parser {
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

func charEq(str string) func(rune int) bool	{
	char, _ := utf8.DecodeRuneInString(str)
	return func (rune int) bool {
		return rune == char
	}
}

func Literal(str string, result interface{}) Parser {
	if len(str) > 0 {
		return Then_(Satisfy(charEq(str)), Literal(str[1:], result))
	}
	return Succeed(result)
}
