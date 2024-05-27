package lox

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Lox struct {
	hadError bool
	parser   Parser
}

func NewLox() *Lox {
	return &Lox{
		hadError: false,
	}
}

func (lox *Lox) run(str string) {
	scanner := NewScanner(str)
	tokens := scanner.ScanTokens()
	parser := NewParser(tokens)
	expr := parser.Parse()
	astPrinter := NewAstPrinter()
	astPrinter.Print(expr)
}

func (lox *Lox) RunFile(path string) {
	lox.run(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	lox.run(string(content))
	if lox.hadError {
		os.Exit(65)
	}
}

func (lox *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, _ := reader.ReadLine()
		lox.run(string(line))
		lox.hadError = false
	}
}

func (lox *Lox) Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Println("[" + strconv.Itoa(line) + "] Error " + where + ": " + message)
}

func ErrorHandler(err error, token *Token, message string) {
	if token.tokenType == EOF {
		report(token.line, "at end", message)
	} else {
		report(token.line, "at '"+token.Lexeme+"'", message)
	}
}
