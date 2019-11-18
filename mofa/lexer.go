package mofa

type TokenType int

const (
	TTIDEN TokenType = iota
	TTLITERAL
	TTNUM
	TTLBRK
	TTRBRK
	TTEOF
)

func isAlpha(c rune) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	switch c {
	case '+', '-', '*', '/', '>', '<', '!', '?', '=':
		return true
	}
	return false
}

func isNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAlphanum(c rune) bool {
	return isAlpha(c) || isNum(c)
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\r' || c == '\n' || c == '\t'
}

type Lexer struct {
	token     string
	tokenType TokenType
	cc        rune
	pc        int
	source    []rune
	sourceLen int
}

func NewLexer(source string) *Lexer {
	lex := &Lexer{}
	lex.source = []rune(source)
	lex.sourceLen = len(lex.source)
	lex.cc = 0
	lex.pc = 0
	lex.nextChar()
	return lex
}

func (lex *Lexer) nextChar() {
	if lex.pc < lex.sourceLen {
		lex.cc = lex.source[lex.pc]
		lex.pc++
	} else {
		lex.cc = 0
	}
}

func (lex *Lexer) NextToken() {
	for isSpace(lex.cc) {
		lex.nextChar()
	}

	if isAlpha(lex.cc) {
		var token string
		token = string(lex.cc)
		lex.nextChar()
		for isAlphanum(lex.cc) {
			token += string(lex.cc)
			lex.nextChar()

		}
		lex.token = token
		lex.tokenType = TTIDEN
	} else if isNum(lex.cc) {
		var token string
		token = string(lex.cc)
		lex.nextChar()

		for isNum(lex.cc) {
			token += string(lex.cc)
			lex.nextChar()
		}
		lex.token = token
		lex.tokenType = TTNUM
	} else if lex.cc == 0 {
		// eof
		lex.token = ""
		lex.tokenType = TTEOF

	} else {
		switch lex.cc {
		case '(':
			lex.token = "("
			lex.tokenType = TTLBRK
			lex.nextChar()
		case ')':
			lex.token = ")"
			lex.tokenType = TTRBRK
			lex.nextChar()
		default:
			panic("unknown char: " + string(lex.cc))
		}
	}
}
