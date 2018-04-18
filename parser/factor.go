package parser

// GetFactor ...
func (p *Program) GetFactor() {
	// ( expr )
	// < literal >
	// var

	p.GetVar()
}
