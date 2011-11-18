package pcomb

type Error struct {
	Position
	Message string
	Errors  []Error
}

func NoErrors() []Error {
	return []Error{}
}

func NewError(p Position, m string) Error {
	return Error{p, m, NoErrors()}
}

func (e *Error) Append(errors ...Error) {
	e.Errors = append(e.Errors, errors...)
}
