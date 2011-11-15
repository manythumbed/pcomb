package pcomb

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
	return p(State{newReader(strings.NewReader(text)})
}
