package pcomb

import (
	"strings"
	"fmt"
)

type State struct {
	reader reader
}

type Error struct {
	Position Position
	Message string
}

var NoErrors = Error{}

type Output struct	{
	Success bool
	Error Error
	Value interface{}
	Next State
}

type Parser func(state State) Output

func (p Parser) parse(text string) Output	{
	return p( State{ newReader(strings.NewReader(text))})
}

func Fail() Parser {
	return func(state State) Output {
		return Output{false, NoErrors, nil, state}
	}
}

func Return(value interface{}) Parser	{
	return func(state State) Output {
		return Output{true, NoErrors, value, state}
	}
}

func (s State) error(message string) Error {
	return Error{*s.reader.current, message}
}

func Satisfy(predicate func(int) bool) Parser {
	return func(state State) Output {
		rune, err := state.reader.take()
		if err != nil {
			return Output{false, state.error("Error reading data"), err, state}
		}
		if predicate(rune) {
			return Output{true, Error{}, rune, state}
		}
		output := Output{false, state.error(fmt.Sprintf("Character %c does not match predicate", rune)), nil,  state}
		_ = state.reader.untake()
		return output
	}
}

func Item() Parser {
	return Satisfy(func(rune int) bool { return true })
}
