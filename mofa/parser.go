package mofa

type Parser struct {
	lexer *Lexer
}

func NewParser (source string) *Parser{
	p:=&Parser{}

	p.lexer = NewLexer(source)

	return p
}

func (p *Parser)Parse() {

}