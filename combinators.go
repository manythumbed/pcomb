package pcomb

import (
	"strings"
)

type State struct {
	reader reader
}

type Output struct	{
	Succeeded bool
	Value interface{}
	Next State
}

type Parser func(state State) Output

func (p Parser) parse(text string) Output	{
	return p( State{ newReader(strings.NewReader(text))})
}

func Fail() Parser {
	return func(state State) Output {
		return Output{false, nil, state}
	}
}

func Return(value interface{}) Parser	{
	return func(state State) Output {
		return Output{true, value, state}
	}
}

func Satisfy(predicate func(int) bool) Parser {
	return func(state State) Output {
		rune, err := state.reader.take()
		if err != nil {
			return Output{false, err, state}
		}
		if predicate(rune) {
			return Output{true, rune, state}
		}
		_ = state.reader.untake()
		return Output{false, nil,  state}
	}
}

func Item() Parser {
	return Satisfy(func(rune int) bool { return true })
}
