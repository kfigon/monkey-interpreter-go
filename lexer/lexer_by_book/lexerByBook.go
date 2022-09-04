package lexerbybook


type llex struct {
	input string
	idx int
}

func newLlexer(input string) *llex {
	return &llex{input: input}
}

type llexToken struct {
	token string
	lexeme string
}

func (l *llex) readChar() rune {
	toRet := l.input[l.idx]
	l.idx++
	return rune(toRet)
}

func (l *llex) peekChar() rune {
	return rune(l.input[l.idx])
}

func (l *llex) eof() bool {
	return l.idx >= len(l.input)
}

func (l *llex) emit() llexToken {
	if l.eof() {
		return llexToken{token: "EOF", lexeme: ""}
	}
	char := l.readChar()
	toPrint := string(char)
	if toPrint == "" {
		char-=1
		char+=1
	}
	switch char {

	case ';': return llexToken{"SEMICOLON", string(char)}
	case '+': {
		if !l.eof() && l.peekChar() == '+' {
			l.readChar()
			return llexToken{"OPERATOR", "++"}
		}
		return llexToken{"OPERATOR", string(char)}
	}
	case '<': return llexToken{"OPERATOR", string(char)}
	case '>': return llexToken{"OPERATOR", string(char)}
	case '-': return llexToken{"OPERATOR", string(char)}
	case '*': return llexToken{"OPERATOR", string(char)}

	case '(': return llexToken{"OPEN_BRACE", string(char)}
	case ')': return llexToken{"CLOSE_BRACE", string(char)}
	case '{': return llexToken{"OPEN_CBRACE", string(char)}
	case '}': return llexToken{"CLOSE_CBRACE", string(char)}

	case '=': {
		if !l.eof() && l.peekChar() == '=' {
			l.readChar()
			return llexToken{"EQUALS", "=="}
		}
		return llexToken{"ASSIGNMENT", "="}
	}
	case '!': {
		if !l.eof() && l.peekChar() == '=' {
			l.readChar()
			return llexToken{"NOT_EQUALS", "!="}
		}
		return llexToken{"NEGATION", "!"}
	}
	default: {
		 if isAlpha(char) {
			consumedString := string(char)+l.eatIdentifier()
			if isKeyword(consumedString) {
				return llexToken{"KEYWORD", consumedString}
			}
			return llexToken{"IDENTIFIER", consumedString}
		 } else if isDigit(char) {
			consumedString := string(char)+l.eatNumber()
			return llexToken{"NUMBER", consumedString}
		 } else if isWhiteSpace(char) {
			 return l.emit()
		 }
	}
	}
	return llexToken{"UNKNOWN_TOKEN", string(char)}
}

func isWhiteSpace(char rune) bool {
	return char ==' ' || char == '\t' || char == '\n' || char == '\r'
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func (l *llex) eatIdentifier() string {
	return l.consumeCharacter(isAlpha)()
}

func (l *llex) eatNumber() string {
	return l.consumeCharacter(isDigit)()
}

func (l *llex) consumeCharacter(fn func(rune)bool) func() string {
	return func() string {
		var out string
		for !l.eof() {
			c := l.peekChar()
			if !fn(c) {
				break
			}
			l.readChar()
			out += string(c)
		}
		return out
	}
}

func isKeyword(input string) bool {
	keywords := map[string]bool{
		"let": true,
		"func": true,
		"for": true,
	}
	return keywords[input]
}
