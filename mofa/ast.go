package mofa

type NodeType int

const (
	NTSEXPR NodeType = iota
	NTNUM
	NTLITERAL
	NTSYM
)

type AstNode struct {
	nodeType NodeType
	sval string
	nval float64
	car *AstNode
	cdr *AstNode
}



