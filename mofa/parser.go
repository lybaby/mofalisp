package mofa

type Parser struct {
	lexer *Lexer
}

func NewParser(source string) *Parser {
	p := &Parser{}

	p.lexer = NewLexer(source)
	p.lexer.NextToken()
	return p
}

func (p *Parser) Parse() {
	for p.lexer.tokenType != TTEOF {
		p.ParseForm()
	}
}


func (p *Parser) ParseForm() {
	switch p.lexer.tokenType {
	case TTIDEN:
		p.lexer.NextToken()
	case TTNUM:
		p.lexer.NextToken()
	case TTLBRK:
		p.ParseSExpr()
	default:
		panic("unknown token:" + p.lexer.token)
	}
}
func (p *Parser) ParseSExpr() {
	if p.lexer.tokenType == TTLBRK {
		p.lexer.NextToken()
		for p.lexer.tokenType != TTEOF && p.lexer.tokenType != TTRBRK {
			
			p.ParseForm()
		}

		if p.lexer.tokenType == TTRBRK {
			p.lexer.NextToken()
		} else {
			panic("require )")
		}

	} else {
		// error
		panic("require (")
	}
}
