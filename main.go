package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"programming-lang/lexer"
	"programming-lang/parser"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("\nDone", time.Since(start))
	}()

	fmt.Println("Welcome to my bad compiler")
	
	cfg := parseCliArgsToConfig()
	if err := validate(cfg); err != nil {
		fmt.Println(err)
		return
	}
	
	if cfg.runRepl {
		handleRepl(cfg)
		return
	} 
	if cfg.lex {
		printFromFile(cfg.filePath)
	}
	if cfg.parse {
		printAstFromFile(cfg.filePath)
	}
}

func readFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error when reading file: %v", err)
	}
	defer file.Close()

	fileByteContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error in reading file: %v", err)
	}
	return string(fileByteContent), nil
}

type config struct {
	filePath string
	lex bool
	runRepl bool
	parse bool
}

func parseCliArgsToConfig() config {
	var cfg config
	flag.StringVar(&cfg.filePath, "file", "", "path to file with code")
	flag.BoolVar(&cfg.lex, "lex", false, "prints lexer output")
	flag.BoolVar(&cfg.runRepl, "repl", false, "run REPL, ignores all other params")
	flag.BoolVar(&cfg.parse, "parse", false, "prints parser output")
	flag.Parse()

	return cfg
}

func validate(c config) error {
	if !c.runRepl && c.filePath == "" {
		return fmt.Errorf("filepath not provided")
	} 
	return nil
}

func handleRepl(cfg config) {
	fmt.Println("Running repl...")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		
		if cfg.lex {
			printTokens(text)
		}
		if cfg.parse {
			lexParsePrint(text)
		}
	}

	fmt.Println("Closing repl")
}

func printFromFile(filePath string) {
	fileContent, err := readFileContent(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	tokens := lexer.Tokenize(fileContent)
	fmt.Println(tokens)
}

func printTokens(input string) {
	fmt.Println(lexer.Tokenize(input))
}

func printAstFromFile(filePath string) {
	fileContent, err := readFileContent(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	lexParsePrint(fileContent)
}

func lexParsePrint(input string) {
	tokens := lexer.Tokenize(input)
	tree := parser.Parse(tokens)
	fmt.Println(tree)
}