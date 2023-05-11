package ast

const (
	IDENT      = "IDENT"
	INST_IDENT = "INST_IDENT"

	EOF      = "EOF"
	NUMBER   = "NUMBER"
	STRING   = "STRING"
	QUESTION = "QUESTION" // ?
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"  // !
	ASTER    = "ASTER" // *
	SLASH    = "SLASH" // /
	LT       = "LT"    // <
	GT       = "GT"    // >
	LE       = "LE"    // <=
	GE       = "GE"    // >=
	EQ       = "EQ"    // ==
	NE       = "NE"    // !=
	ASSIGN   = "ASSIGN"
	COMMA    = "COMMA" // ,
	DOT      = "DOT"   // .
	SEMI     = "SEMI"  // ;
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"
	AT       = "AT"

	NEWLINE    = "NEWLINE"
	WHITESPACE = "WHITESPACE"
	TAB        = "TAB"
	CR         = "CR"

	// Keywords
)

type Type string

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '\n':
		tok = newToken(NEWLINE, 'n')
	case '\t':
		tok = newToken(TAB, l.ch)
	case '\r':
		tok = newToken(CR, l.ch)
	case ' ':
		tok = newToken(WHITESPACE, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '*':
		tok = newToken(ASTER, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '[':
		tok = newToken(LBRACKET, l.ch)
	case ']':
		tok = newToken(RBRACKET, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case ';':
		tok = newToken(SEMI, l.ch)
	case '?':
		tok = newToken(QUESTION, l.ch)
	case '.':
		tok = newToken(DOT, l.ch)
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: GE, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(GT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LE, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(LT, l.ch)
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NE, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '@':
		if isLetter(l.peekChar()) {
			l.readChar()
			tok.Value = l.readIdentifier()
			tok.Type = INST_IDENT
			return tok
		} else {
			tok = newToken(AT, l.ch)
		}
	case 0:
		tok.Value = ""
		tok.Type = EOF
	case '"':
		tok.Type = STRING
		tok.Value = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = IDENT
			return tok
		} else if isDigit(l.ch) {
			tok.Value = l.readNumber()
			tok.Type = NUMBER
			return tok
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func newToken(tokenType string, ch byte) Token {
	return Token{Type: tokenType, Value: string(ch)}
}
