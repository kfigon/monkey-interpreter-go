package lexer

import (
	"fmt"
	"log"
	"regexp"
)

var enableLogs bool = false
var skipWhitespaces bool = true

type Token struct {
	Class  TokenClass
	Lexeme string
}

func (t Token) String() string {
	return fmt.Sprintf("{%v %q}", t.Class, t.Lexeme)
}

type TokenClass int

func (t TokenClass) String() string {
	return classesStrings[t]
}

const (
	Whitespace TokenClass = iota
	Keyword
	Identifier // variables
	Number
	Boolean
	Operator

	OpenParam
	CloseParam
	Semicolon
	Assignment
	EOF
)

var classesStrings = []string{
	"Whitespace",
	"Keyword",
	"Identifier",
	"Number",
	"Boolean",
	"Operator",
	"OpenParam",
	"CloseParam",
	"Semicolon",
	"Assignment",
	"EOF",
}

type tokenizerEntry struct {
	reg   *regexp.Regexp
	class TokenClass
}

var tokenizerEntries []tokenizerEntry = []tokenizerEntry{
	{regexp.MustCompile(`^(\s+)`), Whitespace},

	{regexp.MustCompile(`^(if)($|\s|\()`), Keyword},
	{regexp.MustCompile(`^(else)($|\s|\()`), Keyword},
	{regexp.MustCompile(`^(for)($|\s|\()`), Keyword},
	{regexp.MustCompile(`^(var)($|\s|\()`), Keyword},
	{regexp.MustCompile(`^(return)($|\s|\()`), Keyword},
	{regexp.MustCompile(`^(fn)($|\s|\()`), Keyword},

	{regexp.MustCompile(`^(==)($|\s?)`), Operator},
	{regexp.MustCompile(`^(!=)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\+\+)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\+)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\-\-)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\-)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\*)($|\s?)`), Operator},
	{regexp.MustCompile(`^(\/)($|\s?)`), Operator},
	{regexp.MustCompile(`^(<=)($|\s?)`), Operator},
	{regexp.MustCompile(`^(>=)($|\s?)`), Operator},
	{regexp.MustCompile(`^(<)($|\s?)`), Operator},
	{regexp.MustCompile(`^(>)($|\s?)`), Operator},
	{regexp.MustCompile(`^(!)($|\s?)`), Operator},

	{regexp.MustCompile(`^(=)($|\s?)`), Assignment},

	{regexp.MustCompile(`^(;)`), Semicolon},
	{regexp.MustCompile(`^(\))`), CloseParam},
	{regexp.MustCompile(`^(\()`), OpenParam},
	{regexp.MustCompile(`^({)`), OpenParam},
	{regexp.MustCompile(`^(})`), CloseParam},

	{regexp.MustCompile(`^(true)($|\s|;|,\))`), Boolean},
	{regexp.MustCompile(`^(false)($|\s|;|,\))`), Boolean},
	{regexp.MustCompile(`^([0-9]+\.[0-9]+)`), Number},
	{regexp.MustCompile(`^([0-9]+)`), Number},

	{regexp.MustCompile(`^(\w+)`), Identifier},
}

func Tokenize(input string) []Token {
	var tokens []Token
	var idx uint64

	ln := uint64(len(input))
	for idx < ln {
		rest := input[idx:]

		logLine(idx, ln, rest)

		found, deltaIdx, token := processAvailableTokens(rest)

		if !found {
			log.Println("Unknown token at idx", idx)
			idx++
			continue
		}

		idx += uint64(deltaIdx)
		if skipWhitespaces && token.Class == Whitespace {
			continue
		}
		tokens = append(tokens, token)
	}
	tokens = append(tokens, Token{Class: EOF})
	return tokens
}

func logLine(idx, ln uint64, rest string) {
	if !enableLogs {
		return
	}

	toPrint := 20
	if idx+uint64(toPrint) > ln {
		toPrint = int(ln - idx)
	}
	log.Printf("parsing %q...\n", rest[:toPrint])
}

func processAvailableTokens(input string) (bool, int, Token) {
	for _, entry := range tokenizerEntries {
		substr, ok := findPattern(input, entry.reg)
		if !ok {
			continue
		}

		if enableLogs {
			log.Printf("found %q -> %v, moving up to %v\n", substr, entry.class, len(substr))
		}

		return true, len(substr), Token{Class: entry.class, Lexeme: substr}
	}

	return false, 0, Token{}
}

func findPattern(input string, reg *regexp.Regexp) (string, bool) {
	res := reg.FindStringSubmatch(input)
	if len(res) < 2 {
		return "", false
	}
	return res[1], true
}
