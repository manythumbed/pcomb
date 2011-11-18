package pcomb

type Error struct {
	Position
	Message string
	Errors []Error
}

func NewError(p Position, m string) Error	{
	return Error{p, m, make([]Error, 0)}
}

func (e *Error) Append(errors ...Error)	{
	e.Errors = append(e.Errors, errors...)
}
