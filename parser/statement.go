package parser

// Statement holds information about the current statement that is being parsed
type Statement struct {
	Index  int
	Value  string
	Length int
}

// GetStatement ...
func (p *Program) GetStatement() {
	p.GetExpr()
	p.GetEOS()
}
