package token

type (
	// Value encapsulates the value of the token
	Value struct {
		String string
		True   interface{}
	}

	// Token is a syntactical structure used to hold representative metadata
	Token struct {
		ID       int
		Location [2]int
		Type     string
		Value    Value
	}
)

// SetLocation ...
func (t *Token) SetLocation(start, end int) {
	t.SetStart(start)
	t.SetEnd(end)
}

// SetStart ...
func (t *Token) SetStart(index int) {
	t.Location[0] = index
}

// SetEnd ...
func (t *Token) SetEnd(index int) {
	t.Location[1] = index
}

// NewToken ...
func NewToken(index int) Token {
	return Token{Location: [2]int{index, -1}}
}
