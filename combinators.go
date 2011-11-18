package pcomb

import (
	"fmt"
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

/*
func combineErrors(a, b Output) []Error {
	return append(a.Errors, b.Errors...)
}

func Or(a, b Parser) Parser {
	return func(state State) Output {
		outputA := a(state)
		if outputA.Success && outputA.taken() {
			return outputA
		}
		outputB := b(state)
		if outputB.taken() {
			return outputB
		}
		return Output{false, combineErrors(outputA, outputB), nil, state}
	}
}

func Try(p Parser) Parser {
	return func(state State) Output	{
		output := p(state)
		if output.taken() && !output.Success	{
			return Output{false, output.Errors, output.Value, output.Next}
		}
		return output
	}
}
*/
