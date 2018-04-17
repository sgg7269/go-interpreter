package parser

// GetVar ...
func (p *Program) GetVar() {
	p.GetType()
	p.GetSeparator()
	p.GetIdent()
}
