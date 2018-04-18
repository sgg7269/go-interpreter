package parser

// GetVar ...
func (p *Program) GetVar() {
	p.GetType()
	eos, err := p.GetSeparator()
	if err != nil {

	}
	if eos {

	}
	p.GetIdent()
}
