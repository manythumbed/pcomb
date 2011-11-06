package pcomb

import(
	"utf8"
)

type ParseResult struct	{
	Success bool
	Result interface{}
	Remaining string
}

type Parser func(input string) ParseResult

var Fail Parser = func(input string) ParseResult	{
	return ParseResult{false, nil, input}
}

func Succeed(value interface{}) Parser {
	return func(input string) ParseResult	{
		return ParseResult{true, value, input}
	}
}

type Operation func(value interface{}) Parser

func Then(a Parser, f Operation) Parser {
	return func(input string) ParseResult {
		result := a(input)
		if result.Success {
			return f(result.Result)(result.Remaining)
		}
		return result
	}
}

func Then_(a, b Parser) Parser {
	return Then(a, func(interface{}) Parser { return b })
}

func Or(a, b Parser) Parser {
	return func(input string) ParseResult {
		result := a(input)
		if result.Success {
			return result
		}
		return b(input)
	}
}

func Item() Parser {
	return func (input string) ParseResult	{
		str := utf8.NewString(input)
		if str.RuneCount() > 0 {
			return ParseResult{true, str.Slice(0, 1), str.Slice(1, str.RuneCount())}
		}
		return ParseResult{false, nil, input}
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
