package pcomb

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type position struct {
	line, rune, column int
}

type reader struct {
	*bufio.Reader
	current, previous *position
}

func newReader(r io.Reader) reader {
	return reader{bufio.NewReader(r), &position{}, &position{}}
}

func (r reader) String() string	{
	return fmt.Sprintf("Current %v - Previous %v", r.current, r.previous)
}

func (r *reader) readRune() (int, os.Error) {
	rune, size, err := r.ReadRune()
	if err == nil {
		r.previous = r.current
		r.current = &position{r.current.line, r.current.rune + 1, r.current.column + size}
	}
	return  rune, err
}
