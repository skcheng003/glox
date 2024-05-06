package scan

import (
	"glox/src/token"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func (scanner *Scanner) isNotEnd() bool {
	return scanner.current < len(scanner.source)
}

func (scanner *Scanner) ScanTokens() []token.Token {
	for scanner.isNotEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}
	scanner.tokens = append(scanner.tokens, token.NewToken(token.EOF, "", "", scanner.line))
	return scanner.tokens
}

func (scanner *Scanner) scanToken() {
	b := scanner.advance()
	switch b {
	case '(':
		scanner.addToken(token.LEFT_PAREN)
	case ')':
		scanner.addToken(token.RIGHT_PAREN)
	case '{':
		scanner.addToken(token.LEFT_BRACE)
	case '}':
		scanner.addToken(token.RIGHT_BRACE)
	case ',':
		scanner.addToken(token.COMMA)
	case '.':
		scanner.addToken(token.DOT)
	case '-':
		scanner.addToken(token.MINUS)
	case '+':
		scanner.addToken(token.PLUS)
	case ';':
		scanner.addToken(token.SEMICOLON)
	case '*':
		scanner.addToken(token.STAR)
	case '!':
		scanner.addToken(scanner.selectTokenType(scanner.match('='), token.BANG_EQUAL, token.BANG))
	case '=':
		scanner.addToken(scanner.selectTokenType(scanner.match('='), token.BANG_EQUAL, token.EQUAL))
	case '>':
		scanner.addToken(scanner.selectTokenType(scanner.match('='), token.GREATER_EQUAL, token.GREATER))
	case '<':
		scanner.addToken(scanner.selectTokenType(scanner.match('='), token.LESS_EQUAL, token.LESS))
	case '/':
		if scanner.match('/') {
			if scanner.peek() != '\n' && scanner.isNotEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		scanner.line += 1
	case '"':
		scanner.string()
	default:
		if scanner.isDigit(b) {
			scanner.number()
		} else if scanner.isAlpha(b) {
			scanner.identifier()
		} else {
			// TODO: add some error handle code, lox.Error
		}
	}
}

func (scanner *Scanner) advance() byte {
	b := scanner.source[scanner.current]
	scanner.current += 1
	return b
}

func (scanner *Scanner) addToken(tokenType token.TokenType) {
	scanner.addTokenWithLiteral(tokenType, "")
}

func (scanner *Scanner) addTokenWithLiteral(tokenType token.TokenType, literal any) {
	lexeme := scanner.source[scanner.start:scanner.current]
	newToken := token.NewToken(tokenType, lexeme, literal, scanner.line)
	scanner.tokens = append(scanner.tokens, newToken)
}

func (scanner *Scanner) match(expected byte) bool {
	if !scanner.isNotEnd() {
		return false
	}
	if scanner.source[scanner.current] != expected {
		return false
	}
	scanner.current += 1
	return true
}

func (scanner *Scanner) selectTokenType(flag bool, tokenType1, tokenType2 token.TokenType) token.TokenType {
	if flag {
		return tokenType1
	} else {
		return tokenType2
	}
}

func (scanner *Scanner) peek() byte {
	if !scanner.isNotEnd() {
		return byte(0)
	}
	return scanner.source[scanner.current]
}

func (scanner *Scanner) peekNext() byte {
	if scanner.current+1 >= len(scanner.source) {
		return byte(0)
	}
	return scanner.source[scanner.current+1]
}

func (scanner *Scanner) isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (scanner *Scanner) isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

func (scanner *Scanner) isAlphaNumeric(b byte) bool {
	return scanner.isAlpha(b) || scanner.isDigit(b)
}

func (scanner *Scanner) string() {
	for scanner.peek() != '"' && scanner.isNotEnd() {
		if scanner.peek() == '\n' {
			scanner.line += 1
		}
		scanner.advance()
	}
	if !scanner.isNotEnd() {
		// TODO: error message, "Unterminated string."
	}
	scanner.advance()
	literal := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addTokenWithLiteral(token.STRING, literal)
}

func (scanner *Scanner) number() {
	for scanner.isDigit(scanner.peek()) {
		scanner.advance()
	}
	if scanner.peek() == '.' && scanner.isDigit(scanner.peekNext()) {
		scanner.advance()
		for scanner.isDigit(scanner.peek()) {
			scanner.advance()
		}
	}
	number, _ := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)
	scanner.addTokenWithLiteral(token.NUMBER, number)
}

func (scanner *Scanner) identifier() {
	for scanner.isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}
	key := scanner.source[scanner.start:scanner.current]
	tokenType := token.IDENTIFIER
	if val, ok := keywords[key]; ok {
		tokenType = val
	}
	scanner.addToken(tokenType)
}
