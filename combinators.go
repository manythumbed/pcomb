package pcomb

/*
import (
	"strings"
	"fmt"
)

type State struct {
	reader reader
}

type Error struct {
	Position Position
	Message  string
}

var NoErrors = make([]Error, 0)

type Output struct {
	Success bool
	Errors  []Error
	Value   interface{}
	Next    State
}

func (o Output) taken() bool {
	return o.Next.reader.taken()
}

type Parser func(state State) Output

func (p Parser) parse(text string) Output {
	return p(State{newReader(strings.NewReader(text))})
}

func Fail() Parser {
	return func(state State) Output {
		return Output{false, NoErrors, nil, state}
	}
}

func Return(value interface{}) Parser {
	return func(state State) Output {
		return Output{true, NoErrors, value, state}
	}
}

func (s State) error(message string) []Error {
	return []Error{Error{*s.reader.current, message}}
}

func Satisfy(predicate func(int) bool) Parser {
	return func(state State) Output {
		rune, err := state.reader.take()
		if err != nil {
			return Output{false, state.error(fmt.Sprintf("Error reading data %v", err)), err, state}
		}
		if predicate(rune) {
			return Output{true, NoErrors, rune, state}
		}
		output := Output{false, state.error(fmt.Sprintf("Character (%c) does not match predicate", rune)), nil, state}
		_ = state.reader.untake()
		return output
	}
}

func Item() Parser {
	return Satisfy(func(rune int) bool { return true })
}

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
