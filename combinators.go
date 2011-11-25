package pcomb

import (
	"fmt"
	"utf8"
)

type Output struct {
	Success bool
	Errors  []Error
	Value   interface{}
	State
}

func success(value interface{}, state State) Output {
	return Output{true, NoErrors(), value, state}
}

func failure(message string, state State) Output {
	return Output{false, []Error{NewError(state.Position, message)}, nil, state}
}

func multipleFailure(message string, state State, outputs ...Output) Output {
	errors := []Error{}
	for _, o := range outputs {
		for _, e := range o.Errors {
			errors = append(errors, e)
		}
	}

	error := NewError(state.Position, message)
	error.Errors = errors
	return Output{false, []Error{error}, nil, state}
}

type Parser func(state State) Output

func (p Parser) parse(text string) Output {
	return p(newState(text))
}

var Fail Parser = func(state State) Output {
	return Output{false, NoErrors(), nil, state}
}

func Return(value interface{}) Parser {
	return func(state State) Output {
		return Output{true, NoErrors(), value, state}
	}
}

func Satisfy(predicate func(int) bool) Parser {
	return func(state State) Output {
		rune, next, ok := state.Next()
		if !ok {
			return failure("End of data", state)
		}
		if predicate(rune) {
			return success(rune, next)
		}
		return failure(fmt.Sprintf("Character (%c) does not match predicate", rune), next)
	}
}

func Item() Parser {
	return Satisfy(func(rune int) bool { return true })
}

func Or(a, b Parser) Parser {
	return func(state State) Output {
		outA := a(state)
		if outA.Success || !state.Equals(outA.State) {
			return outA
		}
		outB := b(state)
		if outB.Success || !state.Equals(outB.State) {
			return outB
		}
		return multipleFailure("Both sides of Or expression failed", state, outA, outB)
	}
}

func Try(p Parser) Parser {
	return func(state State) Output	{
		output := p(state)
		if !state.Equals(output.State) && !output.Success	{
			return Output{false, output.Errors, output.Value, state}
		}
		return output
	}
}

const labelFormat = "Attempted to consume %s and failed"

func Tag(p Parser, label string) Parser {
	return func(state State) Output	{
		out := p(state)
		if !state.Equals(out.State)	{
			return out
		}
		return failure(fmt.Sprintf(labelFormat, label), state)
	}
}

func Sequence(p Parser, f func(value interface{}) Parser) Parser {
	return func(state State) Output	{
		out := p(state)
		if out.Success {
			return f(out.Value)(out.State)
		}
		return out
	}
}

func Sequence_(a, b Parser) Parser {
	return Sequence(a, func(ignore interface{}) Parser	{
		return b
	})
}

func literalRune(rune int) Parser {
	return Satisfy(func(r int) bool { return rune == r })
}

func Literal(s string, result interface{}) Parser	{
	match := utf8.NewString(s)
	if len(s) == 0 {
		return Return(result)
	}
	return Sequence_(literalRune(match.At(0)), Literal(match.Slice(1, match.RuneCount()), result))
}

func cons(x interface{}, xs []interface{}) []interface{} {
	if x != nil {
		return append([]interface{}{x}, xs...)
	}
	return xs
}

func emptySlice() []interface{} {
	return make([]interface{}, 0)
}

func ZeroOrMore(p Parser) Parser {
	return Or(OneOrMore(p), Return(emptySlice()))
}

func OneOrMore(p Parser) Parser {
	op := func(x interface{}) Parser {
		consOp := func(xs interface{}) Parser {
			slice, _ := xs.([]interface{})
			return Return(cons(x, slice))
		}
		return Sequence(ZeroOrMore(p), consOp)
	}
	return Sequence(p, op)
}

