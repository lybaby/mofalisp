package mofa

import (
	"strconv"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(source string) *Parser {
	p := &Parser{}

	p.lexer = NewLexer(source)
	p.lexer.NextToken()
	return p
}

func (p *Parser) Parse() *AstNode {

	return p.ParseForm()

}

func (p *Parser) ParseForm() *AstNode {
	switch p.lexer.tokenType {
	case TTIDEN:
		node := &AstNode{nodeType: NTSYM, sval: p.lexer.token}
		p.lexer.NextToken()
		return node
	case TTNUM:
		num, err := strconv.ParseFloat(p.lexer.token, 64)
		if err != nil {
			panic(err.Error())
		}
		node := &AstNode{nodeType: NTNUM, nval: num}
		p.lexer.NextToken()
		return node
	case TTLBRK:
		return p.ParseSExpr()
	default:
		panic("unknown token:" + p.lexer.token)
	}
}
func (p *Parser) ParseSExpr() *AstNode {
	if p.lexer.tokenType == TTLBRK {
		p.lexer.NextToken()

		nodes := make([]*AstNode, 0)

		for p.lexer.tokenType != TTEOF && p.lexer.tokenType != TTRBRK {
			node := p.ParseForm()
			nodes = append(nodes, node)
		}

		cdrNode := &AstNode{nodeType: NTSEXPR}

		for x := len(nodes) - 1; x >= 0; x-- {
			cdrNode = &AstNode{nodeType: NTSEXPR, car: nodes[x], cdr: cdrNode}
		}

		if p.lexer.tokenType == TTRBRK {
			p.lexer.NextToken()
			return cdrNode
		} else {
			panic("require )")
		}

	} else {
		// error
		panic("require (")
	}
}
