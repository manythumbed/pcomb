package pcomb

import (
	"utf8"
)

type Position struct {
	Line, Rune, Column int
}

type State struct {
	text string
	Position
}

func (s State) increment(rune, size int) Position {
	if rune == int('\n') {
		return Position{s.Line + 1, 0, 0}
	}
	return Position{s.Line, s.Rune + 1, s.Column + size}
}

func (s State) Equals(that State) bool {
	return s.text == that.text
}

func newState(text string) State {
	return State{text, Position{1, 0, 0}}
}

func (s State) Next() (rune int, state State, ok bool) {
	if len(s.text) > 0 {
		rune, size := utf8.DecodeRuneInString(s.text)
		return rune, State{s.text[size:], s.increment(rune, size)}, true
	}

	return 0, s, false
}
