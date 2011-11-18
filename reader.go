package pcomb

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const newline = int('\n')

type position struct {
	line, rune, column int
}

func (p position) String() string	{
	return fmt.Sprintf("Line %d Rune %d Column %d", p.line, p.rune, p.column)
}

func (p *position) forward(b int, incLine bool) *position {
	if b > 0 {
		if incLine {
			return &position{p.line + 1, p.rune + 1, p.column + b}
		}
		return &position{p.line, p.rune + 1, p.column + b}
	}
	return p
}

type reader struct {
	*bufio.Reader
	current, previous *position
}

func newReader(r io.Reader) reader {
	return reader{bufio.NewReader(r), &position{1,0,0}, &position{1,0,0}}
}

func (r reader) String() string	{
	return fmt.Sprintf("Current[%v], Previous[%v]", r.current, r.previous)
}

func (r *reader) take() (int, os.Error) {
	rune, size, err := r.ReadRune()
	if err == nil {
		r.previous = r.current
		r.current = r.current.forward(size, rune == newline)
	}
	return  rune, err
}

func (r *reader) untake() os.Error	{
	err := r.UnreadRune()
	if err == nil	{
		r.current = r.previous
	}
	return err
}

func (r reader) taken() bool {
	return r.current.line > r.previous.line || r.current.rune > r.previous.rune
}
